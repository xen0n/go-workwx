//go:build sdkcodegen
// +build sdkcodegen

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const errcodeDocURL = "https://developer.work.weixin.qq.com/document/path/90313"

var absoluteLinkRegexp = regexp.MustCompile(`<a href="(http[^"]+)"(?: rel="nofollow")?>([^<]+)</a>`)

var absoluteLinkReplace = "[$2]($1)"

var anchorLinkRegexp = regexp.MustCompile(`<a href="#([^"]+)"(?: rel="nofollow")?>([^<]+)</a>`)

var anchorLinkReplace = fmt.Sprintf("[$2](%s#$1)", errcodeDocURL)

var h5Regexp = regexp.MustCompile(`<h5[^>]*>.*</h5>`)

var mdLinkRegexp = regexp.MustCompile(`\[([^\]]+)\]\((https?://[^)]+)\)`)

func die(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
	// unreachable
}

func main() {
	// get the fresh documentation!
	var doc *goquery.Document
	var retrieveTime time.Time
	{
		resp, err := http.Get(errcodeDocURL)
		if err != nil {
			die("http get of errcode documentation failed: %+v\n", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			die("non-200 response: %s\n", resp)
		}

		tmp, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			die("parse document failed: %+v\n", err)
		}

		doc = tmp
		retrieveTime = time.Now()
	}

	// prepare to emit code
	destFilename := os.Args[1]
	var sink io.Writer
	{
		file, err := os.Create(destFilename)
		if err != nil {
			die("open '%s' for writing failed: %+v\n", destFilename, err)
		}
		bufWriter := bufio.NewWriter(file)
		sink = bufWriter
		defer func() {
			bufWriter.Flush()
			file.Close()
		}()
	}
	em := &goEmitter{
		Sink: sink,
	}

	err := em.Init(retrieveTime)
	if err != nil {
		die("code emission init failed: %+v\n", err)
	}

	// 本工具最早写作时，所有具体排查方法的小节在页面上是 li > h5 > a 的形式，
	// 正文在 li 里，但截至 2022-02-28 已经变成了平坦的 h5 后面跟一个或多个 p
	// 的形状了。
	// 现在把所有 h5 后面的 p 都预先收集出来，方便下面处理。
	rawSectionContents := collectRawSectionContents(doc)

	numWritten := 0
	// 目前的页面结构是唯一一个 <table> 里面 <tbody> 里面一 <tr> 有按顺序的三个 <td>
	// 错误码, 错误说明, 排查方法
	doc.Find("table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td").First()
		codeStr := td.Text()
		code, err := strconv.ParseInt(codeStr, 10, 64)
		if err != nil {
			die("malformed errcode entry: code=%#v not an int: %+v\n", codeStr, err)
		}

		td = td.Next()
		descStr := td.Text()

		td = td.Next()
		solutionHtml, err := td.Html()
		if err != nil {
			die("failed to get html out of td: %+v\n", err)
		}

		// is there any anchor link to different sections of the same doc?
		var anchorRefs []string
		{
			for _, groups := range anchorLinkRegexp.FindAllStringSubmatch(solutionHtml, -1) {
				anchorName := groups[1]
				if !strings.HasPrefix(anchorName, "10649/") {
					// seems only this format is reference to same doc
					continue
				}
				anchorName = anchorName[6:]

				if unescapedAnchorName, err := url.QueryUnescape(anchorName); err == nil {
					anchorName = unescapedAnchorName
				}

				anchorRefs = append(anchorRefs, anchorName)
			}
		}

		// resolve the referenced section and paste the content into solution
		// for users' convenience
		if len(anchorRefs) > 0 {
			for _, anchor := range anchorRefs {
				solutionHtml += "\n\n"
				solutionHtml += rawSectionContents[anchor]
			}
		}

		// resolve links
		tmp := solutionHtml
		tmp = absoluteLinkRegexp.ReplaceAllString(tmp, absoluteLinkReplace)
		tmp = anchorLinkRegexp.ReplaceAllString(tmp, anchorLinkReplace)
		// hack: remove "10649/" from links
		tmp = strings.ReplaceAll(tmp, "#10649/", "#")

		// unescape things
		// this is VERY crude but working so...
		tmp = strings.ReplaceAll(tmp, "&#34;", `"`)
		tmp = strings.ReplaceAll(tmp, "&lt;", "<")
		tmp = strings.ReplaceAll(tmp, "&gt;", ">")
		tmp = strings.ReplaceAll(tmp, "<br/>", "\n")
		tmp = strings.ReplaceAll(tmp, "<p>", "")
		tmp = strings.ReplaceAll(tmp, "</p>", "")
		tmp = strings.ReplaceAll(tmp, "<ul>", "")
		tmp = strings.ReplaceAll(tmp, "</ul>", "")
		tmp = strings.ReplaceAll(tmp, "<li>", "* ")
		tmp = strings.ReplaceAll(tmp, "</li>", "\n")
		tmp = strings.ReplaceAll(tmp, "<strong>", "**")
		tmp = strings.ReplaceAll(tmp, "</strong>", "**")
		solution := reflowMarkdownLinks(tmp)

		err = em.EmitErrCode(code, descStr, solution)
		if err != nil {
			die("errcode emission failed: %+v\n", err)
		}

		numWritten++
	})

	err = em.Finalize()
	if err != nil {
		die("finalization failed: %+v\n", err)
	}

	fmt.Printf("%d errcodes written.\n", numWritten)
}

func collectRawSectionContents(
	doc *goquery.Document,
) map[string]string {
	result := make(map[string]string)

	var lastSectionHeader string
	var collectedRawContent strings.Builder
	firstContentParagraph := true

	container := doc.Find(".cherry-markdown > div").First()
	container.Children().Each(func(i int, s *goquery.Selection) {
		node := s.Get(0)
		if node.Type == html.ElementNode && strings.ToLower(node.Data) == "h5" {
			if lastSectionHeader != "" {
				// store the previous section
				result[lastSectionHeader] = collectedRawContent.String()

				// prepare for this section
				collectedRawContent.Reset()
				firstContentParagraph = true
			}

			// currently s.Text() is the same as h5[id]
			lastSectionHeader = s.Text()

			return
		}

		// skip everything before the first h5
		if lastSectionHeader == "" {
			return
		}

		rawHTML, err := s.Html()
		if err != nil {
			die("failed to get html out of section: %+v\n", err)
		}

		if firstContentParagraph {
			firstContentParagraph = false
		} else {
			collectedRawContent.WriteString("<br/>")
		}

		collectedRawContent.WriteString(rawHTML)
	})

	return result
}

// Reflows the links in the input Markdown-formatted text so the output is in
// proper go1.19 doc comment form regarding link syntax.
//
// see https://go.dev/doc/comment
func reflowMarkdownLinks(x string) string {
	// construct the output like:
	//
	// 1. input with links trimmed
	// 2. single empty line
	// 3. lines of the format "[text]: link"
	//
	// collect the links to replace along the way
	type linkDesc struct {
		text string
		link string
	}

	var links []linkDesc
	trimmed := mdLinkRegexp.ReplaceAllStringFunc(x, func(fragment string) string {
		match := mdLinkRegexp.FindStringSubmatch(fragment)
		links = append(links, linkDesc{
			text: match[1],
			link: match[2],
		})
		return fmt.Sprintf("[%s]", match[1])
	})

	if len(links) == 0 {
		return x
	}

	var sb strings.Builder
	sb.WriteString(trimmed)
	sb.WriteString("\n\n")
	for _, l := range links {
		sb.WriteRune('[')
		sb.WriteString(l.text)
		sb.WriteString("]: ")
		sb.WriteString(l.link)
		sb.WriteRune('\n')
	}

	return sb.String()
}

//+build sdkcodegen

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const errcodeDocURL = "https://work.weixin.qq.com/api/doc/90000/90139/90313"

var absoluteLinkRegexp = regexp.MustCompile(`<a href="(http[^"]+)">([^<]+)</a>`)

var absoluteLinkReplace = "[$2]($1)"

var anchorLinkRegexp = regexp.MustCompile(`<a href="#([^"]+)">([^<]+)</a>`)

var anchorLinkReplace = fmt.Sprintf("[$2](%s#$1)", errcodeDocURL)

var h5Regexp = regexp.MustCompile(`<h5[^>]*>.*</h5>`)

func die(format string, a ...interface{}) {
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
				anchorRefs = append(anchorRefs, anchorName)
			}
		}

		// resolve the referenced section and paste the content into solution
		// for users' convenience
		//
		// document structure like this: li > h5 > a[name="错误码：xxxxx"]
		// we want the innerHTML of li
		if len(anchorRefs) > 0 {
			for _, anchor := range anchorRefs {
				a := doc.Find(fmt.Sprintf(`a[name="%s"]`, anchor)).First()
				li := a.Parent().Parent()
				liHtml, err := li.Html()
				if err != nil {
					die("failed to get html out of li: %+v\n", err)
				}
				solutionHtml += "\n\n"
				solutionHtml += h5Regexp.ReplaceAllString(liHtml, "")
			}
		}

		// resolve links
		tmp := solutionHtml
		tmp = absoluteLinkRegexp.ReplaceAllString(tmp, absoluteLinkReplace)
		tmp = anchorLinkRegexp.ReplaceAllString(tmp, anchorLinkReplace)

		// unescape things
		// this is VERY crude but working so...
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
		solution := tmp

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

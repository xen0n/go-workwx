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

	"github.com/PuerkitoBio/goquery"
)

const errcodeDocURL = "https://work.weixin.qq.com/api/doc/90000/90139/90313"

var absoluteLinkRegexp = regexp.MustCompile(`<a href="(http[^"]+)">([^<]+)</a>`)

var absoluteLinkReplace = "[$2]($1)"

var anchorLinkRegexp = regexp.MustCompile(`<a href="#([^"]+)">([^<]+)</a>`)

var anchorLinkReplace = fmt.Sprintf("[$2](%s#$1)", errcodeDocURL)

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
	// unreachable
}

func main() {
	// get the fresh documentation!
	var doc *goquery.Document
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

	err := em.Init()
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

		// resolve links
		tmp := solutionHtml
		tmp = absoluteLinkRegexp.ReplaceAllString(tmp, absoluteLinkReplace)
		tmp = anchorLinkRegexp.ReplaceAllString(tmp, anchorLinkReplace)

		// unescape things
		// this is VERY crude but working so...
		tmp = strings.ReplaceAll(tmp, "&lt;", "<")
		tmp = strings.ReplaceAll(tmp, "&gt;", ">")
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

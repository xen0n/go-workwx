package main

import (
	"fmt"

	"gopkg.in/russross/blackfriday.v2"
)

func parseDocument(content []byte) {
	md := blackfriday.New(blackfriday.WithExtensions(blackfriday.CommonExtensions))
	root := md.Parse(content)

	parser := &docParser{
		isDebug: true,
	}
	root.Walk(parser.Walk)
}

type docParser struct {
	isDebug bool
}

func (p *docParser) dbg(format string, a ...interface{}) (n int, err error) {
	if !p.isDebug {
		return 0, nil
	}

	return fmt.Printf(format, a...)
}

func (p *docParser) Walk(n *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if !entering {
		return blackfriday.GoToNext
	}

	p.dbg("node: %+v\n", n)

	return blackfriday.GoToNext
}

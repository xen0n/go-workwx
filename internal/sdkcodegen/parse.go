//+build sdkcodegen

package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func parseDocument(content []byte) *mdTocNode {
	md := blackfriday.New(blackfriday.WithExtensions(blackfriday.CommonExtensions))

	// normalize to LF line ending for correct parsing
	// https://github.com/russross/blackfriday/issues/423
	normalizedContent := normalizeEOL(content)

	root := md.Parse(normalizedContent)

	return reshapeMarkdownAST(root)
}

// Reshape the blackfriday AST into something more workable.
//
// From e.g.:
//
// ```
// Document
// |- H1
// |  |- Text
// |- H2
// |  |- Text
// |  |- Code
// |  |- Text
// |- H3
// |- Table
// |- H3
// |- Table
// ```
//
// To:
//
// ```
// Document(Content: [])
// |- H1(Content: [Text])
//    |- H2(Content: [Text Code Text])
//       |- H3(Content: [Table])
//       |- H3(Content: [Table])
// ```
func reshapeMarkdownAST(root *blackfriday.Node) *mdTocNode {
	stack := []mdNode{
		&mdTocNode{mdContentNode: &mdContentNode{This: root}},
	}

	// traverse the already somewhat-flat blackfriday AST
	for n := root.FirstChild; n != nil; n = n.Next {
		stack = doReshape(n, stack)
	}

	// finalize
	stack = foldTocNode(stack, 0)

	return stack[0].(*mdTocNode)
}

// stack is never empty
func doReshape(
	root *blackfriday.Node,
	stack []mdNode,
) []mdNode {
	// fmt.Printf(">>> doReshape\n    root=%s\n    stack=%s\n", root, stack)
	if root == nil {
		return stack
	}

	switch root.Type {
	case blackfriday.Heading:
		lvl := root.HeadingData.Level
		stack = foldTocNode(stack, lvl)

		// push this heading onto stack for construction
		thisNodeContent := makeContentNode(root)
		stack = append(stack, &mdTocNode{
			mdContentNode: thisNodeContent,
			Level:         lvl,
		})

	default:
		contentNode := makeContentNode(root)
		stack = append(stack, contentNode)
	}

	return stack
}

func makeContentNode(root *blackfriday.Node) *mdContentNode {
	// cull blank text nodes
	switch root.Type {
	case blackfriday.Text:
		if len(root.Literal) == 0 {
			return nil
		}
	}

	result := &mdContentNode{
		This: root,
	}

	for n := root.FirstChild; n != nil; n = n.Next {
		cn := makeContentNode(n)
		if cn == nil {
			// this node is culled
			continue
		}

		result.ThisContent = append(result.ThisContent, cn)
	}

	return result
}

func foldTocNode(stack []mdNode, lvl int) []mdNode {
	// finalize content nodes seen so far onto the nearest outline node
	stktop := stackTop(stack)
	if !isTocNode(stktop) {
		var contentNodes []mdNode
		stack, contentNodes = popUntil(stack, func(n mdNode) bool {
			return !isTocNode(n)
		})
		// fmt.Printf("stack=%s\ncontentNodes=%s\n", stack, contentNodes)
		stktop = stackTop(stack)
		outlineNode := stktop.(*mdTocNode)

		for _, n := range contentNodes {
			outlineNode.Content = append(outlineNode.Content, n.(*mdContentNode))
		}
	}

	// see if we have to finalize some headers
	for {
		stktopTocNode := stackTop(stack).(*mdTocNode)
		stackLvl := stktopTocNode.This.HeadingData.Level

		if lvl > stackLvl || len(stack) == 1 {
			break
		}

		// finalize children up to nearest parent
		// the stack should only consist of outline nodes
		tocChild := stktopTocNode
		stack = stack[:len(stack)-1]
		tocParent := stackTop(stack).(*mdTocNode)
		tocParent.TocChildren = append(tocParent.TocChildren, tocChild)
	}

	return stack
}

func dumpToctree(n *mdTocNode) {
	var sb strings.Builder
	doDumpToctree(&sb, n, 0)
	fmt.Printf("%s", sb.String())
}

func doDumpToctree(sink io.Writer, n *mdTocNode, depth int) {
	fmt.Fprintf(sink, "%s%s", strings.Repeat("  ", depth), n.displayType())
	doDumpContent(sink, n.ThisContent)
	doDumpContent(sink, n.Content)
	fmt.Fprintf(sink, "\n")

	for _, tn := range n.TocChildren {
		doDumpToctree(sink, tn, depth+1)
	}
}

func doDumpContent(sink io.Writer, l []*mdContentNode) {
	if len(l) > 0 {
		fmt.Fprintf(sink, "(")
		first := true
		for _, cn := range l {
			if first {
				first = false
			} else {
				fmt.Fprintf(sink, ", ")
			}
			fmt.Fprintf(sink, "%s", cn.displayType())
		}
		fmt.Fprintf(sink, ")")
	}
}

func stackTop(s []mdNode) mdNode {
	return s[len(s)-1]
}

func popUntil(
	s []mdNode,
	predicate func(mdNode) bool,
) (newStack []mdNode, popped []mdNode) {
	ptr := len(s) - 1
	for {
		if !predicate(s[ptr]) {
			break
		}
		ptr--
	}

	return s[:ptr+1], s[ptr+1:]
}

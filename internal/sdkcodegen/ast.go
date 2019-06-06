package main

import (
	"fmt"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

type mdNode interface {
	sealedForMdNode()

	displayType() string
}

type mdContentNode struct {
	This        *blackfriday.Node
	ThisContent []*mdContentNode
	Content     []*mdContentNode
}

func (x *mdContentNode) sealedForMdNode() {}

func (x *mdContentNode) displayType() string {
	var sb strings.Builder

	sb.WriteString("C<")
	sb.WriteString(x.This.Type.String())
	sb.WriteRune('>')

	return sb.String()
}

type mdTocNode struct {
	*mdContentNode

	Level       int
	TocChildren []*mdTocNode
}

func (x *mdTocNode) sealedForMdNode() {}

func (x *mdTocNode) displayType() string {
	return fmt.Sprintf("T<%s<%d>>", x.This.Type.String(), x.Level)
}

func isTocNode(n mdNode) bool {
	_, ok := n.(*mdTocNode)
	return ok
}

func (x *mdContentNode) String() string {
	var sb strings.Builder

	sb.WriteString(x.displayType())

	if len(x.ThisContent) == 0 && len(x.Content) == 0 {
		goto out
	}

	sb.WriteRune('{')
	if len(x.ThisContent) > 0 {
		sb.WriteString(" tC=")
		fmt.Fprintf(&sb, "%+v", x.ThisContent)
	}
	if len(x.Content) > 0 {
		sb.WriteString(" C=")
		fmt.Fprintf(&sb, "%+v", x.Content)
	}
	sb.WriteRune('}')

out:
	return sb.String()
}

func (x *mdTocNode) String() string {
	var sb strings.Builder

	sb.WriteString(x.displayType())

	if len(x.ThisContent) == 0 && len(x.Content) == 0 && len(x.TocChildren) == 0 {
		goto out
	}

	sb.WriteRune('{')
	if len(x.ThisContent) > 0 {
		sb.WriteString(" tC=")
		fmt.Fprintf(&sb, "%+v", x.ThisContent)
	}
	if len(x.Content) > 0 {
		sb.WriteString(" C=")
		fmt.Fprintf(&sb, "%+v", x.Content)
	}
	if len(x.TocChildren) > 0 {
		sb.WriteString(" T=")
		fmt.Fprintf(&sb, "%+v", x.TocChildren)
	}
	sb.WriteRune('}')

out:
	return sb.String()
}

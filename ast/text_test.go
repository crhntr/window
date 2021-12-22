package ast

import (
	"testing"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
	"github.com/crhntr/window/dom/domtest"
)

var _ dom.Text = (*Text)(nil)

func createTextNode(_ *testing.T) dom.ChildNode {
	return &Text{node: &html.Node{
		Type: html.TextNode,
		Data: "node",
	}}
}

func TestText(t *testing.T) {
	domtest.Node(t, func(t *testing.T) dom.Node {
		return createTextNode(t)
	})
}

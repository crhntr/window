package ast

import (
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
	"github.com/crhntr/window/dom/domtest"
)

func createDocumentNode(t *testing.T) dom.Node {
	node, _ := html.Parse(strings.NewReader(""))
	return &Document{node: node}
}

func TestDocument(t *testing.T) {
	domtest.Node(t, createDocumentNode)
}

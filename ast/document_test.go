package ast

import (
	_ "embed"
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
	"github.com/crhntr/window/dom/domtest"
)

//go:embed testdata/document_01.html
var document01 string

func createDocumentNode(t *testing.T) dom.Node {
	node, _ := html.Parse(strings.NewReader(document01))
	return &Document{node: node}
}

func TestDocument(t *testing.T) {
	domtest.Node(t, createDocumentNode)
	domtest.Document(t, createDocumentNode)
}

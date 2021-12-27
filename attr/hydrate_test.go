package attr

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/crhntr/window/ast"
	"github.com/crhntr/window/dom"
)

func divEl() *html.Node {
	return &html.Node{Type: html.ElementNode, Data: atom.Div.String(), DataAtom: atom.Div}
}

func CreateElement(t *testing.T, s string) dom.Element {
	nodes, err := html.ParseFragment(strings.NewReader(s), divEl())
	if err != nil {
		t.Fatal(err)
	}
	return ast.NewElement(nodes[0])
}

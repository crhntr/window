package template

import (
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
	"github.com/crhntr/window/domtest"
)

func createNode(t *testing.T, nodeType dom.NodeType, input string) dom.Node {
	switch nodeType {
	case dom.NodeTypeElement:
		htmlNode, err := html.ParseFragment(strings.NewReader(input), &html.Node{Type: html.ElementNode})
		if err != nil {
			t.Fatal(err)
		}
		return &Element{Node: *htmlNode[0]}
	case dom.NodeTypeDocument:
		//htmlNode, err := html.ParseFragment(strings.NewReader(input), nil)
		//if err != nil {
		//	panic(err)
		//}
		//return &Document{Node: *htmlNode[0]}
	}
	t.Fatal("unsupported node type")
	return nil
}

func TestElement(t *testing.T) {
	domtest.ElementTagName(t, createNode)
	domtest.ElementNodeType(t, createNode)
}

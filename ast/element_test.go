package ast

import (
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
	"github.com/crhntr/window/dom/domtest"
)

func createElement(t *testing.T, input string) dom.Element {
	htmlNode, err := html.ParseFragment(strings.NewReader(input), &html.Node{Type: html.ElementNode})
	if err != nil {
		t.Fatal(err)
	}
	return &Element{node: htmlNode[0]}
}

func createElementNode(t *testing.T) dom.ChildNode {
	htmlNode, err := html.ParseFragment(strings.NewReader("<div></div>"), &html.Node{Type: html.ElementNode})
	if err != nil {
		t.Fatal(err)
	}
	return &Element{node: htmlNode[0]}
}

func createElementParentNode(t *testing.T) dom.ParentNode {
	return createElementNode(t).(dom.ParentNode)
}

func TestElement(t *testing.T) {
	domtest.Node(t, func(t *testing.T) dom.Node {
		return createElementNode(t)
	})
	t.Run("element children", func(t *testing.T) {
		domtest.ParentNode(t, createElementParentNode, createElementNode)
	})
	t.Run("text children", func(t *testing.T) {
		domtest.ParentNode(t, createElementParentNode, createTextNode)
	})

	domtest.ElementQueries(t, createElement)
	domtest.ElementTextContent(t, createElement)

	domtest.ElementTagName(t, createElement)
	domtest.ElementAttribute(t, createElement)
	domtest.ElementInnerHTML(t, createElement)
	domtest.ElementOuterHTML(t, createElement)
}

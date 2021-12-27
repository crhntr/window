package ast

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/crhntr/window/dom"
)

type Document struct {
	node *html.Node
}

func (d *Document) Head() dom.Element {
	head := firstElementWithTag(d.node, atom.Head.String())
	if head == nil {
		return nil
	}
	return &Element{
		node: head,
	}
}

func (d *Document) Body() dom.Element {
	head := firstElementWithTag(d.node, atom.Body.String())
	if head == nil {
		return nil
	}
	return &Element{
		node: head,
	}
}

// Node

func (d *Document) NodeType() dom.NodeType         { return nodeType(d.node.Type) }
func (d *Document) CloneNode(deep bool) dom.Node   { return cloneNode(d.node, deep) }
func (d *Document) IsSameNode(other dom.Node) bool { return isSameNode(d.node, other) }

//func (d *Document) QuerySelector(query string) dom.Element { return querySelector(d.node, query) }
//func (d *Document) QuerySelectorAll(query string) dom.NodeList {
//	return querySelectorAll(d.node, query)
//}

func (d *Document) Contains(other dom.Node) bool { return contains(d.node, other) }

func (d *Document) TextContent() string { return textContent(d.node) }

// Document

func (*Document) CreateElement(localName string) dom.Element {
	a := atom.Lookup([]byte(localName))
	if a == 0 {
		return &Element{
			node: &html.Node{
				Type:     html.ElementNode,
				DataAtom: 0,
				Data:     localName,
			},
		}
	}

	return &Element{
		node: &html.Node{
			Type:     html.ElementNode,
			DataAtom: a,
			Data:     a.String(),
		},
	}
}

func (*Document) CreateTextNode(text string) dom.Text {
	return &Text{
		node: &html.Node{
			Type: html.TextNode,
			Data: text,
		},
	}
}

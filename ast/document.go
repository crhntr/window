package ast

import (
	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
)

type Document struct {
	node *html.Node
}

// Node

func (d *Document) NodeType() dom.NodeType         { return nodeType(d.node.Type) }
func (d *Document) CloneNode(deep bool) dom.Node   { return cloneNode(d.node, deep) }
func (d *Document) IsSameNode(other dom.Node) bool { return isSameNode(d.node, other) }
func (d *Document) GetElementsByTagName(name string) dom.ElementCollection {
	return getElementsByTagName(d.node, name)
}
func (d *Document) GetElementsByClassName(name string) dom.ElementCollection {
	return getElementsByClassName(d.node, name)
}
func (d *Document) QuerySelector(query string) dom.Element { return querySelector(d.node, query) }
func (d *Document) QuerySelectorAll(query string) dom.NodeList {
	return querySelectorAll(d.node, query)
}
func (d *Document) Contains(other dom.Node) bool { return contains(d.node, other) }

func (d *Document) TextContent() string { return textContent(d.node) }

// Document

func (d *Document) CreateElement(localName string) dom.Element {
	//TODO implement me
	panic("implement me")
}

func (d *Document) CreateElementIs(localName, is string) dom.Element {
	//TODO implement me
	panic("implement me")
}

func (d *Document) CreateTextNode(text string) dom.Text {
	//TODO implement me
	panic("implement me")
}

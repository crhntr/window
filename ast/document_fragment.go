package ast

import (
	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
)

type DocumentFragment struct {
	nodes []*html.Node
}

func (d DocumentFragment) NodeType() dom.NodeType { return dom.NodeTypeDocumentFragment }

func (d DocumentFragment) CloneNode(deep bool) dom.Node {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) IsSameNode(other dom.Node) bool {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) TextContent() string {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) Children() dom.ElementCollection {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) FirstElementChild() dom.Element {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) LastElementChild() dom.Element {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) ChildElementCount() int {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) Prepend(nodes ...dom.ChildNode) {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) Append(nodes ...dom.ChildNode) {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) ReplaceChildren(nodes ...dom.ChildNode) {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) Contains(other dom.Node) bool {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) GetElementsByTagName(name string) dom.ElementCollection {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) GetElementsByClassName(name string) dom.ElementCollection {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) QuerySelector(query string) dom.Element {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) QuerySelectorAll(query string) dom.NodeList {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) HasChildNodes() bool {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) ChildNodes() dom.NodeList {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) FirstChild() dom.ChildNode {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) LastChild() dom.ChildNode {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) InsertBefore(node, child dom.ChildNode) dom.ChildNode {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) AppendChild(node dom.ChildNode) dom.ChildNode {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) ReplaceChild(node, child dom.ChildNode) dom.ChildNode {
	//TODO implement me
	panic("implement me")
}

func (d DocumentFragment) RemoveChild(node dom.ChildNode) dom.ChildNode {
	//TODO implement me
	panic("implement me")
}

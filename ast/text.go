package ast

import (
	"golang.org/x/net/html"

	"github.com/crhntr/window/dom"
)

type Text struct {
	node *html.Node
}

func (t Text) Data() string     { return t.node.Data }
func (t Text) SetData(d string) { t.node.Data = d }

func (t Text) NodeType() dom.NodeType         { return nodeType(t.node.Type) }
func (t Text) IsConnected() bool              { return isConnected(t.node) }
func (t Text) OwnerDocument() dom.Document    { return ownerDocument(t.node) }
func (t Text) Length() int                    { return len(t.node.Data) }
func (t Text) ParentNode() dom.Node           { return parentNode(t.node) }
func (t Text) ParentElement() dom.Element     { return parentElement(t.node) }
func (t Text) PreviousSibling() dom.ChildNode { return previousSibling(t.node) }
func (t Text) NextSibling() dom.ChildNode     { return nextSibling(t.node) }
func (t Text) TextContent() string            { return t.node.Data }
func (t Text) CloneNode(_ bool) dom.Node {
	return Text{
		node: &html.Node{
			Type: html.TextNode,
			Data: t.node.Data,
		},
	}
}
func (t Text) IsSameNode(other dom.Node) bool { return isSameNode(t.node, other) }

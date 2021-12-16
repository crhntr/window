//go:build js

package window

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type DocumentFragment js.Value

func (d DocumentFragment) NodeType() dom.NodeType             { return v(d).nodeType() }
func (d DocumentFragment) IsConnected() bool                  { return v(d).isConnected() }
func (d DocumentFragment) OwnerDocument() dom.Node            { return v(d).ownerDocument() }
func (d DocumentFragment) GetRootNode(composed bool) dom.Node { return v(d).getRootNode(composed) }
func (d DocumentFragment) ParentNode() dom.Node               { return v(d).parentNode() }
func (d DocumentFragment) ParentElement() dom.Element         { return v(d).parentElement() }
func (d DocumentFragment) HasChildNodes() bool                { return v(d).hasChildNodes() }
func (d DocumentFragment) ChildNodes() dom.NodeList           { return v(d).childNodes() }
func (d DocumentFragment) FirstChild() dom.Node               { return v(d).firstChild() }
func (d DocumentFragment) LastChild() dom.Node                { return v(d).lastChild() }
func (d DocumentFragment) PreviousSibling() dom.Node          { return v(d).previousSibling() }
func (d DocumentFragment) NextSibling() dom.Node              { return v(d).nextSibling() }
func (d DocumentFragment) NodeValue() string                  { return v(d).nodeValue() }
func (d DocumentFragment) TextContent() string                { return v(d).textContent() }
func (d DocumentFragment) Normalize()                         { v(d).normalize() }
func (d DocumentFragment) CloneNode(deep bool) dom.Node       { return v(d).cloneNode(deep) }
func (d DocumentFragment) IsEqualNode(other dom.Node) bool    { return v(d).isEqualNode(other) }
func (d DocumentFragment) IsSameNode(other dom.Node) bool     { return v(d).isSameNode(other) }
func (d DocumentFragment) CompareDocumentPosition() dom.DocumentPosition {
	return v(d).compareDocumentPosition()
}
func (d DocumentFragment) Contains(other dom.Node) bool { return v(d).contains(other) }
func (d DocumentFragment) InsertBefore(node, child dom.Node) dom.Node {
	return v(d).insertBefore(node, child)
}
func (d DocumentFragment) AppendChild(node dom.Node) dom.Node { return v(d).appendChild(node) }
func (d DocumentFragment) ReplaceChild(node, child dom.Node) dom.Node {
	return v(d).replaceChild(node, child)
}
func (d DocumentFragment) RemoveChild(node dom.Node) dom.Node { return v(d).removeChild(node) }

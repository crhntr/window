//go:build js && wasm
// +build js,wasm

package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type DocumentFragment js.Value

func (d DocumentFragment) NodeType() dom.NodeType         { return v(d).nodeType() }
func (d DocumentFragment) IsConnected() bool              { return v(d).isConnected() }
func (d DocumentFragment) OwnerDocument() dom.Document    { return v(d).ownerDocument() }
func (d DocumentFragment) ParentNode() dom.Node           { return v(d).parentNode() }
func (d DocumentFragment) ParentElement() dom.Element     { return v(d).parentElement() }
func (d DocumentFragment) HasChildNodes() bool            { return v(d).hasChildNodes() }
func (d DocumentFragment) ChildNodes() dom.NodeList       { return v(d).childNodes() }
func (d DocumentFragment) FirstChild() dom.ChildNode      { return v(d).firstChild() }
func (d DocumentFragment) LastChild() dom.ChildNode       { return v(d).lastChild() }
func (d DocumentFragment) PreviousSibling() dom.ChildNode { return v(d).previousSibling() }
func (d DocumentFragment) NextSibling() dom.ChildNode     { return v(d).nextSibling() }
func (d DocumentFragment) TextContent() string            { return v(d).textContent() }
func (d DocumentFragment) Normalize()                     { v(d).normalize() }
func (d DocumentFragment) CloneNode(deep bool) dom.Node   { return v(d).cloneNode(deep) }
func (d DocumentFragment) IsSameNode(other dom.Node) bool { return v(d).isSameNode(other) }
func (d DocumentFragment) CompareDocumentPosition() dom.DocumentPosition {
	return v(d).compareDocumentPosition()
}
func (d DocumentFragment) Contains(other dom.Node) bool { return v(d).contains(other) }
func (d DocumentFragment) InsertBefore(node, child dom.ChildNode) dom.ChildNode {
	return v(d).insertBefore(node, child)
}
func (d DocumentFragment) AppendChild(node dom.ChildNode) dom.ChildNode {
	return v(d).appendChild(node)
}
func (d DocumentFragment) ReplaceChild(node, child dom.ChildNode) dom.ChildNode {
	return v(d).replaceChild(node, child)
}
func (d DocumentFragment) RemoveChild(node dom.ChildNode) dom.ChildNode {
	return v(d).removeChild(node)
}
func (d DocumentFragment) Children() dom.ElementCollection        { return v(d).children() }
func (d DocumentFragment) FirstElementChild() dom.Element         { return v(d).firstElementChild() }
func (d DocumentFragment) LastElementChild() dom.Element          { return v(d).lastElementChild() }
func (d DocumentFragment) ChildElementCount() int                 { return v(d).childElementCount() }
func (d DocumentFragment) Prepend(nodes ...dom.ChildNode)         { v(d).prepend(nodes) }
func (d DocumentFragment) Append(nodes ...dom.ChildNode)          { v(d).append(nodes) }
func (d DocumentFragment) ReplaceChildren(nodes ...dom.ChildNode) { v(d).replaceChildren(nodes) }
func (d DocumentFragment) GetElementsByTagName(name string) dom.ElementCollection {
	return v(d).getElementsByTagName(name)
}
func (d DocumentFragment) GetElementsByClassName(name string) dom.ElementCollection {
	return v(d).getElementsByClassName(name)
}
func (d DocumentFragment) QuerySelector(query string) dom.Element { return v(d).querySelector(query) }
func (d DocumentFragment) QuerySelectorAll(query string) dom.NodeList {
	return v(d).querySelectorAll(query)
}
func (d DocumentFragment) Length() int { return v(d).length() }

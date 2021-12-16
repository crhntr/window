package window

import (
	"github.com/crhntr/window/dom"
	"syscall/js"
)

type document js.Value

func (d document) NodeType() dom.NodeType             { return v(d).nodeType() }
func (d document) IsConnected() bool                  { return v(d).isConnected() }
func (d document) OwnerDocument() dom.Node            { return v(d).ownerDocument() }
func (d document) GetRootNode(composed bool) dom.Node { return v(d).getRootNode(composed) }
func (d document) ParentNode() dom.Node               { return v(d).parentNode() }
func (d document) ParentElement() dom.Element         { return v(d).parentElement() }
func (d document) HasChildNodes() bool                { return v(d).hasChildNodes() }
func (d document) ChildNodes() dom.NodeList           { return v(d).childNodes() }
func (d document) FirstChild() dom.Node               { return v(d).firstChild() }
func (d document) LastChild() dom.Node                { return v(d).lastChild() }
func (d document) PreviousSibling() dom.Node          { return v(d).previousSibling() }
func (d document) NextSibling() dom.Node              { return v(d).nextSibling() }
func (d document) NodeValue() string                  { return v(d).nodeValue() }
func (d document) TextContent() string                { return v(d).textContent() }
func (d document) Normalize()                         { v(d).normalize() }
func (d document) CloneNode(deep bool) dom.Node       { return v(d).cloneNode(deep) }
func (d document) IsEqualNode(other dom.Node) bool    { return v(d).isEqualNode(other) }
func (d document) IsSameNode(other dom.Node) bool     { return v(d).isSameNode(other) }
func (d document) CompareDocumentPosition() dom.DocumentPosition {
	return v(d).compareDocumentPosition()
}
func (d document) Contains(other dom.Node) bool { return v(d).contains(other) }
func (d document) InsertBefore(node, child dom.Node) dom.Node {
	return v(d).insertBefore(node, child)
}
func (d document) AppendChild(node dom.Node) dom.Node { return v(d).appendChild(node) }
func (d document) ReplaceChild(node, child dom.Node) dom.Node {
	return v(d).replaceChild(node, child)
}
func (d document) RemoveChild(node dom.Node) dom.Node         { return v(d).removeChild(node) }
func (d document) Children() dom.NodeList                     { return v(d).children() }
func (d document) FirstElementChild() dom.Element             { return v(d).firstElementChild() }
func (d document) LastElementChild() dom.Element              { return v(d).lastElementChild() }
func (d document) ChildElementCount() int                     { return v(d).childElementCount() }
func (d document) Prepend(nodes ...dom.Node) dom.Node         { return v(d).prepend(nodes) }
func (d document) Append(nodes ...dom.Node) dom.Node          { return v(d).append(nodes) }
func (d document) ReplaceChildren(nodes ...dom.Node) dom.Node { return v(d).replaceChildren(nodes) }
func (d document) GetElementsByTagName(name string) dom.ElementCollection {
	return v(d).getElementsByTagName(name)
}
func (d document) GetElementsByTagNameNS(namespace, localName string) dom.ElementCollection {
	return v(d).getElementsByTagNameNS(namespace, localName)
}
func (d document) GetElementsByClassName(name string) dom.ElementCollection {
	return v(d).getElementsByClassName(name)
}
func (d document) QuerySelector(query string) dom.Element { return v(d).querySelector(query) }
func (d document) QuerySelectorAll(query string) dom.NodeList {
	return v(d).querySelectorAll(query)
}
func (d document) CreateElement(localName string) dom.Element {
	return v(d).createElement(localName)
}
func (d document) CreateElementIs(localName, is string) dom.Element {
	return v(d).createElementIs(localName, is)
}
func (d document) CreateElementNS(namespace, localName string) dom.Element {
	return v(d).createElementNS(namespace, localName)
}
func (d document) CreateElementNSIS(namespace, localName, is string) dom.Element {
	return v(d).createElementNSIS(namespace, localName, is)
}
func (d document) CreateDocumentFragment() DocumentFragment { return v(d).createDocumentFragment() }
func (d document) CreateTextNode(text string) dom.Text      { return v(d).createTextNode(text) }

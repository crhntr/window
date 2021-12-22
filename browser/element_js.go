package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type Element js.Value

func (e Element) NodeType() dom.NodeType         { return v(e).nodeType() }
func (e Element) IsConnected() bool              { return v(e).isConnected() }
func (e Element) OwnerDocument() dom.Document    { return v(e).ownerDocument() }
func (e Element) ParentNode() dom.Node           { return v(e).parentNode() }
func (e Element) ParentElement() dom.Element     { return v(e).parentElement() }
func (e Element) HasChildNodes() bool            { return v(e).hasChildNodes() }
func (e Element) ChildNodes() dom.NodeList       { return v(e).childNodes() }
func (e Element) FirstChild() dom.ChildNode      { return v(e).firstChild() }
func (e Element) LastChild() dom.ChildNode       { return v(e).lastChild() }
func (e Element) PreviousSibling() dom.ChildNode { return v(e).previousSibling() }
func (e Element) NextSibling() dom.ChildNode     { return v(e).nextSibling() }
func (e Element) TextContent() string            { return v(e).textContent() }
func (e Element) Normalize()                     { v(e).normalize() }
func (e Element) CloneNode(deep bool) dom.Node   { return v(e).cloneNode(deep) }
func (e Element) IsSameNode(other dom.Node) bool { return v(e).isSameNode(other) }
func (e Element) CompareDocumentPosition() dom.DocumentPosition {
	return v(e).compareDocumentPosition()
}
func (e Element) Contains(other dom.Node) bool { return v(e).contains(other) }
func (e Element) InsertBefore(node, child dom.ChildNode) dom.ChildNode {
	return v(e).insertBefore(node, child)
}
func (e Element) AppendChild(node dom.ChildNode) dom.ChildNode { return v(e).appendChild(node) }
func (e Element) ReplaceChild(node, child dom.ChildNode) dom.ChildNode {
	return v(e).replaceChild(node, child)
}
func (e Element) RemoveChild(node dom.ChildNode) dom.ChildNode { return v(e).removeChild(node) }
func (e Element) Children() dom.ElementCollection              { return v(e).children() }
func (e Element) ChildElementCount() int                       { return v(e).childElementCount() }
func (e Element) FirstElementChild() dom.Element               { return v(e).firstElementChild() }
func (e Element) LastElementChild() dom.Element                { return v(e).lastElementChild() }
func (e Element) Prepend(nodes ...dom.ChildNode)               { v(e).prepend(nodes) }
func (e Element) Append(nodes ...dom.ChildNode)                { v(e).append(nodes) }
func (e Element) ReplaceChildren(nodes ...dom.ChildNode)       { v(e).replaceChildren(nodes) }
func (e Element) GetElementsByTagName(name string) dom.ElementCollection {
	return v(e).getElementsByTagName(name)
}
func (e Element) GetElementsByClassName(name string) dom.ElementCollection {
	return v(e).getElementsByClassName(name)
}
func (e Element) QuerySelector(query string) dom.Element     { return v(e).querySelector(query) }
func (e Element) QuerySelectorAll(query string) dom.NodeList { return v(e).querySelectorAll(query) }

func (e Element) TagName() string                     { return v(e).tagName() }
func (e Element) ID() string                          { return v(e).iD() }
func (e Element) ClassName() string                   { return v(e).className() }
func (e Element) GetAttribute(name string) string     { return v(e).getAttribute(name) }
func (e Element) SetAttribute(name, value string)     { v(e).setAttribute(name, value) }
func (e Element) RemoveAttribute(name string)         { v(e).removeAttribute(name) }
func (e Element) ToggleAttribute(name string) bool    { return v(e).toggleAttribute(name) }
func (e Element) HasAttribute(name string) bool       { return v(e).hasAttribute(name) }
func (e Element) Closest(selector string) dom.Element { return v(e).closest(selector) }
func (e Element) Matches(selector string) bool        { return v(e).matches(selector) }

func (e Element) SetInnerHTML(s string) { v(e).setInnerHTML(s) }
func (e Element) InnerHTML() string     { return v(e).innerHTML() }
func (e Element) SetOuterHTML(s string) { v(e).setOuterHTML(s) }
func (e Element) OuterHTML() string     { return v(e).outerHTML() }
func (e Element) SetInnerText(s string) { v(e).setInnerText(s) }
func (e Element) InnerText() string     { return v(e).innerText() }

func (e Element) Length() int { return v(e).length() }

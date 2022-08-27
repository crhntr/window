//go:build js && wasm

package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type Input js.Value

func (in Input) Value() string             { return v(in).value() }
func (in Input) SetValue(value string)     { v(in).setValue(value) }
func (in Input) Disabled() bool            { return v(in).disabled() }
func (in Input) SetDisabled(disabled bool) { v(in).setDisabled(disabled) }

// The following implement Element

func (in Input) NodeType() dom.NodeType         { return v(in).nodeType() }
func (in Input) IsConnected() bool              { return v(in).isConnected() }
func (in Input) OwnerDocument() dom.Document    { return v(in).ownerDocument() }
func (in Input) ParentNode() dom.Node           { return v(in).parentNode() }
func (in Input) ParentElement() dom.Element     { return v(in).parentElement() }
func (in Input) HasChildNodes() bool            { return v(in).hasChildNodes() }
func (in Input) ChildNodes() dom.NodeList       { return v(in).childNodes() }
func (in Input) FirstChild() dom.ChildNode      { return v(in).firstChild() }
func (in Input) LastChild() dom.ChildNode       { return v(in).lastChild() }
func (in Input) PreviousSibling() dom.ChildNode { return v(in).previousSibling() }
func (in Input) NextSibling() dom.ChildNode     { return v(in).nextSibling() }
func (in Input) TextContent() string            { return v(in).textContent() }
func (in Input) Normalize()                     { v(in).normalize() }
func (in Input) CloneNode(deep bool) dom.Node   { return v(in).cloneNode(deep) }
func (in Input) IsSameNode(other dom.Node) bool { return v(in).isSameNode(other) }
func (in Input) CompareDocumentPosition() dom.DocumentPosition {
	return v(in).compareDocumentPosition()
}
func (in Input) Contains(other dom.Node) bool { return v(in).contains(other) }
func (in Input) InsertBefore(node, child dom.ChildNode) dom.ChildNode {
	return v(in).insertBefore(node, child)
}
func (in Input) AppendChild(node dom.ChildNode) dom.ChildNode { return v(in).appendChild(node) }
func (in Input) ReplaceChild(node, child dom.ChildNode) dom.ChildNode {
	return v(in).replaceChild(node, child)
}
func (in Input) RemoveChild(node dom.ChildNode) dom.ChildNode { return v(in).removeChild(node) }
func (in Input) Children() dom.ElementCollection              { return v(in).children() }
func (in Input) ChildElementCount() int                       { return v(in).childElementCount() }
func (in Input) FirstElementChild() dom.Element               { return v(in).firstElementChild() }
func (in Input) LastElementChild() dom.Element                { return v(in).lastElementChild() }
func (in Input) Prepend(nodes ...dom.ChildNode)               { v(in).prepend(nodes) }
func (in Input) Append(nodes ...dom.ChildNode)                { v(in).append(nodes) }
func (in Input) ReplaceChildren(nodes ...dom.ChildNode)       { v(in).replaceChildren(nodes) }

func (in Input) GetElementByID(id string) dom.Element { return v(in).getElementById(id) }
func (in Input) GetElementsByTagName(name string) dom.ElementCollection {
	return v(in).getElementsByTagName(name)
}
func (in Input) GetElementsByClassName(name string) dom.ElementCollection {
	return v(in).getElementsByClassName(name)
}
func (in Input) QuerySelector(query string) dom.Element     { return v(in).querySelector(query) }
func (in Input) QuerySelectorAll(query string) dom.NodeList { return v(in).querySelectorAll(query) }

func (in Input) TagName() string                     { return v(in).tagName() }
func (in Input) ID() string                          { return v(in).iD() }
func (in Input) ClassName() string                   { return v(in).className() }
func (in Input) Attribute(name string) string        { return v(in).getAttribute(name) }
func (in Input) SetAttribute(name, value string)     { v(in).setAttribute(name, value) }
func (in Input) RemoveAttribute(name string)         { v(in).removeAttribute(name) }
func (in Input) ToggleAttribute(name string) bool    { return v(in).toggleAttribute(name) }
func (in Input) HasAttribute(name string) bool       { return v(in).hasAttribute(name) }
func (in Input) Closest(selector string) dom.Element { return v(in).closest(selector) }
func (in Input) Matches(selector string) bool        { return v(in).matches(selector) }

func (in Input) SetInnerHTML(s string)      { v(in).setInnerHTML(s) }
func (in Input) InnerHTML() string          { return v(in).innerHTML() }
func (in Input) SetOuterHTML(s string)      { v(in).setOuterHTML(s) }
func (in Input) OuterHTML() string          { return v(in).outerHTML() }
func (in Input) SetInnerText(s string)      { v(in).setInnerText(s) }
func (in Input) InnerText() string          { return v(in).innerText() }
func (in Input) SetTextContent(text string) { v(in).setTextContent(text) }

func (in Input) Length() int { return v(in).length() }

func (in Input) InsertAdjacentHTML(pos dom.InsertAdjacentHTMLPosition, text string) {
	v(in).insertAdjacentHTML(pos, text)
}

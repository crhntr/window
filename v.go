//go:build js

package window

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type v js.Value

func (val v) nodeType() dom.NodeType  { return nodeType(js.Value(val).Get("nodeType")) }
func (val v) isConnected() bool       { return js.Value(val).Get("isConnected").Bool() }
func (val v) ownerDocument() dom.Node { return valueToNode(js.Value(val).Get("ownerDocument")) }
func (val v) getRootNode(composed bool) dom.Node {
	return valueToNode(js.Value(val).Call("getRootNode", getRootNodeOptions(composed)))
}
func (val v) parentNode() dom.Node       { return valueToNode(js.Value(val).Get("parentNode")) }
func (val v) parentElement() dom.Element { return Element(js.Value(val).Call("parentElement")) }
func (val v) hasChildNodes() bool        { return js.Value(val).Call("hasChildNodes").Bool() }
func (val v) childNodes() dom.NodeList   { return NodeList(js.Value(val).Call("childNodes")) }
func (val v) firstChild() dom.Node       { return valueToNode(js.Value(val).Call("firstChild")) }
func (val v) lastChild() dom.Node        { return valueToNode(js.Value(val).Call("lastChild")) }
func (val v) previousSibling() dom.Node {
	return valueToNode(js.Value(val).Call("previousSibling"))
}
func (val v) nextSibling() dom.Node { return valueToNode(js.Value(val).Call("nextSibling")) }
func (val v) nodeValue() string     { return js.Value(val).Call("nodeValue").String() }
func (val v) textContent() string   { return js.Value(val).Call("textContent").String() }
func (val v) normalize()            { js.Value(val).Call("normalize") }
func (val v) cloneNode(deep bool) dom.Node {
	return valueToNode(js.Value(val).Call("cloneNode", deep))
}
func (val v) isEqualNode(other dom.Node) bool {
	return js.Value(val).Call("isEqualNode", nodeToValue(other)).Bool()
}
func (val v) isSameNode(other dom.Node) bool {
	return js.Value(val).Call("isSameNode", nodeToValue(other)).Bool()
}
func (val v) compareDocumentPosition() dom.DocumentPosition {
	return dom.DocumentPosition(js.Value(val).Call("compareDocumentPosition").Int())
}
func (val v) contains(other dom.Node) bool {
	return js.Value(val).Call("contains", nodeToValue(other)).Bool()
}
func (val v) insertBefore(node, child dom.Node) dom.Node {
	return valueToNode(js.Value(val).Call("insertBefore", node, child))
}
func (val v) appendChild(node dom.Node) dom.Node {
	return valueToNode(js.Value(val).Call("appendChild", node))
}
func (val v) replaceChild(node, child dom.Node) dom.Node {
	return valueToNode(js.Value(val).Call("replaceChild", node, child))
}
func (val v) removeChild(node dom.Node) dom.Node {
	return valueToNode(js.Value(val).Call("removeChild", node))
}
func (val v) children() dom.NodeList { return NodeList(js.Value(val).Call("children")) }
func (val v) firstElementChild() dom.Element {
	return Element(js.Value(val).Call("firstElementChild"))
}
func (val v) lastElementChild() dom.Element {
	return Element(js.Value(val).Call("lastElementChild"))
}
func (val v) childElementCount() int { return js.Value(val).Get("childElementCount").Int() }
func (val v) prepend(nodes []dom.Node) dom.Node {
	return Element(js.Value(val).Call("prepend", nodesToEmptyInterfaceSlice(nodes)...))
}
func (val v) append(nodes []dom.Node) dom.Node {
	return document(js.Value(val).Call("append", nodesToEmptyInterfaceSlice(nodes)...))
}
func (val v) replaceChildren(nodes []dom.Node) dom.Node {
	return document(js.Value(val).Call("replaceChildren", nodesToEmptyInterfaceSlice(nodes)...))
}
func (val v) getElementsByTagName(name string) dom.ElementCollection {
	return ElementCollection(js.Value(val).Call("getElementsByTagName", name))
}
func (val v) getElementsByTagNameNS(namespace, localName string) dom.ElementCollection {
	return ElementCollection(js.Value(val).Call("getElementsByTagNameNS", namespace, localName))
}
func (val v) getElementsByClassName(name string) dom.ElementCollection {
	return ElementCollection(js.Value(val).Call("getElementsByClassName", name))
}
func (val v) querySelector(query string) dom.Element {
	return Element(js.Value(val).Call("querySelector", query))
}
func (val v) querySelectorAll(query string) dom.NodeList {
	return NodeList(js.Value(val).Call("querySelectorAll", query))
}
func (val v) createElement(localName string) dom.Element {
	return Element(js.Value(val).Call("createElement", localName))
}
func (val v) createElementIs(localName, is string) dom.Element {
	return Element(js.Value(val).Call("createElementIs", localName, js.ValueOf(map[string]interface{}{"is": is})))
}
func (val v) createElementNS(namespace, localName string) dom.Element {
	return Element(js.Value(val).Call("createElementNS", namespace, localName))
}
func (val v) createElementNSIS(namespace, localName, is string) dom.Element {
	return Element(js.Value(val).Call("createElementNS", namespace, localName, js.ValueOf(map[string]interface{}{"is": is})))
}
func (val v) createTextNode(text string) dom.Text {
	return Text(js.Value(val).Call("createTextNode", text))
}
func (val v) tagName() string   { return js.Value(val).Get("tagName").String() }
func (val v) iD() string        { return js.Value(val).Get("id").String() }
func (val v) className() string { return js.Value(val).Get("className").String() }
func (val v) getAttribute(name string) string {
	return js.Value(val).Call("getAttribute", name).String()
}
func (val v) getAttributeNS(namespace, name string) string {
	return js.Value(val).Call("getAttributeNS", namespace, name).String()
}
func (val v) setAttribute(name, value string) { js.Value(val).Call("setAttribute", name, value) }
func (val v) setAttributeNS(namespace, name, value string) {
	js.Value(val).Call("setAttributeNS", namespace, name, value)
}
func (val v) removeAttribute(name string) { js.Value(val).Call("removeAttribute", name) }
func (val v) removeAttributeNS(namespace, name string) {
	js.Value(val).Call("removeAttributeNS", namespace, name)
}
func (val v) toggleAttribute(name string) bool {
	return js.Value(val).Call("toggleAttribute", name, true).Bool()
}
func (val v) hasAttribute(name string) bool {
	return js.Value(val).Call("hasAttribute", name).Bool()
}
func (val v) hasAttributeNS(namespace, name string) bool {
	return js.Value(val).Call("setAttribute", namespace, name).Bool()
}
func (val v) closest(selector string) dom.Element {
	return Element(js.Value(val).Call("closest", selector))
}
func (val v) matches(selector string) bool {
	return js.Value(val).Call("setAttribute", selector).Bool()
}

func (val v) setInnerHTML(s string) { js.Value(val).Set("innerHTML", s) }
func (val v) innerHTML() string     { return js.Value(val).Get("innerHTML").String() }
func (val v) setOuterHTML(s string) { js.Value(val).Set("outerHTML", s) }
func (val v) outerHTML() string     { return js.Value(val).Get("outerHTML").String() }
func (val v) setInnerText(s string) { js.Value(val).Set("innerText", s) }
func (val v) innerText() string     { return js.Value(val).Get("innerText").String() }

func (val v) data() string         { return js.Value(val).Get("data").String() }
func (val v) split(n int) dom.Text { return Text(js.Value(val).Call("split", n)) }
func (val v) wholeText() string    { return js.Value(val).Call("wholeText").String() }

func (val v) createDocumentFragment() DocumentFragment {
	return DocumentFragment(js.Value(val).Call("createDocumentFragment"))
}

func (val v) length() int                { return js.Value(val).Length() }
func (val v) item(index int) dom.Element { return Element(js.Value(val).Call("item", index)) }
func (val v) namedItem(name string) dom.Element {
	return Element(js.Value(val).Call("namedItem", name))
}

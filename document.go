// +build js,wasm

package window

import (
	"bytes"
	"fmt"
	"strings"
	"syscall/js"
)

type document byte

const (
	Document document = 0
)

var doc = win.Get("document")

func (document document) JSValue() js.Value { return doc }

func (document document) Get(key string) js.Value           { return doc.Get(key) }
func (document document) Equal(w js.Value) bool             { return doc.Equal(w) }
func (document document) Set(key string, value interface{}) { document.Set(key, value) }
func (document document) Call(m string, args ...interface{}) js.Value {
	return doc.Call(m, args...)
}
func (document document) Type() js.Type              { return doc.Type() }
func (document document) Truthy() bool               { return doc.Truthy() }
func (document document) IsNull() bool               { return doc.IsNull() }
func (document document) IsUndefined() bool          { return doc.IsUndefined() }
func (document document) InstanceOf(t js.Value) bool { return doc.InstanceOf(t) }

func (document document) Append(node ...Node)  { appendNodes(doc, node) }
func (document document) Prepend(node ...Node) { prepend(doc, node) }
func (document document) QuerySelector(query string, a ...interface{}) Element {
	return querySelector(doc, query, a)
}
func (document document) QuerySelectorAll(query string, a ...interface{}) NodeList {
	return querySelectorAll(doc, query, a)
}
func (document document) ReplaceChildren(node ...Node) { replaceChildren(doc, node) }
func (document document) ChildElementCount() int       { return childElementCount(doc) }
func (document document) Children() HTMLCollection     { return children(doc) }
func (document document) FirstElementChild() Element   { return firstElementChild(doc) }
func (document document) LastElementChild() Element    { return lastElementChild(doc) }

func (document document) Node() Node            { return document }
func (document document) NodeType() NodeType    { return nodeType(document) }
func (document document) FirstChild() ChildNode { return firstChild(document) }
func (document document) IsConnected() bool     { return isConnected(document) }
func (document document) LastChild() ChildNode  { return lastChild(document) }
func (document document) ChildNodes() NodeList  { return childNodes(document) }
func (document document) AppendChild(child NodeWrapper) ChildNode {
	return appendChild(document, child)
}
func (document document) RemoveChild(child NodeWrapper) Node { return removeChild(document, child) }
func (document document) ReplaceChild(newChild, oldChild NodeWrapper) Node {
	return replaceChild(document, newChild, oldChild)
}
func (document document) Contains(child NodeWrapper) bool { return contains(document, child) }
func (document document) NextSibling() ChildNode          { return nextSibling(document) }
func (document document) PreviousSibling() ChildNode      { return previousSibling(document) }
func (document document) ParentNode() Node                { return parentNode(document) }
func (document document) ParentElement() Element          { return parentElement(document) }
func (document document) CloneNode(isDeep bool) Node      { return cloneNode(document, isDeep) }
func (document document) HasChildNodes() bool             { return hasChildNodes(document) }
func (document document) Normalize()                      { normalize(document) }

func (document document) AddEventListener(eventName string, listener EventListener) func() {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		listener.HandleEvent(Event(args[0]))
		return nil
	})

	win.Call("addEventListener", eventName, fn)

	return func() {
		defer fn.Release()
		win.Call("removeEventListener", eventName, fn)
	}
}

func (document document) AddEventListenerFunc(eventName string, listener EventListenerFunc) func() {
	return document.AddEventListener(eventName, listener)
}

func (document document) AddEventListenerChannel(eventName string, c chan Event) func() {
	return document.AddEventListenerFunc(eventName, func(event Event) {
		c <- event
	})
}

func (document document) NewElementFromTemplate(name string, data interface{}) (Element, error) {
	buf := bytes.NewBuffer(nil)
	err := templates.ExecuteTemplate(buf, name, data)
	if err != nil {
		return Element(js.Null()), err
	}

	div := document.Call("createElement", "div")
	div.Set("innerHTML", strings.TrimSpace(buf.String()))

	v := div.Get("firstChild")
	if !v.Truthy() {
		return Element(js.Null()), fmt.Errorf("could not get created element")
	}

	return Element(v), nil
}

// NewElement creates a new element from the format string and a call to Sprintf
// callers should ensure the element was created successfully. Sprintf sanitizes all parameters
// but assumes the format string is safe html.
//
//  el = Document.NewElement(`<div class=%[2]q>Hello, %[1]s!</div>`, "world", "greeting")
//  if !el.Truthy() {
//    // handle error
//  }
//
func (document document) NewElement(format string, a ...interface{}) Element {
	tmpDiv := document.Call("createElement", "div")
	tmpDiv.Set("innerHTML", Sprintf(format, a...))
	return Element(tmpDiv.Get("firstChild"))
}

func (document document) CreateTextNode(text string) Text {
	return Text(document.Call("createTextNode", text))
}

func (document document) CreateElement(tagName string) Element {
	return Element(document.Call("createElement", tagName))
}

func (document document) GetElementByID(id string) Element {
	v := document.Call("getElementById", id)
	if !v.Truthy() {
		return Element(js.Null())
	}
	return Element(v)
}

func (document document) IsSameNode(node Node) bool { return isSameNode(document, node) }

func (document document) IsEqualNode(node Node) bool { return isEqualNode(document, node) }

func (document document) AsElement() Element { return nodeAsElement(document) }
func (document document) AsText() Text       { return nodeAsText(document) }

func (document document) Body() Element { return Element(document.Get("body")) }
func (document document) Head() Element { return Element(document.Get("head")) }

func (document document) Title() string         { return document.Get("title").String() }
func (document document) SetTitle(title string) { document.Set("title", title) }

type DocumentFragment js.Value

func (doc DocumentFragment) IsNull() bool { return doc.JSValue().IsNull() }

func (doc DocumentFragment) Append(nodes ...Node)  { appendNodes(doc, nodes) }
func (doc DocumentFragment) Prepend(nodes ...Node) { prepend(doc, nodes) }
func (doc DocumentFragment) QuerySelector(query string, a ...interface{}) Element {
	return querySelector(doc, query, a)
}
func (doc DocumentFragment) QuerySelectorAll(query string, a ...interface{}) NodeList {
	return querySelectorAll(doc, query, a)
}
func (doc DocumentFragment) ReplaceChildren(node ...Node) { replaceChildren(doc, node) }

func (doc DocumentFragment) Node() Node                              { return doc }
func (doc DocumentFragment) NodeType() NodeType                      { return nodeType(doc) }
func (doc DocumentFragment) FirstChild() ChildNode                   { return firstChild(doc) }
func (doc DocumentFragment) IsConnected() bool                       { return isConnected(doc) }
func (doc DocumentFragment) LastChild() ChildNode                    { return lastChild(doc) }
func (doc DocumentFragment) ChildNodes() NodeList                    { return childNodes(doc) }
func (doc DocumentFragment) AppendChild(child NodeWrapper) ChildNode { return appendChild(doc, child) }
func (doc DocumentFragment) RemoveChild(child NodeWrapper) Node      { return removeChild(doc, child) }
func (doc DocumentFragment) ReplaceChild(newChild, oldChild NodeWrapper) Node {
	return replaceChild(doc, newChild, oldChild)
}
func (doc DocumentFragment) Contains(child NodeWrapper) bool { return contains(doc, child) }
func (doc DocumentFragment) NextSibling() ChildNode          { return nextSibling(doc) }
func (doc DocumentFragment) PreviousSibling() ChildNode      { return previousSibling(doc) }
func (doc DocumentFragment) ParentNode() Node                { return parentNode(doc) }
func (doc DocumentFragment) ParentElement() Element          { return parentElement(doc) }
func (doc DocumentFragment) CloneNode(isDeep bool) Node      { return cloneNode(doc, isDeep) }
func (doc DocumentFragment) HasChildNodes() bool             { return hasChildNodes(doc) }
func (doc DocumentFragment) Normalize()                      { normalize(doc) }

func (document document) CreateDocumentFragment(format string, a ...interface{}) DocumentFragment {
	content := format

	if len(a) > 0 {
		content = Sprintf(format, a...)
	}

	return DocumentFragment(document.Call("createDocumentFragment", content))
}

func (doc DocumentFragment) Children() HTMLCollection {
	return HTMLCollection(doc.JSValue().Get("children"))
}

func (doc DocumentFragment) FirstElementChild() Element {
	return Element(doc.JSValue().Get("firstElementChild"))
}

func (doc DocumentFragment) LastElementChild() Element {
	return Element(doc.JSValue().Get("lastElementChild"))
}

func (doc DocumentFragment) ChildElementCount() int {
	return doc.JSValue().Get("childElementCount").Int()
}

func (doc DocumentFragment) JSValue() js.Value { return js.Value(doc) }

func (doc DocumentFragment) IsSameNode(node Node) bool { return isSameNode(doc, node) }

func (doc DocumentFragment) IsEqualNode(node Node) bool { return isEqualNode(doc, node) }

func (doc DocumentFragment) AsElement() Element { return nodeAsElement(doc) }
func (doc DocumentFragment) AsText() Text       { return nodeAsText(doc) }

type HTMLCollection js.Value

func (col HTMLCollection) Length() int {
	return col.JSValue().Length()
}

func (col HTMLCollection) Index(i int) js.Value {
	return col.JSValue().Index(i)
}

func (col HTMLCollection) Item(i int) js.Value {
	return col.JSValue().Call("item", i)
}

func (col HTMLCollection) NamedItem(nameOrID string) js.Value {
	return col.JSValue().Call("namedItem", nameOrID)
}

func (col HTMLCollection) JSValue() js.Value { return js.Value(col) }

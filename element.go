//go:build js && wasm
// +build js,wasm

package window

import (
	"fmt"
	"syscall/js"
)

type Element js.Value

func (el Element) JSValue() js.Value { return js.Value(el) }

func (el Element) Get(key string) js.Value                     { return js.Value(el).Get(key) }
func (el Element) Equal(w js.Value) bool                       { return js.Value(el).Equal(w) }
func (el Element) Set(key string, value interface{})           { js.Value(el).Set(key, value) }
func (el Element) Call(m string, args ...interface{}) js.Value { return js.Value(el).Call(m, args...) }
func (el Element) Type() js.Type                               { return js.Value(el).Type() }
func (el Element) Truthy() bool                                { return js.Value(el).Truthy() }
func (el Element) IsNull() bool                                { return js.Value(el).IsNull() }
func (el Element) IsUndefined() bool                           { return js.Value(el).IsUndefined() }
func (el Element) InstanceOf(t js.Value) bool                  { return js.Value(el).InstanceOf(t) }

func (el Element) ChildElementCount() int     { return childElementCount(el) }
func (el Element) Children() HTMLCollection   { return children(el) }
func (el Element) FirstElementChild() Element { return firstElementChild(el) }
func (el Element) LastElementChild() Element  { return lastElementChild(el) }
func (el Element) Append(nodes ...Node)       { appendNodes(el, nodes) }
func (el Element) Prepend(nodes ...Node)      { prepend(el, nodes) }
func (el Element) QuerySelector(query string, a ...interface{}) Element {
	return querySelector(el, query, a)
}
func (el Element) QuerySelectorAll(query string, a ...interface{}) NodeList {
	return querySelectorAll(el, query, a)
}
func (el Element) ReplaceChildren(nodes ...Node) { replaceChildren(el, nodes) }

func (el Element) Node() Node                              { return el }
func (el Element) NodeType() NodeType                      { return nodeType(el) }
func (el Element) FirstChild() ChildNode                   { return firstChild(el) }
func (el Element) IsConnected() bool                       { return isConnected(el) }
func (el Element) LastChild() ChildNode                    { return lastChild(el) }
func (el Element) ChildNodes() NodeList                    { return childNodes(el) }
func (el Element) AppendChild(child NodeWrapper) ChildNode { return appendChild(el, child) }
func (el Element) RemoveChild(child NodeWrapper) Node      { return removeChild(el, child) }
func (el Element) ReplaceChild(newChild, oldChild NodeWrapper) Node {
	return replaceChild(el, newChild, oldChild)
}
func (el Element) Contains(child NodeWrapper) bool { return contains(el, child) }
func (el Element) NextSibling() ChildNode          { return nextSibling(el) }
func (el Element) PreviousSibling() ChildNode      { return previousSibling(el) }
func (el Element) ParentNode() Node                { return parentNode(el) }
func (el Element) ParentElement() Element          { return parentElement(el) }
func (el Element) CloneNode(isDeep bool) Node      { return cloneNode(el, isDeep) }
func (el Element) HasChildNodes() bool             { return hasChildNodes(el) }
func (el Element) Normalize()                      { normalize(el) }

func (el Element) Attribute(key string) string {
	attrVal := el.JSValue().Call("getAttribute", key)
	if attrVal.IsNull() {
		return ""
	}
	return attrVal.String()
}

func (el Element) SetAttribute(key, val string, args ...interface{}) {
	el.JSValue().Call("setAttribute", key, fmt.Sprintf(val, args...))
}

func (el Element) RemoveAttribute(key string) {
	if el.JSValue().IsNull() {
		return
	}
	el.JSValue().Call("removeAttribute", key)
}

func (el Element) AddClass(class string) {
	if list := js.Value(el).Get("classList"); list.Truthy() {
		list.Call("add", class)
	}
}

func (el Element) RemoveClass(class string) {
	el.JSValue().Get("classList").Call("remove", class)
}

func (el Element) ReplaceClass(old, new string) {
	el.JSValue().Get("classList").Call("replace", old, new)
}

func (el Element) ToggleClass(class string) {
	el.JSValue().Get("classList").Call("toggle", class)
}

func (el Element) HasClass(class string) bool {
	return el.JSValue().Get("classList").Call("contains", class).Bool()
}

func (el Element) Clone() Element {
	return Element(el.JSValue().Call("cloneNode", true))
}

func (el Element) InnerText() string { return el.JSValue().Get("innerText").String() }

func (el Element) SetInnerText(format string, a ...interface{}) {
	el.JSValue().Set("innerText", fmt.Sprintf(format, a...))
}

func (el Element) InnerHTML() string { return el.JSValue().Get("innerHTML").String() }

func (el Element) SetInnerHTML(str string) { el.JSValue().Set("innerHTML", str) }

func (el Element) SetInnerHTMLf(format string, vs ...interface{}) {
	el.JSValue().Set("innerHTML", Sprintf(format, vs...))
}

func (el Element) OuterHTML() string { return el.JSValue().Get("outerHTML").String() }

func (el Element) SetOuterHTML(str string) { el.JSValue().Set("outerHTML", str) }

func (el Element) SetOuterHTMLf(format string, vs ...interface{}) {
	el.JSValue().Set("outerHTML", Sprintf(format, vs...))
}

func (el Element) ReplaceElement(newEl Element) { el.Parent().ReplaceChild(newEl, el) }

func (el Element) Matches(query string) bool { return el.JSValue().Call("matches", query).Bool() }

func (el Element) Closest(query string) Element {
	return Element(el.JSValue().Call("closest", query))
}

func (el Element) HasAncestor(query string) bool {
	return el.Closest(query).JSValue().Truthy()
}

func (el Element) AddEventListener(eventName string, listener EventListener) func() {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		listener.HandleEvent(Event(args[0]))
		return nil
	})

	el.JSValue().Call("addEventListener", eventName, fn)

	return func() {
		defer fn.Release()

		if !el.JSValue().Truthy() {
			return
		}

		el.JSValue().Call("removeEventListener", eventName, fn)
	}
}

func (el Element) AddEventListenerFunc(eventName string, listener EventListenerFunc) func() {
	return el.AddEventListener(eventName, listener)
}

func (el Element) AddEventListenerChannel(eventName string, c chan Event) func() {
	return el.AddEventListenerFunc(eventName, func(event Event) {
		c <- event
	})
}

func (el Element) Tag() string {
	return el.JSValue().Get("tagName").String()
}

func (el Element) Parent() Element {
	if el.JSValue().Type() == js.TypeUndefined || el.JSValue().Type() == js.TypeNull {
		return Element(js.Null())
	}
	return Element(el.JSValue().Get("parentElement"))
}

func (el Element) Log() {
	js.Global().Get("console").Call("log", el)
}

func toInterfaceSlice(dst []interface{}, slice []js.Wrapper) {
	for i := range slice {
		dst[i] = slice[i]
	}
}
func (el Element) InsertBefore(element js.Wrapper) {

	el.Parent().JSValue().Call("insertBefore", element, el)
}

func (el Element) InsertAfter(element Node) {
	if next := el.NextSiblingElement(); next.JSValue().Truthy() {
		next.InsertBefore(element)
		return
	}
	el.Parent().AppendChild(element)
}

// InsertHTMLBefore calls insertAdjacentHTML with beforebegin
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) InsertHTMLBefore(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentHTML", "beforebegin", Sprintf(format, vs...))
}

// PrependHTML calls insertAdjacentHTML with afterbegin
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) PrependHTML(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentHTML", "afterbegin", Sprintf(format, vs...))
}

// AppendHTML calls insertAdjacentHTML with beforeend
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) AppendHTML(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentHTML", "beforeend", Sprintf(format, vs...))
}

// InsertHTMLAfter calls insertAdjacentHTML with afterend
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) InsertHTMLAfter(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentHTML", "afterend", Sprintf(format, vs...))
}

// InsertTextBefore calls insertAdjacentText with beforebegin
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) InsertTextBefore(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentText", "beforebegin", Sprintf(format, vs...))
}

// PrependText calls insertAdjacentText with afterbegin
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) PrependText(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentText", "afterbegin", Sprintf(format, vs...))
}

// AppendText calls insertAdjacentText with beforeend
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) AppendText(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentText", "beforeend", Sprintf(format, vs...))
}

// InsertTextAfter calls insertAdjacentText with afterend
// it is not safe to use with user input as it calls fmt.Sprintf
func (el Element) InsertTextAfter(format string, vs ...interface{}) {
	el.JSValue().Call("insertAdjacentText", "afterend", Sprintf(format, vs...))
}

// IndexOf returns the index at which a given element can be found in the array, or -1 if it is not present.
func (el Element) IndexOf(child js.Wrapper) int {
	return js.Global().Get("Array").Get("prototype").Get("indexOf").Call("call", el.JSValue().Get("children"), child).Int()
}

// NextSiblingElement wraps nextElementSibling
func (el Element) NextSiblingElement() Element {
	return Element(el.JSValue().Get("nextElementSibling"))
}

func (el Element) Style() js.Value { return el.JSValue().Get("style") }

func (el Element) IsSameNode(node Node) bool { return isSameNode(el, node) }

func (el Element) IsEqualNode(node Node) bool { return isEqualNode(el, node) }

func (el Element) Remove()                  { childNodeRemove(el) }
func (el Element) Before(node ...Node)      { childNodeBefore(el, node) }
func (el Element) After(node ...Node)       { childNodeAfter(el, node) }
func (el Element) ReplaceWith(node ...Node) { childNodeReplaceWith(el, node) }

func (el Element) AsElement() Element { return nodeAsElement(el) }
func (el Element) AsText() Text       { return nodeAsText(el) }

func (el Element) DispatchEvent(e Event) {
	el.Call("dispatchEvent", e)
}

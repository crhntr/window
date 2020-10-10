// +build js,wasm

package window

import (
	"fmt"
	"syscall/js"
)

type Element js.Value

func (el Element) Get(key string) js.Value                     { return js.Value(el).Get(key) }
func (el Element) Set(key string, value interface{})           { js.Value(el).Set(key, value) }
func (el Element) Call(m string, args ...interface{}) js.Value { return js.Value(el).Call(m, args...) }
func (el Element) JSValue() js.Value                           { return js.Value(el) }
func (el Element) Type() js.Type                               { return js.Value(el).Type() }
func (el Element) Truthy() bool                                { return js.Value(el).Truthy() }
func (el Element) IsNull() bool                                { return js.Value(el).IsNull() }
func (el Element) IsUndefined() bool                           { return js.Value(el).IsUndefined() }
func (el Element) InstanceOf(t js.Value) bool                  { return js.Value(el).InstanceOf(t) }

func (el Element) Attribute(key string) string {
	return el.Call("getAttribute", key).String()
}

func (el Element) SetAttribute(key, val string, args ...interface{}) {
	el.Call("setAttribute", key, fmt.Sprintf(val, args...))
}

func (el Element) AddClass(class string) {
	if list := js.Value(el).Get("classList"); list.Truthy() {
		list.Call("add", class)
	}
}

func (el Element) RemoveClass(class string) {
	el.Get("classList").Call("remove", class)
}

func (el Element) ReplaceClass(old, new string) {
	el.Get("classList").Call("replace", old, new)
}

func (el Element) ToggleClass(class string) {
	el.Get("classList").Call("toggle", class)
}

func (el Element) HasClass(class string) bool {
	return el.Get("classList").Call("contains", class).Bool()
}

func (el Element) Clone() Element {
	return Element(el.Call("cloneNode", true))
}

func (el Element) InnerHTML() string {
	return el.Get("innerHTML").String()
}

func (el Element) SetInnerHTML(str string) {
	el.Set("innerHTML", str)
}

func (el Element) SetInnerHTMLf(format string, vs ...interface{}) {
	el.Set("innerHTML", fmt.Sprintf(format, vs...))
}

func (el Element) ChildCount() int {
	return el.Get("childElementCount").Int()
}

func (el Element) AppendChild(child js.Wrapper) {
	el.Call("appendChild", child)
}

// Prepend inserts elements before the first child of the Element
// https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/prepend
func (el Element) Prepend(elements ...js.Wrapper) {
	if len(elements) == 0 {
		return
	}
	args := make([]interface{}, len(elements))
	toInterfaceSlice(args, elements)
	el.Call("prepend", args...)
}

// ReplaceChild should only be called if the caller is sure
// existingChild exists.
func (el Element) ReplaceChild(newChild, existingChild js.Wrapper) {
	el.Call("replaceChild", newChild, existingChild)
}

func (el Element) Matches(query string) bool {
	return el.Call("matches", query).Bool()
}

func (el Element) Closest(query string) Element {
	return Element(el.Call("closest", query))
}

func (el Element) HasAncestor(query string) bool {
	return el.Closest(query).Truthy()
}

func (el Element) QuerySelector(query string, args ...interface{}) Element {
	query = fmt.Sprintf(query, args...)

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(js.Error); ok {
				panic(fmt.Errorf("query selector failed for %q: %w", query, err))
			}
		}
	}()

	return Element(el.Call("querySelector", query))
}

func (el Element) QuerySelectorAll(query string, args ...interface{}) []Element {
	var elements []Element

	query = fmt.Sprintf(query, args...)

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(js.Error); ok {
				panic(fmt.Errorf("query selector failed for %q: %w", query, err))
			}
		}
	}()

	matches := el.Call("querySelectorAll", query)

	if !matches.Truthy() {
		return nil
	}

	for index := 0; index < matches.Length(); index++ {
		elements = append(elements, Element(matches.Index(index)))
	}

	return elements
}

func (el Element) AddEventListener(eventName string, listener EventListener) func() {
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

func (el Element) AddEventListenerFunc(eventName string, listener EventListenerFunc) func() {
	return el.AddEventListener(eventName, listener)
}

func (el Element) AddEventListenerChannel(eventName string, c chan Event) func() {
	return el.AddEventListenerFunc(eventName, func(event Event) {
		c <- event
	})
}

func (el Element) Tag() string {
	return el.Get("tagName").String()
}

func (el Element) Remove() {
	if el.Type() == js.TypeUndefined || el.Type() == js.TypeNull {
		return
	}
	el.Call("remove")
}

func (el Element) Parent() Element {
	if el.Type() == js.TypeUndefined || el.Type() == js.TypeNull {
		return Element(js.Null())
	}
	return Element(el.Get("parentElement"))
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
	el.Parent().Call("insertBefore", element, el)
}

func (el Element) InsertHTMLBefore(format string, vs ...interface{}) {
	el.Call("insertAdjacentHTML", "beforebegin", fmt.Sprintf(format, vs...))
}

func (el Element) PrependHTML(format string, vs ...interface{}) {
	el.Call("insertAdjacentHTML", "afterbegin", fmt.Sprintf(format, vs...))
}

func (el Element) AppendHTML(format string, vs ...interface{}) {
	el.Call("insertAdjacentHTML", "beforeend", fmt.Sprintf(format, vs...))
}

func (el Element) InsertAfter(element js.Wrapper) {
	if next := el.NextSiblingElement(); next.Truthy() {
		next.InsertBefore(element)
		return
	}
	el.Parent().AppendChild(element)
}

func (el Element) InsertHTMLAfter(format string, vs ...interface{}) {
	el.Call("insertAdjacentHTML", "afterend", fmt.Sprintf(format, vs...))
}

func (el Element) InsertTextBefore(format string, vs ...interface{}) {
	el.Call("insertAdjacentText", "beforebegin", fmt.Sprintf(format, vs...))
}

func (el Element) PrependText(format string, vs ...interface{}) {
	el.Call("insertAdjacentText", "afterbegin", fmt.Sprintf(format, vs...))
}

func (el Element) AppendText(format string, vs ...interface{}) {
	el.Call("insertAdjacentText", "beforeend", fmt.Sprintf(format, vs...))
}

func (el Element) InsertTextAfter(format string, vs ...interface{}) {
	el.Call("insertAdjacentText", "afterend", fmt.Sprintf(format, vs...))
}

// IndexOf returns the index at which a given element can be found in the array, or -1 if it is not present.
func (el Element) IndexOf(child js.Wrapper) int {
	return js.Global().Get("Array").Get("prototype").Get("indexOf").Call("call", el.Get("children"), child).Int()
}

// NextSiblingElement wraps nextElementSibling
func (el Element) NextSiblingElement() Element {
	return Element(el.Get("nextElementSibling"))
}

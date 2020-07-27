// +build js,wasm

package window

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"syscall/js"
)

type Element js.Value

func (document document) NewElementFromTemplate(tmp *template.Template, name string, data interface{}) (Element, error) {
	buf := bytes.NewBuffer(nil)
	err := tmp.ExecuteTemplate(buf, name, data)
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

func (document document) NewElement(format string, vs ...interface{}) (Element, error) {
	tmp := document.Call("createElement", "div")
	tmp.Set("innerHTML", strings.TrimSpace(fmt.Sprintf(format, vs...)))

	v := tmp.Get("firstChild")
	if !v.Truthy() {
		return Element(js.Null()), fmt.Errorf("could not get created element")
	}

	return Element(v), nil
}

func (document document) GetElementByID(id string) Element {
	v := document.Call("getElementById", id)
	if !v.Truthy() {
		return Element(js.Null())
	}
	return Element(v)
}

func (document document) QuerySelector(query string, args ...interface{}) Element {
	query = fmt.Sprintf(query, args...)

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(js.Error); ok {
				panic(fmt.Errorf("query selector failed for %q: %w", query, err))
			}
		}
	}()

	return Element(document.Call("querySelector", query))
}

func (document document) QuerySelectorAll(query string, args ...interface{}) []Element {
	query = fmt.Sprintf(query, args...)

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(js.Error); ok {
				panic(fmt.Errorf("query selector failed for %q: %w", query, err))
			}
		}
	}()

	matches := document.Call("querySelectorAll", query)

	if !matches.Truthy() {
		return nil
	}

	length := matches.Length()

	elements := make([]Element, length)

	for index := 0; index < length; index++ {
		elements[index] = Element(matches.Index(index))
	}

	return elements
}

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

func (el Element) AppendChild(child Element) {
	el.Call("appendChild", child)
}

func (el Element) Matches(query string) bool {
	return el.Call("matches", query).Bool()
}

func (el Element) Closest(query string) Element {
	return Element(el.Call("closest", query))
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

func (el Element) AddEventListener(eventName string, listener EventListener) {
	el.Call("addEventListener", eventName, js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		listener.HandleEvent(Event(args[0]))
		return nil
	}))
}

func (fn EventListenerFunc) HandleEvent(event Event) { fn(event) }

func (el Element) AddEventListenerFunc(eventName string, listener EventListenerFunc) {
	el.AddEventListener(eventName, listener)
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

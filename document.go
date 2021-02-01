// +build js,wasm

package window

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"syscall/js"
)

type document byte

const (
	Document document = 0
)

var doc = js.Global().Get("document")

func (document document) JSValue() js.Value                           { return doc }
func (document document) Get(key string) js.Value                     { return doc.Get(key) }
func (document document) Set(key string, value interface{})           { doc.Set(key, value) }
func (document document) Call(m string, args ...interface{}) js.Value { return doc.Call(m, args...) }
func (document document) Type() js.Type                               { return doc.Type() }

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

// NewElement creates a new element from the format string and a call to fmt.Sprintf
// callers should ensure the element was created successfully
//
//  el = Document.NewElement(`<div class=%[2]q>Hello, %[1]s!</div>`, struct{
//  	Class, Name string
//  }{
// 		Class: "greeting",
//		Name: "world",
//  })
//  if !el.Truthy() {
//    // handle error
//  }
//
func (document document) NewElement(templateHTML string, data interface{}) Element {
	t, err := template.New("").Parse(templateHTML)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	tmpDiv := document.Call("createElement", "div")
	tmpDiv.Set("innerHTML", buf.String())
	return Element(tmpDiv.Get("firstChild"))
}

func (document document) CreateTextNode(text string) js.Value {
	return document.Call("createTextNode", text)
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

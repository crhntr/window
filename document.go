// +build js,wasm

package window

import (
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

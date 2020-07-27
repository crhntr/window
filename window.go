// +build js,wasm

package window

import (
	"net/url"
	"syscall/js"
)

var win = js.Global()

func JSValue() js.Value                           { return win }
func Get(key string) js.Value                     { return win.Get(key) }
func Set(key string, value interface{})           { win.Set(key, value) }
func Call(m string, args ...interface{}) js.Value { return win.Call(m, args...) }
func Type() js.Type                               { return win.Type() }

func AddEventListener(eventName string, listener EventListener) {
	win.Call("addEventListener", eventName, js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		listener.HandleEvent(Event(args[0]))
		return nil
	}))
}

func AddEventListenerFunc(eventName string, listener EventListenerFunc) {
	AddEventListener(eventName, listener)
}

func AddEventListenerChannel(eventName string, c chan Event) {
	AddEventListenerFunc(eventName, func(event Event) {
		c <- event
	})
}

func URL() (*url.URL, error) {
	return url.Parse(Get("location").Get("href").String())
}

func Body() Element {
	return Element(Document.Call("getElementsByTagName", "body").Index(0))
}

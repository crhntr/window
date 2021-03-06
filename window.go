// +build js,wasm

package window

import (
	"net/url"
	"syscall/js"
)

var win = js.Global()

func JSValue() js.Value                           { return win }
func Get(key string) js.Value                     { return win.Get(key) }
func Delete(key string)                           { win.Delete(key) }
func Set(key string, value interface{})           { win.Set(key, value) }
func Call(m string, args ...interface{}) js.Value { return win.Call(m, args...) }
func Type() js.Type                               { return win.Type() }

// AddEventListener should be called once per eventName. It returns a wrapper around "removeEventListener".
func AddEventListener(eventName string, listener EventListener) func() {
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

func AddEventListenerFunc(eventName string, listener EventListenerFunc) func() {
	return AddEventListener(eventName, listener)
}

func AddEventListenerChannel(eventName string, c chan Event) func() {
	return AddEventListenerFunc(eventName, func(event Event) {
		c <- event
	})
}

func URL() (*url.URL, error) {
	return url.Parse(Get("location").Get("href").String())
}

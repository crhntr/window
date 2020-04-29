// +build js,wasm

package dom

import (
	"syscall/js"
)

type document byte

const (
	Document document = 0
)

var (
	doc = js.Global().Get("document")
)

func (_ document) JSValue() js.Value                           { return doc }
func (_ document) Get(key string) js.Value                     { return doc.Get(key) }
func (_ document) Set(key string, value interface{})           { doc.Set(key, value) }
func (_ document) Call(m string, args ...interface{}) js.Value { return doc.Call(m, args...) }
func (_ document) Type() js.Type                               { return doc.Type() }

// func (_ document) Truthy() bool { return doc.Truthy() }
// func (_ document) IsNull() bool { return doc.IsNull() }
// func (_ document) IsUndefined() bool { return doc.IsUndefined() }
// func (_ document) InstanceOf(t js.Value) bool { return doc.InstanceOf(t) }

func (el document) AddEventListener(eventName string, listener EventListener) {
	el.Call("addEventListener", eventName, js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		listener.HandleEvent(Event(args[0]))
		return nil
	}))
}

func (el document) AddEventListenerFunc(eventName string, listener EventListenerFunc) {
	el.AddEventListener(eventName, listener)
}

// +build js,wasm

package window

import (
	"syscall/js"
)

type Event js.Value

func (el Event) Get(key string) js.Value { return js.Value(el).Get(key) }

func (el Event) Set(key string, value interface{}) { js.Value(el).Set(key, value) }

func (el Event) Call(m string, args ...interface{}) js.Value { return js.Value(el).Call(m, args...) }

func (el Event) JSValue() js.Value { return js.Value(el) }

func (el Event) Type() js.Type { return js.Value(el).Type() }

func (el Event) EventType() string { return js.Value(el).Get("type").String() }

func (el Event) Truthy() bool { return js.Value(el).Truthy() }

func (el Event) IsNull() bool { return js.Value(el).IsNull() }

func (el Event) IsUndefined() bool { return js.Value(el).IsUndefined() }

func (el Event) InstanceOf(t js.Value) bool { return js.Value(el).InstanceOf(t) }

func (el Event) Target() js.Value { return el.Get("target") }

func (el Event) PreventDefault() { el.Call("preventDefault") }

func (el Event) Log() { js.Global().Get("console").Call("log", el) }

// KeyCode can be used on keyup and keydown events
func (el Event) KeyCode() int { return el.Get("keyCode").Int() }

type EventListener interface {
	HandleEvent(event Event)
}

type EventListenerFunc func(Event)

//go:build js && wasm
// +build js,wasm

package window

import (
	"syscall/js"
)

type Event js.Value

func (ev Event) Get(key string) js.Value { return js.Value(ev).Get(key) }

func (ev Event) Set(key string, value interface{}) { js.Value(ev).Set(key, value) }

func (ev Event) Call(m string, args ...interface{}) js.Value { return js.Value(ev).Call(m, args...) }

func (ev Event) JSValue() js.Value { return js.Value(ev) }

func (ev Event) Type() js.Type { return js.Value(ev).Type() }

func (ev Event) EventType() string { return js.Value(ev).Get("type").String() }

func (ev Event) Truthy() bool { return js.Value(ev).Truthy() }

func (ev Event) IsNull() bool { return js.Value(ev).IsNull() }

func (ev Event) IsUndefined() bool { return js.Value(ev).IsUndefined() }

func (ev Event) InstanceOf(t js.Value) bool { return js.Value(ev).InstanceOf(t) }

func (ev Event) Target() js.Value { return ev.Get("target") }

func (ev Event) Equal(w js.Value) bool { return js.Value(ev).Equal(w) }

func (ev Event) TargetElement() Element { return Element(ev.Target()) }

func (ev Event) PreventDefault() { ev.Call("preventDefault") }

func (ev Event) StopPropagation() { ev.Call("stopPropagation") }

func (ev Event) Log() { js.Global().Get("console").Call("log", ev) }

// KeyCode can be used on keyup and keydown events
func (ev Event) KeyCode() int { return ev.Get("keyCode").Int() }

type EventListener interface {
	HandleEvent(event Event)
}

type EventListenerFunc func(Event)

func (fn EventListenerFunc) HandleEvent(event Event) { fn(event) }

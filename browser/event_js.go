package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type Event js.Value

func (e Event) Target() dom.Element {
	return valueToElement(js.Value(e).Get("target"))
}

type EventListenerFunc js.Func

func (e EventListenerFunc) Release() { js.Func(e).Release() }

func NewEventListenerFunc(fn func(Event)) EventListenerFunc {
	return EventListenerFunc(js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(Event(args[0]))
		return nil
	}))
}

func (e Element) AddEventListener(name string, listener EventListenerFunc) {
	v(e).addEventListener(name, listener)
}

func (in Input) AddEventListener(name string, listener EventListenerFunc) {
	v(in).addEventListener(name, listener)
}

func (d Document) AddEventListener(name string, listener EventListenerFunc) {
	v(d).addEventListener(name, listener)
}

func (val v) addEventListener(name string, listener EventListenerFunc) {
	js.Value(val).Call("addEventListener", name, js.Func(listener))
}

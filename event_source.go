//go:build js && wasm
// +build js,wasm

package window

import (
	"sync"
	"syscall/js"
)

type ServerEventHandlerFunc = func(id, data string)

type EventSource struct {
	js.Value
	handlers map[string][]js.Func
	sync.Mutex
	closed bool

	LogEvents bool
}

func NewEventSource(srcURL string) *EventSource {
	return &EventSource{
		Value:    js.Global().Get("EventSource").New(srcURL),
		handlers: make(map[string][]js.Func),
	}
}

func (es *EventSource) Handle(eventName string, handler ServerEventHandlerFunc) {
	es.Lock()
	es.Unlock()

	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		msg := args[0]

		data := msg.Get("data").String()
		id := msg.Get("lastEventId").String()

		if es.LogEvents {
			js.Global().Get("console").Call("log", "EventSource", eventName, id, js.Global().Get("JSON").Call("parse", data))
		}

		go handler(id, data)

		return nil
	})

	es.handlers[eventName] = append(es.handlers[eventName], fn)

	es.Call("addEventListener", eventName, fn)
}

func (es *EventSource) Close() {
	es.Lock()
	es.Unlock()

	if es == nil || es.closed {
		return
	}
	es.closed = true

	for eventName, fns := range es.handlers {
		for _, fn := range fns {
			es.Call("removeEventListener", eventName, fn)
			fn.Release()
		}
	}
	es.handlers = nil

	if es.Truthy() {
		es.Value.Call("close")
	}
}

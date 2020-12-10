// +build js,wasm

package window

import (
	"syscall/js"
)

type ServerEventHandlerFunc = func(id, data string)

type EventSource struct {
	js.Value
	handlers map[string][]js.Func

	LogEvents bool
}

func NewEventSource(srcURL string) EventSource {
	return EventSource{
		Value:    js.Global().Get("EventSource").New(srcURL),
		handlers: make(map[string][]js.Func),
	}
}

func (es *EventSource) Handle(eventName string, handler ServerEventHandlerFunc) {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		msg := args[0]

		data := msg.Get("data").String()
		id := msg.Get("lastEventId").String()

		go handler(id, data)

		if es.LogEvents {
			js.Global().Get("console").Call("log", "EventSource", id, js.Global().Get("JSON").Call("parse", data))
		}

		return nil
	})

	es.handlers[eventName] = append(es.handlers[eventName], fn)

	es.Call("addEventListener", eventName, fn)
}

func (es *EventSource) Close() {
	for eventName, fns := range es.handlers {
		for _, fn := range fns {
			es.Call("removeEventListener", eventName, fn)
			fn.Release()
		}
	}
}

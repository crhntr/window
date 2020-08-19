// +build js,wasm

package window

import (
	"errors"
	"syscall/js"
)

type ServerEventHandlerFunc = func(id, data string)

type EventSource struct {
	js.Value
	Handlers map[string]ServerEventHandlerFunc

	LogEvents bool
}

func NewEventSource(srcURL string) (*EventSource, error) {
	esv := js.Global().Get("EventSource").New(srcURL)

	if !esv.Truthy() {
		return nil, errors.New("created object is falsy")
	}

	return &EventSource{
		LogEvents: true,
		Value:     esv,
		Handlers:  make(map[string]ServerEventHandlerFunc),
	}, nil
}

func (es *EventSource) Handle(eventName string, handler ServerEventHandlerFunc) func() {
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

	es.Call("addEventListener", eventName, fn)

	return func() {
		defer fn.Release()

		es.Call("removeEventListener", eventName, fn)
	}
}

package window

import (
	"reflect"
	"syscall/js"

	"github.com/crhntr/window/dom"
)

func Window() dom.Window { return dom.Window(js.Global()) }

func Document() dom.Document {
	return dom.Document(js.Global().Get("document"))
}

func NewObject() js.Value { return js.Global().Get("Object").New() }

func NewUint8ClampedArray(length int) js.Value {
	return js.Global().Get("Uint8ClampedArray").New(length)
}

func WrapEventListenerFunc[E dom.EventValue](fn func(event E)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(E(args[0]))
		return nil
	})
}

func NewEventSource(url string, withCredentials bool) dom.EventSource {
	return dom.EventSource(js.Global().Get("EventSource").New(url, map[string]interface{}{"withCredentials": withCredentials}))
}

func ConsoleLog(a ...any) {
	jsV := reflect.TypeOf(js.Value{})
	for i := range a {
		if _, ok := a[i].(js.Value); ok {
			continue
		}
		va := reflect.ValueOf(a[i])
		if va.CanConvert(jsV) {
			a[i] = va.Convert(jsV).Interface().(js.Value)
		}
	}
	js.Global().Get("console").Call("log", a...)
}

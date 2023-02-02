//go:build js && wasm

package window

import (
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

func ConsoleLog(a ...any) {
	for i := range a {
		switch x := a[i].(type) {
		case dom.Text:
			a[i] = js.Value(x)
		case dom.HTMLElement:
			a[i] = js.Value(x)
		case dom.SVGElement:
			a[i] = js.Value(x)
		case dom.GenericEvent:
			a[i] = js.Value(x)
		case dom.UIEvent:
			a[i] = js.Value(x)
			// TODO: generate a function for this switch
		}
	}
	js.Global().Get("console").Call("log", a...)
}

//go:build js && wasm
// +build js,wasm

package window

import (
	"syscall/js"

	"github.com/crhntr/window/browser"
)

var Document = browser.Document(js.Global().Get("document"))

func NewObject() js.Value { return js.Global().Get("Object").New() }

func NewUint8ClampedArray(length int) js.Value {
	return js.Global().Get("Uint8ClampedArray").New(length)
}

func AddEventListener(name string, listener browser.EventListenerFunc) {
	js.Global().Call("addEventListener", name, js.Func(listener))
}

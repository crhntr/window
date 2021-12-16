//go:build js && wasm

package window

import (
	"encoding/base64"
	"syscall/js"
	"testing"

	"github.com/crhntr/please"
)

func TestWindow_Get(t *testing.T) {
	result := Get("self")
	please.ExpectTrue(t, result.Truthy())
	please.ExpectTrue(t, result.InstanceOf(js.Global().Get("Window")))
}

func TestWindow_SetDeleteGet(t *testing.T) {
	Set("exampleField", js.ValueOf(9000))
	please.ExpectTrue(t, Get("exampleField").Truthy())
	Delete("exampleField")
	please.ExpectFalse(t, Get("exampleField").Truthy())
}

func TestWindow_Call(t *testing.T) {
	decoded := Call("atob", base64.StdEncoding.EncodeToString([]byte("Hello")))
	please.ExpectEqual(t, decoded.String(), "Hello")
}

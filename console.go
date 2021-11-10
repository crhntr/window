//go:build js && wasm
// +build js,wasm

package window

import (
	"syscall/js"
)

type console byte

const (
	Console console = 0
)

var con = win.Get("console")

func (console console) JSValue() js.Value                           { return con }
func (console console) Get(key string) js.Value                     { return con.Get(key) }
func (console console) Set(key string, value interface{})           { con.Set(key, value) }
func (console console) Call(m string, args ...interface{}) js.Value { return con.Call(m, args...) }
func (console console) Type() js.Type                               { return con.Type() }

func (console console) Log(objects ...interface{}) {
	con.Call("log", objects...)
}

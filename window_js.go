//go:build js

package window

import (
	"syscall/js"
)

var window = js.Global()
var Document = document(Get("document"))

func Call(m string, args ...interface{}) js.Value { return window.Call(m, args...) }
func Delete(key string)                           { window.Delete(key) }
func Get(key string) js.Value                     { return window.Get(key) }
func Set(key string, value js.Value)              { window.Set(key, value) }

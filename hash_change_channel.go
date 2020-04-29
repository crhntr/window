// +build js,wasm

package dom

import (
	"fmt"
	"net/url"
	"syscall/js"
)

func HashChanges() <-chan string {
	changeFile := make(chan string)
	js.Global().Get("window").Call("addEventListener", "hashchange",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			go func() {
				u, err := url.Parse(js.Global().Get("window").Get("location").Get("hash").String())
				if err != nil {
					fmt.Println("could not parse URL hash on change", err)
					return
				}
				changeFile <- u.Fragment
			}()
			return nil
		}), false)
	return changeFile
}

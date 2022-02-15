//go:build js && wasm
// +build js,wasm

package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type ElementCollection js.Value

func (e ElementCollection) Length() int                       { return v(e).length() }
func (e ElementCollection) Item(index int) dom.Element        { return v(e).item(index) }
func (e ElementCollection) NamedItem(name string) dom.Element { return v(e).namedItem(name) }

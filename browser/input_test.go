//go:build js && wasm
// +build js,wasm

package browser

import "github.com/crhntr/window/dom"

var _ dom.InputElement = Input{}

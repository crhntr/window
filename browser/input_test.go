//go:build js && wasm

package browser

import "github.com/crhntr/window/dom"

var _ dom.InputElement = Input{}
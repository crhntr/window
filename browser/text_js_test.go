//go:build js

package browser

import (
	"syscall/js"
	"testing"

	"github.com/crhntr/window/dom"
)

var _ dom.Text = Text{}

func createTextNode(t *testing.T) dom.ChildNode {
	return Text(js.Global().Get("document").Call("createTextNode", "node"))
}

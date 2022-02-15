//go:build js && wasm
// +build js,wasm

package browser

import (
	"syscall/js"
	"testing"

	"github.com/crhntr/window/dom"
	"github.com/crhntr/window/dom/domtest"
)

var _ dom.DocumentFragment = DocumentFragment{}

func createDocumentFragmentParentNode(t *testing.T) dom.ParentNode {
	return DocumentFragment(js.Global().Get("document").Call("createDocumentFragment"))
}

func TestDocumentFragment(t *testing.T) {
	domtest.ParentNode(t, createDocumentFragmentParentNode, createElementNode)
	domtest.ParentNode(t, createDocumentFragmentParentNode, createTextNode)
}

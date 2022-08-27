//go:build js && wasm

package browser

import (
	"syscall/js"
	"testing"

	"github.com/crhntr/please"

	"github.com/crhntr/window/dom/domtest"

	"github.com/crhntr/window/dom"
)

var _ dom.Document = Document{}

func TestDocument(t *testing.T) {
	doc := js.Global().Get("document")
	please.ExpectTrue(t, doc.InstanceOf(js.Global().Get("Document")))

	domtest.Document(t, func(t *testing.T) dom.Node {
		return Document(js.Global().Get("document"))
	})
}

func TestDocument_ChildElementCount(t *testing.T) {
	//doc := Document(js.Global().Get("document"))
	//please.ExpectEqual(t, doc.(), 1)
}

func TestDocument_QuerySelector(t *testing.T) {
	doc := Document(js.Global().Get("document"))
	button := doc.QuerySelector("button")
	please.ExpectTrue(t, button != nil)
}

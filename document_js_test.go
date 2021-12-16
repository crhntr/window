//go:build js

package window

import (
	"syscall/js"
	"testing"

	"github.com/crhntr/please"
	"github.com/crhntr/window/dom"
)

var _ dom.Document = document{}

func TestDocument(t *testing.T) {
	please.ExpectTrue(t, js.Value(Document).InstanceOf(Get("Document")))
}

func TestDocument_ChildElementCount(t *testing.T) {
	please.ExpectEqual(t, Document.ChildElementCount(), 1)
}

func TestDocument_SetGetDelete(t *testing.T) {
	js.Value(Document).Set("exampleField", js.ValueOf(9000))
	please.ExpectTrue(t, js.Value(Document).Get("exampleField").Truthy())
	js.Value(Document).Delete("exampleField")
	please.ExpectFalse(t, js.Value(Document).Get("exampleField").Truthy())
}

func TestDocument_QuerySelector(t *testing.T) {
	button := Document.QuerySelector("button")
	please.ExpectTrue(t, button != nil)
}

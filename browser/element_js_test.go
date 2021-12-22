package browser

import (
	"syscall/js"
	"testing"

	"github.com/crhntr/window/dom"
	"github.com/crhntr/window/dom/domtest"
)

var _ dom.Element = Element{}
var _ dom.ElementCollection = ElementCollection{}

func createElement(t *testing.T, s string) dom.Element {
	div := Document(js.Global().Get("document")).CreateElement("div")
	div.SetInnerHTML(s)
	return div.FirstElementChild()
}

func createElementNode(t *testing.T) dom.ChildNode {
	div := Document(js.Global().Get("document")).CreateElement("div")
	if js.Value(div.(Element)).IsUndefined() {
		t.Errorf("div is undefined")
	}
	if js.Value(div.(Element)).IsNull() {
		t.Errorf("div is null")
	}
	return div
}

func createElementParentNode(t *testing.T) dom.ParentNode {
	return createElementNode(t).(dom.ParentNode)
}

func TestElement(t *testing.T) {
	domtest.ChildNode(t, createElementNode)
	domtest.ParentNode(t, createElementParentNode, createElementNode)
	domtest.ParentNode(t, createElementParentNode, createElementNode)

	domtest.ElementTextContent(t, createElement)
	domtest.ElementTagName(t, createElement)
	domtest.ElementAttribute(t, createElement)
	domtest.ElementInnerHTML(t, createElement)
	domtest.ElementOuterHTML(t, createElement)
}

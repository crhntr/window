//go:build js && wasm

package window_test

import (
	"testing"

	"github.com/crhntr/window"
)

func TestWindow(t *testing.T) {
	w := window.Window()
	w.Name()
	w.Document()
	w.InnerHeight()
	w.InnerWidth()
	w.IsSecureContext()
}

func TestDocument(t *testing.T) {
	d := window.Document()
	if d.Body() == nil {
		t.Fail()
	}
	if d.Body() == nil {
		t.Fail()
	}
}

func TestConsoleLog(t *testing.T) {
	window.ConsoleLog("Hello, string!")

	text := window.Document().CreateTextNode("Hello, text node!")
	window.ConsoleLog(text)

	h1 := window.Document().CreateElement("h1")
	h1.SetAttribute("data-message", "some-element")
	window.Document().Body().Append(h1)
	window.ConsoleLog(h1)
}

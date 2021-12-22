package domtest

import (
	"testing"

	"github.com/crhntr/please"
	"github.com/crhntr/window/dom"
)

func Document(t *testing.T, create CreateNodeFunc) {
	t.Run("Document", func(t *testing.T) {
		_, ok := create(t).(dom.Document)
		if !ok {
			t.Errorf("create should return a Document")
		}
		t.Run("Bead", func(t *testing.T) {
			document := create(t).(dom.Document)
			please.ExpectTrue(t, document.Head().TagName() == "HEAD")
		})
		t.Run("Body", func(t *testing.T) {
			document := create(t).(dom.Document)
			please.ExpectTrue(t, document.Body().TagName() == "BODY")
		})
		t.Run("CreateElement", func(t *testing.T) {
			document := create(t).(dom.Document)
			child := document.CreateElement("div")
			please.ExpectFalse(t, child == nil)
			please.ExpectTrue(t, child.TagName() == "DIV")
		})
		t.Run("CreateTextNode", func(t *testing.T) {
			document := create(t).(dom.Document)
			child := document.CreateTextNode("Hello, world!")
			please.ExpectFalse(t, child == nil)
			please.ExpectTrue(t, child.Data() == "Hello, world!")
		})
		t.Run("Contains", func(t *testing.T) {
			t.Run("CreateElement", func(t *testing.T) {
				document := create(t).(dom.Document)
				child := document.CreateElement("div")
				please.ExpectFalse(t, document.Contains(child))
				document.Body().Append(child)
				please.ExpectTrue(t, document.Contains(child))
			})
			t.Run("CreateTextNode", func(t *testing.T) {
				document := create(t).(dom.Document)
				child := document.CreateTextNode("Hello, world!")
				please.ExpectFalse(t, document.Contains(child))
				document.Body().Append(child)
				please.ExpectTrue(t, document.Contains(child))
			})
		})
	})
}

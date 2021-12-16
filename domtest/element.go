package domtest

import (
	"testing"

	"github.com/crhntr/please"

	"github.com/crhntr/window/dom"
)

type CreateNodeFunc func(t *testing.T, nodeType dom.NodeType, input string) dom.Node

func ElementTagName(t *testing.T, create CreateNodeFunc) {
	t.Run("tagName", func(t *testing.T) {
		el, ok := create(t, dom.NodeTypeElement, `<div></div>`).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}
		result := el.TagName()
		please.ExpectEqual(t, result, "div")
	})
}

func ElementNodeType(t *testing.T, create CreateNodeFunc) {
	t.Run("nodeType", func(t *testing.T) {
		el, ok := create(t, dom.NodeTypeElement, `<div></div>`).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}
		result := el.NodeType()
		please.ExpectEqual(t, result, dom.NodeTypeElement)
	})
}

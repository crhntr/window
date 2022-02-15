//go:build js && wasm
// +build js,wasm

package browser

import (
	"strings"
	"syscall/js"

	"github.com/crhntr/window/dom"
)

func nodeType(node js.Value) dom.NodeType { return dom.NodeType(node.Get("nodeType").Int()) }

func nodeToValue(n dom.Node) js.Value {
	if n == nil {
		// not sure if this should be js.Undefined()?
		return js.Null()
	}
	switch val := n.(type) {
	case Input:
		return js.Value(val)
	case Element:
		return js.Value(val)
	case DocumentFragment:
		return js.Value(val)
	case Document:
		return js.ValueOf(val)
	case Text:
		return js.Value(val)
	default:
		panic("unknown node type")
	}
}

func valueToNode(val js.Value) dom.Node {
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	switch nodeType(val) {
	case dom.NodeTypeElement:
		return valueToElement(val)
	case dom.NodeTypeDocumentFragment:
		return DocumentFragment(val)
	case dom.NodeTypeDocument:
		return Document(val)
	case dom.NodeTypeText:
		return Text(val)
	default:
		panic("unknown node type")
	}
}

func valueToChildNode(val js.Value) dom.ChildNode {
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	switch nodeType(val) {
	case dom.NodeTypeElement:
		return valueToElement(val)
	case dom.NodeTypeDocumentFragment:
		return DocumentFragment(val)
	case dom.NodeTypeText:
		return Text(val)
	default:
		panic("unknown node type")
	}
}

func valueToElement(val js.Value) dom.Element {
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	switch strings.ToLower(v(val).tagName()) {
	case "input", "select":
		return Input(val)
	default:
		return Element(val)
	}
}

func nodesToValuesAsEmptyInterfaceSlice(nodes []dom.ChildNode) []interface{} {
	list := make([]interface{}, len(nodes))
	for i := range nodes {
		list[i] = nodeToValue(nodes[i])
	}
	return list
}

func getRootNodeOptions(composed bool) js.Value {
	return js.ValueOf(map[string]interface{}{"composed": composed})
}

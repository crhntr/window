//go:build js

package window

import (
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
	case Element:
		return js.Value(val)
	case DocumentFragment:
		return js.Value(val)
	case document:
		return js.ValueOf(val)
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
		return Element(val)
	case dom.NodeTypeDocumentFragment:
		return DocumentFragment(val)
	case dom.NodeTypeDocument:
		return document(val)
	default:
		panic("unknown node type")
	}
}

func nodesToEmptyInterfaceSlice(nodes []dom.Node) []interface{} {
	list := make([]interface{}, len(nodes))
	for i := range nodes {
		list[i] = nodes[i]
	}
	return list
}

func getRootNodeOptions(composed bool) js.Value {
	return js.ValueOf(map[string]interface{}{"composed": composed})
}

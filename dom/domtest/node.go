package domtest

import (
	"testing"

	"github.com/crhntr/please"
	"github.com/crhntr/window/dom"
)

type CreateNodeFunc func(t *testing.T) dom.Node

func Node(t *testing.T, create CreateNodeFunc) {
	t.Run("NodeType", func(t *testing.T) {
		node := create(t)

		result := node.NodeType()

		switch node.(type) {
		case dom.Text:
			please.ExpectEqual(t, result, dom.NodeTypeText)
		case dom.Element:
			please.ExpectEqual(t, result, dom.NodeTypeElement)
		case dom.Document:
			please.ExpectEqual(t, result, dom.NodeTypeDocument)
		case dom.DocumentFragment:
			please.ExpectEqual(t, result, dom.NodeTypeDocumentFragment)
		case dom.Comment:
			please.ExpectEqual(t, result, dom.NodeTypeComment)
		default:
			t.Errorf("unknown node type %[1]T: %[1]v", node)
		}
	})

	shouldNotPanic := func(t *testing.T) {
		t.Helper()
		r := recover()
		if r != nil {
			t.Errorf("it should not panic, but got: %v", r)
		}
	}

	t.Run("CloneNode", func(t *testing.T) {
		node := create(t)

		t.Run("deep", func(t *testing.T) {
			defer shouldNotPanic(t)
			node.CloneNode(true)
		})

		t.Run("shallow", func(t *testing.T) {
			defer shouldNotPanic(t)
			node.CloneNode(false)
		})
	})

	t.Run("IsSameNode", func(t *testing.T) {
		defer shouldNotPanic(t)
		node := create(t)
		node.IsSameNode(node)
	})
}

type CreateChildNodeFunc func(t *testing.T) dom.ChildNode

func ChildNode(t *testing.T, create CreateChildNodeFunc) {
	Node(t, func(t *testing.T) dom.Node {
		return create(t)
	})
	shouldNotPanic := func(t *testing.T) {
		t.Helper()
		r := recover()
		if r != nil {
			t.Errorf("it should not panic, but got: %v", r)
		}
	}

	t.Run("IsConnected", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.IsConnected()
	})

	t.Run("OwnerDocument", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.OwnerDocument()
	})

	t.Run("ParentNode", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.ParentNode()
	})

	t.Run("ParentElement", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.ParentElement()
	})

	t.Run("PreviousSibling", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.PreviousSibling()
	})

	t.Run("NextSibling", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.NextSibling()
	})

	t.Run("TextContent", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.TextContent()
	})

	t.Run("Length", func(t *testing.T) {
		defer shouldNotPanic(t)
		node, ok := create(t).(dom.ChildNode)
		if !ok {
			t.SkipNow()
		}
		node.Length()
	})
}

type CreateParentNodeFunc func(t *testing.T) dom.ParentNode

// ParentNode tests parent child relationships. It should be passed a parent with no
// children. Most of the getter type methods are not tested independently but are tested
// in conjunction with mutating methods.
func ParentNode(t *testing.T, createParent CreateParentNodeFunc, createNode CreateChildNodeFunc) {
	t.Run("created parent node is empty", func(t *testing.T) {
		parent := createParent(t)
		please.ExpectEqual(t, parent.ChildNodes().Length(), 0)
		please.ExpectEqual(t, parent.FirstChild(), nil)
		please.ExpectEqual(t, parent.LastChild(), nil)
		please.ExpectFalse(t, parent.HasChildNodes())
	})

	t.Run("AppendChild", func(t *testing.T) {
		t.Run("append one node", func(t *testing.T) {
			parent := createParent(t)

			child := createNode(t)
			please.ExpectEqual(t, child.PreviousSibling(), nil)
			please.ExpectEqual(t, child.NextSibling(), nil)

			result := parent.AppendChild(child)

			please.ExpectTrue(t, result.IsSameNode(child))
			please.ExpectTrue(t, child.ParentNode().IsSameNode(parent))
			please.ExpectEqual(t, child.PreviousSibling(), nil)
			please.ExpectEqual(t, child.NextSibling(), nil)

			please.ExpectTrue(t, parent.Contains(child))

			stranger := createNode(t)
			please.ExpectFalse(t, parent.Contains(stranger))

			t.Run("ChildNodes", func(t *testing.T) {
				nodes := parent.ChildNodes()
				please.ExpectEqual(t, nodes.Length(), 1)
				please.ExpectTrue(t, nodes.Item(0).IsSameNode(child))
			})
		})

		t.Run("append two element nodes", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)

			_ = parent.AppendChild(child1)
			result := parent.AppendChild(child2)
			please.ExpectTrue(t, result.IsSameNode(child2))
			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)

			please.ExpectTrue(t, child1.ParentNode().IsSameNode(parent))
			please.ExpectTrue(t, child2.ParentNode().IsSameNode(parent))

			please.ExpectEqual(t, child1.PreviousSibling(), nil)
			please.ExpectTrue(t, child1.NextSibling().IsSameNode(child2))
			please.ExpectEqual(t, child2.NextSibling(), nil)
			please.ExpectTrue(t, child2.PreviousSibling().IsSameNode(child1))

			please.ExpectTrue(t, parent.HasChildNodes())

			please.ExpectTrue(t, parent.Contains(child1))
			please.ExpectTrue(t, parent.Contains(child2))

			t.Run("ChildNodes", func(t *testing.T) {
				nodes := parent.ChildNodes()
				please.ExpectEqual(t, nodes.Length(), 2)
				please.ExpectTrue(t, nodes.Item(0).IsSameNode(child1))
				please.ExpectTrue(t, nodes.Item(1).IsSameNode(child2))
			})
		})

		t.Run("append existing node", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)

			_ = parent.AppendChild(child2)
			_ = parent.AppendChild(child1)
			_ = parent.AppendChild(child2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)

			please.ExpectEqual(t, child1.PreviousSibling(), nil)
			please.ExpectTrue(t, child1.NextSibling().IsSameNode(child2))
			please.ExpectEqual(t, child2.NextSibling(), nil)
			please.ExpectTrue(t, child2.PreviousSibling().IsSameNode(child1))

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)
			please.ExpectTrue(t, parent.FirstChild().IsSameNode(child1))
			please.ExpectTrue(t, parent.LastChild().IsSameNode(child2))
			please.ExpectTrue(t, parent.HasChildNodes())

			please.ExpectTrue(t, parent.Contains(child1))
			please.ExpectTrue(t, parent.Contains(child2))

			t.Run("ChildNodes", func(t *testing.T) {
				nodes := parent.ChildNodes()
				please.ExpectEqual(t, nodes.Length(), 2)
				please.ExpectTrue(t, nodes.Item(0).IsSameNode(child1))
				please.ExpectTrue(t, nodes.Item(1).IsSameNode(child2))
			})
		})
	})

	t.Run("InsertBefore", func(t *testing.T) {
		parent := createParent(t)

		child1 := createNode(t)
		child2 := createNode(t)

		_ = parent.AppendChild(child2)
		result := parent.InsertBefore(child1, child2)
		please.ExpectTrue(t, result.IsSameNode(child1))

		please.ExpectTrue(t, child1.ParentNode().IsSameNode(parent))
		please.ExpectTrue(t, child2.ParentNode().IsSameNode(parent))

		please.ExpectEqual(t, child1.PreviousSibling(), nil)
		please.ExpectTrue(t, child1.NextSibling().IsSameNode(child2))
		please.ExpectEqual(t, child2.NextSibling(), nil)
		please.ExpectTrue(t, child2.PreviousSibling().IsSameNode(child1))

		please.ExpectEqual(t, parent.ChildNodes().Length(), 2)
		please.ExpectTrue(t, parent.FirstChild().IsSameNode(child1))
		please.ExpectTrue(t, parent.LastChild().IsSameNode(child2))
		please.ExpectTrue(t, parent.HasChildNodes())

		please.ExpectTrue(t, parent.Contains(child1))
		please.ExpectTrue(t, parent.Contains(child2))
	})

	t.Run("RemoveChild", func(t *testing.T) {
		t.Run("with an empty parent", func(t *testing.T) {
			parent := createParent(t)
			child := createNode(t)
			parent.AppendChild(child)

			result := parent.RemoveChild(child)
			please.ExpectTrue(t, result.IsSameNode(child))
			please.ExpectEqual(t, parent.ChildNodes().Length(), 0)
			please.ExpectEqual(t, parent.FirstChild(), nil)
			please.ExpectEqual(t, parent.LastChild(), nil)
		})

		t.Run("with one sibling after", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)

			result := parent.RemoveChild(child1)
			please.ExpectTrue(t, result.IsSameNode(child1))
			please.ExpectEqual(t, parent.ChildNodes().Length(), 1)
			please.ExpectTrue(t, parent.FirstChild().IsSameNode(child2))
			please.ExpectTrue(t, parent.LastChild().IsSameNode(child2))
		})

		t.Run("with one sibling before", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)

			result := parent.RemoveChild(child2)
			please.ExpectTrue(t, result.IsSameNode(child2))
			please.ExpectEqual(t, parent.ChildNodes().Length(), 1)
			please.ExpectTrue(t, parent.FirstChild().IsSameNode(child1))
			please.ExpectTrue(t, parent.LastChild().IsSameNode(child1))
		})

		t.Run("with siblings before and after", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			child3 := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			parent.AppendChild(child3)

			result := parent.RemoveChild(child2)
			please.ExpectTrue(t, result.IsSameNode(child2))
			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)
			please.ExpectTrue(t, parent.FirstChild().IsSameNode(child1))
			please.ExpectTrue(t, parent.LastChild().IsSameNode(child3))
			please.ExpectTrue(t, child1.NextSibling().IsSameNode(child3))
			please.ExpectTrue(t, child3.PreviousSibling().IsSameNode(child1))
		})
	})

	t.Run("ReplaceChild", func(t *testing.T) {
		t.Run("first child", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			replacement := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			result := parent.ReplaceChild(replacement, child1)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)
			please.ExpectTrue(t, result.IsSameNode(child1))
			please.ExpectTrue(t, parent.FirstChild().IsSameNode(replacement))
		})

		t.Run("last child", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			replacement := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			result := parent.ReplaceChild(replacement, child2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)
			please.ExpectTrue(t, result.IsSameNode(child2))
			please.ExpectTrue(t, parent.LastChild().IsSameNode(replacement))
		})

		t.Run("middle child", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			child3 := createNode(t)
			replacement := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			parent.AppendChild(child3)
			result := parent.ReplaceChild(replacement, child2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 3)
			please.ExpectTrue(t, result.IsSameNode(child2))
			please.ExpectTrue(t, child1.NextSibling().IsSameNode(replacement))
		})
	})

	t.Run("ReplaceChildren", func(t *testing.T) {
		t.Run("children exist but no new nodes", func(t *testing.T) {
			parent := createParent(t)
			parent.ReplaceChildren()
			please.ExpectEqual(t, parent.ChildNodes().Length(), 0)
		})

		t.Run("when replacing child nodes with multiple replacement nodes", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			child3 := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			parent.AppendChild(child3)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 3)

			replacement1 := createNode(t)
			replacement2 := createNode(t)

			parent.ReplaceChildren(replacement1, replacement2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)

			nodes := parent.ChildNodes()
			please.ExpectTrue(t, nodes.Item(0).IsSameNode(replacement1))
			please.ExpectTrue(t, nodes.Item(1).IsSameNode(replacement2))
		})

		t.Run("when replacing child nodes with multiple replacement nodes", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			child3 := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			parent.AppendChild(child3)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 3)

			replacement1 := createNode(t)
			replacement2 := createNode(t)

			parent.ReplaceChildren(replacement1, replacement2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)

			nodes := parent.ChildNodes()
			please.ExpectTrue(t, nodes.Item(0).IsSameNode(replacement1))
			please.ExpectTrue(t, nodes.Item(1).IsSameNode(replacement2))
		})
	})

	t.Run("Prepend", func(t *testing.T) {
		t.Run("with existing children", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			child3 := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			parent.AppendChild(child3)

			additional1 := createNode(t)
			additional2 := createNode(t)

			parent.Prepend(additional1, additional2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 5)

			please.ExpectTrue(t, parent.FirstChild().IsSameNode(additional1))
			please.ExpectTrue(t, additional1.NextSibling().IsSameNode(additional2))
			please.ExpectTrue(t, additional2.NextSibling().IsSameNode(child1))
		})

		t.Run("without existing children", func(t *testing.T) {
			parent := createParent(t)

			additional1 := createNode(t)
			additional2 := createNode(t)

			parent.Prepend(additional1, additional2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)
			please.ExpectTrue(t, parent.FirstChild().IsSameNode(additional1))
			please.ExpectTrue(t, parent.FirstChild().(dom.ChildNode).NextSibling().IsSameNode(additional2))
			please.ExpectEqual(t, parent.FirstChild().(dom.ChildNode).NextSibling().NextSibling(), nil)
		})
	})

	t.Run("Append", func(t *testing.T) {
		t.Run("with existing children", func(t *testing.T) {
			parent := createParent(t)

			child1 := createNode(t)
			child2 := createNode(t)
			child3 := createNode(t)

			parent.AppendChild(child1)
			parent.AppendChild(child2)
			parent.AppendChild(child3)

			additional1 := createNode(t)
			additional2 := createNode(t)

			parent.Append(additional1, additional2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 5)
			please.ExpectTrue(t, child3.NextSibling().IsSameNode(additional1))
			please.ExpectTrue(t, additional1.PreviousSibling().IsSameNode(child3))
			please.ExpectTrue(t, additional1.NextSibling().IsSameNode(additional2))
			please.ExpectEqual(t, additional2.NextSibling(), nil)
			please.ExpectTrue(t, parent.LastChild().IsSameNode(additional2))
		})

		t.Run("without existing children", func(t *testing.T) {
			parent := createParent(t)

			additional1 := createNode(t)
			additional2 := createNode(t)

			parent.Append(additional1, additional2)

			please.ExpectEqual(t, parent.ChildNodes().Length(), 2)
			please.ExpectTrue(t, parent.FirstChild().IsSameNode(additional1))
			please.ExpectTrue(t, parent.FirstChild().NextSibling().IsSameNode(additional2))
			please.ExpectEqual(t, parent.FirstChild().NextSibling().NextSibling(), nil)
		})
	})
}

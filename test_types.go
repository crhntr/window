package window

import (
	"syscall/js"
)

var (
	_ Node = Document
	_ Node = Element(js.Null())
	_ Node = DocumentFragment(js.Null())
	_ Node = Text(js.Null())

	_ ParentNode = document(0)
	_ ParentNode = Element(js.Null())
	_ ParentNode = DocumentFragment(js.Null())

	_ ChildNode = Element(js.Null())
	_ ChildNode = Text(js.Null())
)

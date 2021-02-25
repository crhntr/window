package window

import (
	"syscall/js"
)

// Element types checker
var (
	_ js.Wrapper  = Element{}
	_ Node        = Element{}
	_ NodeWrapper = Element{}
	_ ChildNode   = Element{}
	_ ParentNode  = Element{}
)

// Text types checker
var (
	_ js.Wrapper  = Text{}
	_ Node        = Text{}
	_ NodeWrapper = Text{}
	_ ChildNode   = Text{}
)

// document types checker
var (
	_ js.Wrapper  = document(0)
	_ Node        = document(0)
	_ NodeWrapper = document(0)
	_ ParentNode  = document(0)
)

// DocumentFragment types checker
var (
	_ js.Wrapper  = DocumentFragment{}
	_ Node        = DocumentFragment{}
	_ NodeWrapper = DocumentFragment{}
	_ ParentNode  = DocumentFragment{}
)

package window

import (
	"syscall/js"
)

func (document document) Selection() Selection {
	return Selection(document.Call("getSelection"))
}

type (
	Selection js.Value
	Range     js.Value
)

func (selection Selection) Range(index int) Range {
	return Range(js.Value(selection).Call("getRangeAt", index))
}

func (selection Selection) RangeCount() int {
	return js.Value(selection).Get("rangeCount").Int()
}

func (selection Selection) AddRange(ran Range) {
	js.Value(selection).Call("addRange", ran)
}

func (selection Selection) RemoveAllRanges() {
	js.Value(selection).Call("removeAllRanges")
}

func (selection Selection) JSValue() js.Value { return js.Value(selection) }

func (selection Selection) Get(key string) js.Value           { return selection.JSValue().Get(key) }
func (selection Selection) Equal(w js.Value) bool             { return selection.JSValue().Equal(w) }
func (selection Selection) Set(key string, value interface{}) { selection.JSValue().Set(key, value) }
func (selection Selection) Call(m string, args ...interface{}) js.Value {
	return selection.JSValue().Call(m, args...)
}
func (selection Selection) Type() js.Type              { return selection.JSValue().Type() }
func (selection Selection) Truthy() bool               { return selection.JSValue().Truthy() }
func (selection Selection) IsNull() bool               { return selection.JSValue().IsNull() }
func (selection Selection) IsUndefined() bool          { return selection.JSValue().IsUndefined() }
func (selection Selection) InstanceOf(t js.Value) bool { return selection.JSValue().InstanceOf(t) }

func (document document) NewRange() Range {
	return Range(win.New("Range"))
}

func (ran Range) CommonAncestor(node js.Value) Node {
	return nodeFactory(ran.JSValue().Call("commonAncestorContainer", node))
}

func (ran Range) StartContainer() Node {
	return nodeFactory(ran.JSValue().Get("startContainer"))
}

func (ran Range) EndContainer() Node {
	return nodeFactory(ran.JSValue().Get("endContainer"))
}

func (ran Range) StartOffset() int {
	return ran.JSValue().Get("startOffset").Int()
}

func (ran Range) EndOffset() int {
	return ran.JSValue().Get("endOffset").Int()
}

func (ran Range) SetStart(node Node, offset int) {
	ran.JSValue().Call("setEnd", node, offset)
}

func (ran Range) SetEnd(node Node, offset int) {
	ran.JSValue().Call("setEnd", node, offset)
}

func (ran Range) SetStartBefore(node js.Value) {
	ran.JSValue().Call("setStartBefore", node)
}

func (ran Range) SetEndBefore(node js.Value) {
	ran.JSValue().Call("setEndBefore", node)
}

func (ran Range) SetStartAfter(node js.Value) {
	ran.JSValue().Call("setStartAfter", node)
}

func (ran Range) SetEndAfter(node js.Value) {
	ran.JSValue().Call("setEndAfter", node)
}

func (ran Range) Select(node js.Value) {
	ran.JSValue().Call("selectNode", node)
}

func (ran Range) CreateContextualFragment(content string) DocumentFragment {
	return DocumentFragment(ran.JSValue().Call("createContextualFragment", content))
}

func (ran Range) JSValue() js.Value { return js.Value(ran) }

func (ran Range) Get(key string) js.Value           { return ran.Get(key) }
func (ran Range) Equal(w js.Value) bool             { return ran.Equal(w) }
func (ran Range) Set(key string, value interface{}) { ran.Set(key, value) }
func (ran Range) Call(m string, args ...interface{}) js.Value {
	return ran.Call(m, args...)
}
func (ran Range) Type() js.Type              { return ran.Type() }
func (ran Range) Truthy() bool               { return ran.Truthy() }
func (ran Range) IsNull() bool               { return ran.IsNull() }
func (ran Range) IsUndefined() bool          { return ran.IsUndefined() }
func (ran Range) InstanceOf(t js.Value) bool { return ran.InstanceOf(t) }

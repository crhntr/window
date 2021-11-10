//go:build js && wasm
// +build js,wasm

package window

import (
	"syscall/js"
)

type Text js.Value

func (t Text) Equal(value js.Value) bool { return js.Value(t).Equal(value) }
func (t Text) IsNull() bool              { return js.Value(t).IsNull() }

func (t Text) Node() Node                              { return t }
func (t Text) JSValue() js.Value                       { return js.Value(t) }
func (t Text) NodeType() NodeType                      { return nodeType(t) }
func (t Text) FirstChild() ChildNode                   { return firstChild(t) }
func (t Text) IsConnected() bool                       { return isConnected(t) }
func (t Text) LastChild() ChildNode                    { return lastChild(t) }
func (t Text) ChildNodes() NodeList                    { return childNodes(t) }
func (t Text) AppendChild(child NodeWrapper) ChildNode { return appendChild(t, child) }
func (t Text) RemoveChild(child NodeWrapper) Node      { return removeChild(t, child) }
func (t Text) ReplaceChild(newChild, oldChild NodeWrapper) Node {
	return replaceChild(t, newChild, oldChild)
}
func (t Text) Contains(child NodeWrapper) bool { return contains(t, child) }
func (t Text) NextSibling() ChildNode          { return nextSibling(t) }
func (t Text) PreviousSibling() ChildNode      { return previousSibling(t) }
func (t Text) ParentNode() Node                { return parentNode(t) }
func (t Text) ParentElement() Element          { return parentElement(t) }
func (t Text) CloneNode(isDeep bool) Node      { return cloneNode(t, isDeep) }
func (t Text) HasChildNodes() bool             { return hasChildNodes(t) }
func (t Text) Normalize()                      { normalize(t) }

func (t Text) IsSameNode(node Node) bool  { return isSameNode(t, node) }
func (t Text) IsEqualNode(node Node) bool { return isEqualNode(t, node) }

func (t Text) Remove()                  { childNodeRemove(t) }
func (t Text) Before(node ...Node)      { childNodeBefore(t, node) }
func (t Text) After(node ...Node)       { childNodeAfter(t, node) }
func (t Text) ReplaceWith(node ...Node) { childNodeReplaceWith(t, node) }

func (t Text) WholeText() string { return js.Value(t).Call("wholeTExt").String() }

func (t Text) Length() int { return js.Value(t).Length() }

func (t Text) AsElement() Element { return nodeAsElement(t) }
func (t Text) AsText() Text       { return nodeAsText(t) }

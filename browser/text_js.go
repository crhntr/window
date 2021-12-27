//go:build js && wasm

package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type Text js.Value

func (t Text) NodeType() dom.NodeType                        { return v(t).nodeType() }
func (t Text) IsConnected() bool                             { return v(t).isConnected() }
func (t Text) OwnerDocument() dom.Document                   { return v(t).ownerDocument() }
func (t Text) ParentNode() dom.Node                          { return v(t).parentNode() }
func (t Text) ParentElement() dom.Element                    { return v(t).parentElement() }
func (t Text) PreviousSibling() dom.ChildNode                { return v(t).previousSibling() }
func (t Text) NextSibling() dom.ChildNode                    { return v(t).nextSibling() }
func (t Text) TextContent() string                           { return v(t).textContent() }
func (t Text) Normalize()                                    { v(t).normalize() }
func (t Text) CloneNode(deep bool) dom.Node                  { return v(t).cloneNode(deep) }
func (t Text) IsEqualNode(other dom.Node) bool               { return v(t).isEqualNode(other) }
func (t Text) IsSameNode(other dom.Node) bool                { return v(t).isSameNode(other) }
func (t Text) CompareDocumentPosition() dom.DocumentPosition { return v(t).compareDocumentPosition() }

func (t Text) Data() string         { return v(t).data() }
func (t Text) SetData(s string)     { v(t).setData(s) }
func (t Text) Split(n int) dom.Text { return v(t).split(n) }
func (t Text) WholeText() string    { return v(t).wholeText() }

func (t Text) Length() int { return v(t).length() }

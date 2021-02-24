package window

import (
	"fmt"
	"syscall/js"
)

type AsElementer interface {
	AsElement() Element
}

type AsTexter interface {
	AsText() Text
}

func nodeAsElement(node js.Wrapper) Element {
	return Element(node.JSValue())
}

func nodeAsText(node js.Wrapper) Text {
	return Text(node.JSValue())
}

type NodeWrapper interface {
	js.Wrapper

	Node() Node
}

type Node interface {
	AsElementer
	AsTexter

	NodeWrapper

	IsNull() bool

	NodeType() NodeType
	FirstChild() ChildNode
	IsConnected() bool
	LastChild() ChildNode
	ChildNodes() NodeList
	AppendChild(child NodeWrapper) ChildNode
	RemoveChild(child NodeWrapper) Node
	ReplaceChild(newChild, oldChild NodeWrapper) Node
	Contains(child NodeWrapper) bool
	NextSibling() ChildNode
	PreviousSibling() ChildNode
	ParentNode() Node
	ParentElement() Element
	CloneNode(isDeep bool) Node
	HasChildNodes() bool
	Normalize()
	IsSameNode(Node) bool
	IsEqualNode(Node) bool
}

func nodeFactory(value js.Wrapper) Node {
	if wrapper, ok := value.(NodeWrapper); ok {
		return wrapper.Node()
	}

	if v := value.JSValue(); v.IsNull() {
		return Element(v)
	}

	switch nodeType(value) {
	case NodeTypeElement:
		return Element(value.JSValue())
	case NodeTypeDocument:
		return Document
	case NodeTypeText:
		return Text(value.JSValue())
	case NodeTypeDocumentFragment:
		return DocumentFragment(value.JSValue())
	default:
		return nil
	}
}

func nodeType(node js.Wrapper) NodeType {
	return NodeType(node.JSValue().Get("nodeType").Int())
}

type NodeType int

const (
	NodeTypeUnknown NodeType = iota
	NodeTypeElement
	NodeTypeAttribute
	NodeTypeText
	NodeTypeCdataSection
	NodeTypeEntityReference
	NodeTypeEntity
	NodeTypeProcessingInstruction
	NodeTypeComment
	NodeTypeDocument
	NodeTypeDocumentType
	NodeTypeDocumentFragment
	NodeTypeNotation
)

func (nt NodeType) IsElement() bool {
	return nt == NodeTypeElement
}

func (nt NodeType) IsText() bool {
	return nt == NodeTypeText
}

func (nt NodeType) IsDocumentFragment() bool {
	return nt == NodeTypeDocumentFragment
}

func firstChild(node js.Wrapper) ChildNode {
	return nodeFactory(node.JSValue().Get("firstChild")).(ChildNode)
}

func isConnected(node js.Wrapper) bool {
	return node.JSValue().Get("isConnected").Bool()
}

func lastChild(node js.Wrapper) ChildNode {
	return nodeFactory(node.JSValue().Get("lastChild")).(ChildNode)
}

func childNodes(node js.Wrapper) NodeList {
	return NodeList(node.JSValue().Get("childNodes"))
}

func appendChild(node, child js.Wrapper) ChildNode {
	return nodeFactory(node.JSValue().Call("appendChild", child)).(ChildNode)
}

func removeChild(node, child js.Wrapper) Node {
	return nodeFactory(node.JSValue().Call("removeChild", child))
}

func replaceChild(node, newChild, oldChild js.Wrapper) Node {
	return nodeFactory(node.JSValue().Call("replaceChild", newChild, oldChild))
}

func contains(node, child js.Wrapper) bool {
	return node.JSValue().Call("contains", child).Bool()
}

func nextSibling(node js.Wrapper) ChildNode {
	return nodeFactory(node.JSValue().Get("nextSibling")).(ChildNode)
}

func previousSibling(node js.Wrapper) ChildNode {
	return nodeFactory(node.JSValue().Get("previousSibling")).(ChildNode)
}

func parentNode(node js.Wrapper) Node {
	return nodeFactory(node.JSValue().Get("parentNode"))
}

func parentElement(node js.Wrapper) Element {
	return Element(node.JSValue().Get("parentElement"))
}

func cloneNode(node js.Wrapper, isDeep bool) Node {
	return nodeFactory(node.JSValue().Call("cloneNode", isDeep))
}

type DocumentPosition int

const (
	DocumentPositionDisconnected DocumentPosition = 1 << iota
	DocumentPositionPreceding
	DocumentPositionFollowing
	DocumentPositionContains
	DocumentPositionContainedBy
	DocumentPositionImplementationSpecific
)

func compareDocumentPosition(node, other Node) DocumentPosition {
	return DocumentPosition(node.JSValue().Call("compareDocumentPosition", other).Int())
}

func (pos DocumentPosition) IsDisconnected() bool {
	return pos&DocumentPositionDisconnected != 0
}
func (pos DocumentPosition) IsPreceding() bool {
	return pos&DocumentPositionPreceding != 0
}
func (pos DocumentPosition) IsFollowing() bool {
	return pos&DocumentPositionFollowing != 0
}
func (pos DocumentPosition) IsContains() bool {
	return pos&DocumentPositionContains != 0
}
func (pos DocumentPosition) IsContainedBy() bool {
	return pos&DocumentPositionContainedBy != 0
}
func (pos DocumentPosition) IsImplementationSpecific() bool {
	return pos&DocumentPositionImplementationSpecific != 0
}

func nodeContains(node, other js.Wrapper) bool {
	return node.JSValue().Call("contains", node).Bool()
}

func hasChildNodes(node js.Wrapper) bool {
	return node.JSValue().Call("hasChildNodes", node).Bool()
}

func insertBefore(node, child js.Wrapper) bool {
	return node.JSValue().Call("insertBefore", child).Bool()
}

func isEqualNode(node, other js.Wrapper) bool {
	value := node.JSValue()
	return other.JSValue().IsNull() == value.IsNull() && node.JSValue().Call("isEqualNode", other).Bool()
}

func isSameNode(node, other js.Wrapper) bool {
	value := node.JSValue()
	return !other.JSValue().IsNull() && !value.IsNull() && value.Call("isSameNode", other).Bool()
}

func normalize(node js.Wrapper) { node.JSValue().Call("normalize") }

type ParentNode interface {
	Node

	// properties
	ChildElementCount() int
	Children() HTMLCollection
	FirstElementChild() Element
	LastElementChild() Element

	// methods
	Append(...Node)
	Prepend(...Node)
	QuerySelector(query string, a ...interface{}) Element
	QuerySelectorAll(query string, a ...interface{}) NodeList
	ReplaceChildren(...Node)
}

func childElementCount(parent js.Wrapper) int { return parent.JSValue().Get("childElementCount").Int() }
func children(parent js.Wrapper) HTMLCollection {
	return HTMLCollection(parent.JSValue().Get("children"))
}
func firstElementChild(parent js.Wrapper) Element {
	return Element(parent.JSValue().Get("firstElementChild"))
}
func lastElementChild(parent js.Wrapper) Element {
	return Element(parent.JSValue().Get("lastElementChild"))
}
func appendNodes(parent js.Wrapper, nodes []Node) {
	parent.JSValue().Call("append", nodesToInterfaceSlice(nodes)...)
}
func prepend(parent js.Wrapper, nodes []Node) {
	parent.JSValue().Call("prepend", nodesToInterfaceSlice(nodes)...)
}
func querySelector(parent js.Wrapper, query string, a []interface{}) Element {
	query = fmt.Sprintf(query, a...)
	defer recoverSelector(query)
	return Element(parent.JSValue().Call("querySelector", query))
}
func querySelectorAll(parent js.Wrapper, query string, a []interface{}) NodeList {
	query = fmt.Sprintf(query, a...)
	defer recoverSelector(query)
	return NodeList(parent.JSValue().Call("querySelectorAll", query))
}
func replaceChildren(parent js.Wrapper, nodes []Node) {
	parent.JSValue().Call("replaceChildren", nodesToInterfaceSlice(nodes)...)
}

func recoverSelector(query string) {
	if r := recover(); r != nil {
		if err, ok := r.(js.Error); ok {
			panic(fmt.Errorf("query selector failed for %q: %w", query, err))
		}
	}
}

func nodesToInterfaceSlice(slice []Node) []interface{} {
	list := make([]interface{}, len(slice))

	for i, el := range slice {
		list[i] = el
	}

	return list
}

type ChildNode interface {
	Node

	Remove()
	Before(...Node)
	After(...Node)
	ReplaceWith(...Node)
}

func childNodeRemove(node js.Wrapper) {
	v := node.JSValue()
	if v.IsNull() {
		return
	}
	v.Call("remove")
}
func childNodeBefore(node js.Wrapper, list []Node) {
	node.JSValue().Call("before", nodesToInterfaceSlice(list)...)
}
func childNodeAfter(node js.Wrapper, list []Node) {
	node.JSValue().Call("after", nodesToInterfaceSlice(list)...)
}
func childNodeReplaceWith(node js.Wrapper, list []Node) {
	node.JSValue().Call("replaceWith", nodesToInterfaceSlice(list)...)
}

type NodeList js.Value

func (list NodeList) Length() int {
	return js.Value(list).Length()
}

func (list NodeList) Item(i int) ChildNode {
	return nodeFactory(js.Value(list).Call("item", i)).(ChildNode)
}

func (list NodeList) Slice() []ChildNode {
	nodes := make([]ChildNode, 0, list.Length())
	for i := 0; i < list.Length(); i++ {
		nodes = append(nodes, list.Item(i))
	}
	return nodes
}

func (list NodeList) ElementSlice() []Element {
	nodes := make([]Element, 0, list.Length())
	for i := 0; i < list.Length(); i++ {
		node := list.Item(i)
		if !node.NodeType().IsElement() {
			continue
		}
		nodes = append(nodes, Element(node.JSValue()))
	}
	return nodes
}

func (list NodeList) ForEach(fn func(node ChildNode, index int, list NodeList)) {
	wrappedFn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(nodeFactory(args[0]).(ChildNode), args[1].Int(), NodeList(args[2]))
		return nil
	})
	defer wrappedFn.Release()

	js.Value(list).Call("forEach", wrappedFn)
}

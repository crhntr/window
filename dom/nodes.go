package dom

// Node is based on a subset of the methods and types in
// https://dom.spec.whatwg.org/#interface-node as of 2021-12-10
//
// ParentNode and ChildNode contain more of the methods from the spec.
type Node interface {
	NodeType() NodeType
	CloneNode(deep bool) Node
	IsSameNode(other Node) bool
	TextContent() string
}

type ChildNode interface {
	Node

	IsConnected() bool
	OwnerDocument() Document
	ParentNode() Node
	ParentElement() Element
	PreviousSibling() ChildNode
	NextSibling() ChildNode

	// CompareDocumentPosition(other Node) DocumentPosition

	// Length should be based on https://dom.spec.whatwg.org/#concept-node-length
	Length() int

	// LookupPrefix(namespace string)
	// LookupNamespaceURI(prefix string)
	// IsDefaultNamespace(namespace string) bool
}

// ParentNode is based on https://dom.spec.whatwg.org/#interface-parentnode. It also includes some fields and
// methods from Node that only make sense for non-leaf nodes such as Element, DocumentFragment, and Document.
type ParentNode interface {
	Node
	ElementParent

	Prepend(nodes ...ChildNode)
	Append(nodes ...ChildNode)
	ReplaceChildren(nodes ...ChildNode)

	// the following methods are from node; however, they only make sense for parent nodes

	HasChildNodes() bool
	ChildNodes() NodeList
	FirstChild() ChildNode
	LastChild() ChildNode
	InsertBefore(node, child ChildNode) ChildNode
	AppendChild(node ChildNode) ChildNode
	ReplaceChild(node, child ChildNode) ChildNode
	RemoveChild(node ChildNode) ChildNode

	NodeContainer
}

type NodeContainer interface {
	Contains(other Node) bool
}

type ElementParent interface {
	Children() ElementCollection
	FirstElementChild() Element
	LastElementChild() Element
	ChildElementCount() int
	GetElementByID(id string) Element
	GetElementsByTagName(name string) ElementCollection
	GetElementsByClassName(name string) ElementCollection
}

type SelectorMethods interface {
	Closest(selector string) Element
	Matches(selector string) bool

	QuerySelector(query string) Element
	QuerySelectorAll(query string) NodeList
}

// Normalizer may be implemented by a Node and should follow https://dom.spec.whatwg.org/#dom-node-normalize
type Normalizer interface {
	Normalize()
}

type DocumentPosition int

// DocumentPosition is based on const values in
// https://dom.spec.whatwg.org/#interface-node (reviewed on 2021-12-10)
const (
	DocumentPositionDisconnected DocumentPosition = 1 << iota
	DocumentPositionPreceding
	DocumentPositionFollowing
	DocumentPositionContains
	DocumentPositionContainedBy
	DocumentPositionImplementationSpecific
)

// NodeType is based on const values in
// https://dom.spec.whatwg.org/#interface-node (reviewed on 2021-12-10)
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

func (nt NodeType) String() string {
	switch nt {
	case NodeTypeElement:
		return "Element"
	case NodeTypeAttribute:
		return "Attribute"
	case NodeTypeText:
		return "Text"
	case NodeTypeCdataSection:
		return "CdataSection"
	case NodeTypeEntityReference:
		return "EntityReference"
	case NodeTypeEntity:
		return "Entity"
	case NodeTypeProcessingInstruction:
		return "ProcessingInstruction"
	case NodeTypeComment:
		return "Comment"
	case NodeTypeDocument:
		return "Document"
	case NodeTypeDocumentType:
		return "DocumentType"
	case NodeTypeDocumentFragment:
		return "DocumentFragment"
	case NodeTypeNotation:
		return "Notation"

	default:
		fallthrough
	case NodeTypeUnknown:
		return "Unknown"
	}
}

type NodeList interface {
	Length() int
	Item(int) Node
}

type Text interface {
	ChildNode

	Data() string
	SetData(string)

	// Split(n int) Text // CONSIDER: maybe implement this
	// WholeText() string // CONSIDER: maybe implement this
}

type Document interface {
	Node

	CreateElement(localName string) Element
	CreateTextNode(text string) Text

	// CreateDocumentFragment() node

	NodeContainer

	Body() Element
	Head() Element
}

// Element is based on
//
// InnerText methods are ignored due to rendering complexity; however, implementations may add them
// based on InnerTextSetter.
type Element interface {
	Node
	ChildNode
	ParentNode

	TagName() string
	ID() string
	ClassName() string

	Attribute(name string) string
	SetAttribute(name, value string)
	RemoveAttribute(name string)
	ToggleAttribute(name string) bool
	HasAttribute(name string) bool

	SetInnerHTML(s string)
	InnerHTML() string
	SetOuterHTML(s string)
	OuterHTML() string

	InsertAdjacentHTML(InsertAdjacentHTMLPosition, string)
}

// InputElement may be implemented by input Elements
type InputElement interface {
	Element

	Value() string
	SetValue(value string)

	Disabled() bool
	SetDisabled(disabled bool)
}

type InsertAdjacentHTMLPosition string

const (
	PositionBeforeBegin = InsertAdjacentHTMLPosition("beforebegin")
	PositionAfterBegin  = InsertAdjacentHTMLPosition("afterbegin")
	PositionBeforeEnd   = InsertAdjacentHTMLPosition("beforeend")
	PositionAfterEnd    = InsertAdjacentHTMLPosition("afterend")
)

func (pos InsertAdjacentHTMLPosition) String() string {
	return string(pos)
}

type InnerTextSetter interface {
	SetInnerText(s string)
	InnerText() string
}

type ElementCollection interface {
	// Length returns the number of elements in the collection.
	Length() int

	// Item returns the element with index from the collection. The elements are sorted in tree order.
	Item(index int) Element

	// NamedItem returns the first element with ID or name from the collection.
	NamedItem(name string) Element
}

type DocumentFragment interface {
	ParentNode
}

type Comment interface {
	Node

	Data() string
	SetData() string
}

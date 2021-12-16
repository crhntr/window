package dom

// Node is based on a subset of the methods and types in
// https://dom.spec.whatwg.org/#interface-node as of 2021-12-10
type Node interface {
	NodeType() NodeType

	IsConnected() bool
	OwnerDocument() Node
	GetRootNode(composed bool) Node
	ParentNode() Node
	ParentElement() Element
	HasChildNodes() bool
	ChildNodes() NodeList
	FirstChild() Node
	LastChild() Node
	PreviousSibling() Node
	NextSibling() Node

	NodeValue() string
	TextContent() string
	Normalize()

	CloneNode(deep bool) Node
	IsEqualNode(other Node) bool
	IsSameNode(other Node) bool

	CompareDocumentPosition() DocumentPosition
	Contains(other Node) bool

	// LookupPrefix(namespace string)
	// LookupNamespaceURI(prefix string)
	// IsDefaultNamespace(namespace string) bool

	InsertBefore(node, child Node) Node
	AppendChild(node Node) Node
	ReplaceChild(node, child Node) Node
	RemoveChild(node Node) Node
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

type NodeList interface {
	Length() int
	Item(int) Node
}

type Text interface {
	Data() string
	Split(n int) Text
	WholeText() string
}

type Document interface {
	Node
	ParentNode

	CreateElement(localName string) Element
	CreateElementIs(localName, is string) Element
	CreateElementNS(namespace, localName string) Element
	CreateElementNSIS(namespace, localName, is string) Element

	// CreateDocumentFragment() Node

	CreateTextNode(text string) Text
}

type ParentNode interface {
	Children() NodeList
	FirstElementChild() Element
	LastElementChild() Element
	ChildElementCount() int

	Prepend(nodes ...Node) Node
	Append(nodes ...Node) Node
	ReplaceChildren(nodes ...Node) Node

	GetElementsByTagName(name string) ElementCollection
	GetElementsByTagNameNS(namespace, name string) ElementCollection
	GetElementsByClassName(name string) ElementCollection

	QuerySelector(query string) Element
	QuerySelectorAll(query string) NodeList
}

type Element interface {
	Node
	ParentNode

	TagName() string
	ID() string
	ClassName() string

	GetAttribute(name string) string
	GetAttributeNS(namespace, name string) string
	SetAttribute(name, value string)
	SetAttributeNS(namespace, name, value string)
	RemoveAttribute(name string)
	RemoveAttributeNS(namespace, name string)
	ToggleAttribute(name string) bool
	HasAttribute(name string) bool
	HasAttributeNS(namespace, name string) bool

	Closest(selector string) Element
	Matches(selector string) bool

	SetInnerHTML(s string)
	InnerHTML() string
	SetOuterHTML(s string)
	OuterHTML() string
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
	Node
}

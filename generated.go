package window

import (
	"syscall/js"
	"time"
)

type Event interface {
	Bubbles() bool
	Cancelable() bool
	Composed() bool
	CurrentTarget() js.Value
	DefaultPrevented() bool
	EventPhase() EventPhase
	IsTrusted() bool
	Target() js.Value
	TimeStamp() time.Time
	Type() string
	PreventDefault()
	StopImmediatePropagation()
	StopPropagation()
}
type Node interface {
	BaseURI() string
	ChildNodes() NodeList
	FirstChild() Node
	IsConnected() bool
	LastChild() Node
	NextSibling() Node
	NodeName() string
	NodeType() NodeType
	NodeValue() Node
	ParentNode() Node
	ParentElement() Element
	PreviousSibling() Node
	TextContent() string
	AppendChild(aChild Node) Node
	CloneNode(deep bool) Node
	CompareDocumentPosition(otherNode Node) DocumentPosition
	Contains(node Node) bool
	GetRootNode(options GetRootNodeOptions) Node
	HasChildNodes() bool
	InsertBefore(newNode Node, referenceNode Node) bool
	IsDefaultNamespace(namespaceURL string) bool
	IsEqualNode(otherNode Node) bool
	IsOtherNode(otherNode Node) bool
	LookupPrefix(prefix string) string
	LookupNamespaceURI(prefix string) string
	Normalize()
	RemoveChild(child Node)
	ReplaceChild(newChild Node, oldChild Node)
}
type Element interface {
	BaseURI() string
	ChildNodes() NodeList
	FirstChild() Node
	IsConnected() bool
	LastChild() Node
	NextSibling() Node
	NodeName() string
	NodeType() NodeType
	NodeValue() Node
	ParentNode() Node
	ParentElement() Element
	PreviousSibling() Node
	TextContent() string
	ChildElementCount() int
	Children() HTMLCollection
	ClassList() DOMTokenList
	ClassName() string
	ClientHeight() int
	ClientLeft() int
	ClientTop() int
	ClientWidth() int
	FirstElementChild() Element
	ID() string
	InnerHTML() string
	LocalName() string
	NamespaceURI() string
	NextElementSibling() Element
	OuterHTML() string
	Part() DOMTokenList
	Prefix() string
	PreviousElementSibling() Element
	ScrollHeight() int
	ScrollLeft() int
	ScrollTop() int
	ScrollWidth() int
	Slot() string
	TagName() string
	AppendChild(aChild Node) Node
	CloneNode(deep bool) Node
	CompareDocumentPosition(otherNode Node) DocumentPosition
	Contains(node Node) bool
	GetRootNode(options GetRootNodeOptions) Node
	HasChildNodes() bool
	InsertBefore(newNode Node, referenceNode Node) bool
	IsDefaultNamespace(namespaceURL string) bool
	IsEqualNode(otherNode Node) bool
	IsOtherNode(otherNode Node) bool
	LookupPrefix(prefix string) string
	LookupNamespaceURI(prefix string) string
	Normalize()
	RemoveChild(child Node)
	ReplaceChild(newChild Node, oldChild Node)
	AddEventListener(eventType string, listener func(event Event), options AddEventListenerOptions, useCapture bool)
	After(node Node)
	Append(params Node)
	Before(params Node)
	Closest(selector string) Element
	DispatchEvent(event Event) bool
	GetAttribute(name string) string
	GetAttributeNames() string
	GetBoundingClientRect() DOMRect
	GetElementsByClassName(names string) HTMLCollection
	GetElementsByTagName(tagName string) HTMLCollection
	GetElementsByTagNameNS(namespaceURI string, localName string) HTMLCollection
	HasAttribute(name string) bool
	HasAttributeNS(namespaceURI string, localName string) bool
	HasAttributes() bool
	HasPointerCapture(pointerId int) bool
	InsertAdjacentElement(position AdjacentPosition, element Element) Element
	InsertAdjacentHTML(position AdjacentPosition, text string)
	InsertAdjacentText(position AdjacentPosition, text string)
	Matches(selectors string) HTMLCollection
	Prepend(params Node)
	QuerySelector(selectors string) Element
	QuerySelectorAll(selectors string) NodeList
	ReleasePointerCapture(pointerId int)
	Remove()
	RemoveAttribute(attrName string)
	RemoveAttributeNS(namespaceURI string, localName string)
	RemoveEventListener(eventType string, listener func(event Event), options AddEventListenerOptions, useCapture bool)
	ReplaceChildren(node Node)
	ReplaceWith(node Node)
	Scroll(options ScrollOptions)
	ScrollBy(options ScrollOptions)
	ScrollIntoView(options ScrollIntoViewOptions)
	ScrollTo(options ScrollOptions)
	SetAttribute(name string, value string)
	SetAttributeNS(namespace string, name string, value string)
	SetPointerCapture(pointerId int)
	ToggleAttribute(name string)
}
type DOMRect js.Value

func wrapDOMRect(value js.Value) DOMRect {
	return DOMRect(value)
}
func (val DOMRect) X() int {
	return js.Value(val).Call("x").Int()
}
func (val DOMRect) Y() int {
	return js.Value(val).Call("y").Int()
}
func (val DOMRect) Width() int {
	return js.Value(val).Call("width").Int()
}
func (val DOMRect) Height() int {
	return js.Value(val).Call("height").Int()
}
func (val DOMRect) Top() int {
	return js.Value(val).Call("top").Int()
}
func (val DOMRect) Right() int {
	return js.Value(val).Call("right").Int()
}
func (val DOMRect) Bottom() int {
	return js.Value(val).Call("bottom").Int()
}
func (val DOMRect) Left() int {
	return js.Value(val).Call("left").Int()
}

type NodeList js.Value

func wrapNodeList(value js.Value) NodeList {
	return NodeList(value)
}
func (val NodeList) Length() int {
	return js.Value(val).Length()
}
func (val NodeList) Item(index int) Node {
	return wrapNode(js.Value(val).Call("item", index))
}

type HTMLCollection js.Value

func wrapHTMLCollection(value js.Value) HTMLCollection {
	return HTMLCollection(value)
}
func (val HTMLCollection) Length() int {
	return js.Value(val).Length()
}
func (val HTMLCollection) Item(index int) Node {
	return wrapNode(js.Value(val).Call("item", index))
}
func (val HTMLCollection) NamedItem(name string) Node {
	return wrapNode(js.Value(val).Call("namedItem", name))
}

type DOMTokenList js.Value

func wrapDOMTokenList(value js.Value) DOMTokenList {
	return DOMTokenList(value)
}
func (val DOMTokenList) Length() int {
	return js.Value(val).Length()
}
func (val DOMTokenList) Value() string {
	return js.Value(val).Call("value").String()
}
func (val DOMTokenList) Item(index int) string {
	return js.Value(val).Call("item", index).String()
}
func (val DOMTokenList) Contains(token string) bool {
	return js.Value(val).Call("contains", token).Bool()
}
func (val DOMTokenList) Add(token string)  {
	js.Value(val).Call("add", token)
}
func (val DOMTokenList) Remove(token string)  {
	js.Value(val).Call("remove", token)
}
func (val DOMTokenList) Replace(oldToken string, newToken string)  {
	js.Value(val).Call("replace", oldToken, newToken)
}
func (val DOMTokenList) Supports(token string) bool {
	return js.Value(val).Call("supports", token).Bool()
}
func (val DOMTokenList) Toggle(token string) bool {
	return js.Value(val).Call("toggle", token).Bool()
}
func (val DOMTokenList) ToggleForce(token string, force bool) bool {
	return js.Value(val).Call("toggle", token, force).Bool()
}

type HTMLElement js.Value

func wrapHTMLElement(value js.Value) HTMLElement {
	return HTMLElement(value)
}
func (val HTMLElement) BaseURI() string {
	return js.Value(val).Call("baseURI").String()
}
func (val HTMLElement) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Call("childNodes"))
}
func (val HTMLElement) FirstChild() Node {
	return wrapNode(js.Value(val).Call("firstChild"))
}
func (val HTMLElement) IsConnected() bool {
	return js.Value(val).Call("isConnected").Bool()
}
func (val HTMLElement) LastChild() Node {
	return wrapNode(js.Value(val).Call("lastChild"))
}
func (val HTMLElement) NextSibling() Node {
	return wrapNode(js.Value(val).Call("nextSibling"))
}
func (val HTMLElement) NodeName() string {
	return js.Value(val).Call("nodeName").String()
}
func (val HTMLElement) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Call("nodeType"))
}
func (val HTMLElement) NodeValue() Node {
	return wrapNode(js.Value(val).Call("nodeValue"))
}
func (val HTMLElement) ParentNode() Node {
	return wrapNode(js.Value(val).Call("parentNode"))
}
func (val HTMLElement) ParentElement() Element {
	return wrapElement(js.Value(val).Call("parentElement"))
}
func (val HTMLElement) PreviousSibling() Node {
	return wrapNode(js.Value(val).Call("previousSibling"))
}
func (val HTMLElement) TextContent() string {
	return js.Value(val).Call("textContent").String()
}
func (val HTMLElement) ChildElementCount() int {
	return js.Value(val).Call("childElementCount").Int()
}
func (val HTMLElement) Children() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("children"))
}
func (val HTMLElement) ClassList() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Call("classList"))
}
func (val HTMLElement) ClassName() string {
	return js.Value(val).Call("className").String()
}
func (val HTMLElement) ClientHeight() int {
	return js.Value(val).Call("clientHeight").Int()
}
func (val HTMLElement) ClientLeft() int {
	return js.Value(val).Call("clientLeft").Int()
}
func (val HTMLElement) ClientTop() int {
	return js.Value(val).Call("clientTop").Int()
}
func (val HTMLElement) ClientWidth() int {
	return js.Value(val).Call("clientWidth").Int()
}
func (val HTMLElement) FirstElementChild() Element {
	return wrapElement(js.Value(val).Call("firstElementChild"))
}
func (val HTMLElement) ID() string {
	return js.Value(val).Call("id").String()
}
func (val HTMLElement) InnerHTML() string {
	return js.Value(val).Call("innerHTML").String()
}
func (val HTMLElement) LocalName() string {
	return js.Value(val).Call("localName").String()
}
func (val HTMLElement) NamespaceURI() string {
	return js.Value(val).Call("namespaceURI").String()
}
func (val HTMLElement) NextElementSibling() Element {
	return wrapElement(js.Value(val).Call("nextElementSibling"))
}
func (val HTMLElement) OuterHTML() string {
	return js.Value(val).Call("outerHTML").String()
}
func (val HTMLElement) Part() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Call("part"))
}
func (val HTMLElement) Prefix() string {
	return js.Value(val).Call("prefix").String()
}
func (val HTMLElement) PreviousElementSibling() Element {
	return wrapElement(js.Value(val).Call("previousElementSibling"))
}
func (val HTMLElement) ScrollHeight() int {
	return js.Value(val).Call("scrollHeight").Int()
}
func (val HTMLElement) ScrollLeft() int {
	return js.Value(val).Call("scrollLeft").Int()
}
func (val HTMLElement) ScrollTop() int {
	return js.Value(val).Call("scrollTop").Int()
}
func (val HTMLElement) ScrollWidth() int {
	return js.Value(val).Call("scrollWidth").Int()
}
func (val HTMLElement) Slot() string {
	return js.Value(val).Call("slot").String()
}
func (val HTMLElement) TagName() string {
	return js.Value(val).Call("tagName").String()
}
func (val HTMLElement) AccessKey() string {
	return js.Value(val).Call("accessKey").String()
}
func (val HTMLElement) AccessKeyLabel() string {
	return js.Value(val).Call("accessKeyLabel").String()
}
func (val HTMLElement) IsContentEditable() bool {
	return js.Value(val).Call("isContentEditable").Bool()
}
func (val HTMLElement) Dataset() DOMStringMap {
	return wrapDOMStringMap(js.Value(val).Call("dataset"))
}
func (val HTMLElement) Dir() string {
	return js.Value(val).Call("dir").String()
}
func (val HTMLElement) Draggable() bool {
	return js.Value(val).Call("draggable").Bool()
}
func (val HTMLElement) EnterKeyHint() string {
	return js.Value(val).Call("enterKeyHint").String()
}
func (val HTMLElement) Hidden() bool {
	return js.Value(val).Call("hidden").Bool()
}
func (val HTMLElement) InnerText() string {
	return js.Value(val).Call("innerText").String()
}
func (val HTMLElement) Lang() string {
	return js.Value(val).Call("lang").String()
}
func (val HTMLElement) OffsetHeight() float64 {
	return js.Value(val).Call("offsetHeight").Float()
}
func (val HTMLElement) OffsetLeft() float64 {
	return js.Value(val).Call("offsetLeft").Float()
}
func (val HTMLElement) OffsetParent() float64 {
	return js.Value(val).Call("offsetParent").Float()
}
func (val HTMLElement) OffsetTop() float64 {
	return js.Value(val).Call("offsetTop").Float()
}
func (val HTMLElement) OffsetWidth() float64 {
	return js.Value(val).Call("offsetWidth").Float()
}
func (val HTMLElement) OuterText() string {
	return js.Value(val).Call("outerText").String()
}
func (val HTMLElement) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(js.Value(val).Call("style"))
}
func (val HTMLElement) TabIndex() int {
	return js.Value(val).Call("tabIndex").Int()
}
func (val HTMLElement) Title() string {
	return js.Value(val).Call("title").String()
}
func (val HTMLElement) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", aChild))
}
func (val HTMLElement) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val HTMLElement) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", otherNode))
}
func (val HTMLElement) Contains(node Node) bool {
	return js.Value(val).Call("contains", node).Bool()
}
func (val HTMLElement) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", options))
}
func (val HTMLElement) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val HTMLElement) InsertBefore(newNode Node, referenceNode Node) bool {
	return js.Value(val).Call("insertBefore", newNode, referenceNode).Bool()
}
func (val HTMLElement) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val HTMLElement) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", otherNode).Bool()
}
func (val HTMLElement) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", otherNode).Bool()
}
func (val HTMLElement) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val HTMLElement) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val HTMLElement) Normalize()  {
	js.Value(val).Call("normalize")
}
func (val HTMLElement) RemoveChild(child Node)  {
	js.Value(val).Call("removeChild", child)
}
func (val HTMLElement) ReplaceChild(newChild Node, oldChild Node)  {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}
func (val HTMLElement) AddEventListener(eventType string, listener func(event Event), options AddEventListenerOptions, useCapture bool)  {
	js.Value(val).Call("addEventListener", eventType, listener, options, useCapture)
}
func (val HTMLElement) After(node Node)  {
	js.Value(val).Call("after", node)
}
func (val HTMLElement) Append(params Node)  {
	js.Value(val).Call("append", params)
}
func (val HTMLElement) Before(params Node)  {
	js.Value(val).Call("before", params)
}
func (val HTMLElement) Closest(selector string) Element {
	return wrapElement(js.Value(val).Call("closest", selector))
}
func (val HTMLElement) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val HTMLElement) GetAttribute(name string) string {
	return js.Value(val).Call("getAttribute", name).String()
}
func (val HTMLElement) GetAttributeNames() string {
	return js.Value(val).Call("getAttributeNames").String()
}
func (val HTMLElement) GetBoundingClientRect() DOMRect {
	return wrapDOMRect(js.Value(val).Call("getBoundingClientRect"))
}
func (val HTMLElement) GetElementsByClassName(names string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByClassName", names))
}
func (val HTMLElement) GetElementsByTagName(tagName string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByTagName", tagName))
}
func (val HTMLElement) GetElementsByTagNameNS(namespaceURI string, localName string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByTagNameNS", namespaceURI, localName))
}
func (val HTMLElement) HasAttribute(name string) bool {
	return js.Value(val).Call("hasAttribute", name).Bool()
}
func (val HTMLElement) HasAttributeNS(namespaceURI string, localName string) bool {
	return js.Value(val).Call("hasAttributeNS", namespaceURI, localName).Bool()
}
func (val HTMLElement) HasAttributes() bool {
	return js.Value(val).Call("hasAttributes").Bool()
}
func (val HTMLElement) HasPointerCapture(pointerId int) bool {
	return js.Value(val).Call("hasPointerCapture", pointerId).Bool()
}
func (val HTMLElement) InsertAdjacentElement(position AdjacentPosition, element Element) Element {
	return wrapElement(js.Value(val).Call("insertAdjacentElement", position, element))
}
func (val HTMLElement) InsertAdjacentHTML(position AdjacentPosition, text string)  {
	js.Value(val).Call("insertAdjacentHTML", position, text)
}
func (val HTMLElement) InsertAdjacentText(position AdjacentPosition, text string)  {
	js.Value(val).Call("insertAdjacentText", position, text)
}
func (val HTMLElement) Matches(selectors string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("matches", selectors))
}
func (val HTMLElement) Prepend(params Node)  {
	js.Value(val).Call("prepend", params)
}
func (val HTMLElement) QuerySelector(selectors string) Element {
	return wrapElement(js.Value(val).Call("querySelector", selectors))
}
func (val HTMLElement) QuerySelectorAll(selectors string) NodeList {
	return wrapNodeList(js.Value(val).Call("querySelectorAll", selectors))
}
func (val HTMLElement) ReleasePointerCapture(pointerId int)  {
	js.Value(val).Call("releasePointerCapture", pointerId)
}
func (val HTMLElement) Remove()  {
	js.Value(val).Call("remove")
}
func (val HTMLElement) RemoveAttribute(attrName string)  {
	js.Value(val).Call("removeAttribute", attrName)
}
func (val HTMLElement) RemoveAttributeNS(namespaceURI string, localName string)  {
	js.Value(val).Call("removeAttributeNS", namespaceURI, localName)
}
func (val HTMLElement) RemoveEventListener(eventType string, listener func(event Event), options AddEventListenerOptions, useCapture bool)  {
	js.Value(val).Call("removeEventListener", eventType, listener, options, useCapture)
}
func (val HTMLElement) ReplaceChildren(node Node)  {
	js.Value(val).Call("replaceChildren", node)
}
func (val HTMLElement) ReplaceWith(node Node)  {
	js.Value(val).Call("replaceWith", node)
}
func (val HTMLElement) Scroll(options ScrollOptions)  {
	js.Value(val).Call("scroll", options)
}
func (val HTMLElement) ScrollBy(options ScrollOptions)  {
	js.Value(val).Call("scrollBy", options)
}
func (val HTMLElement) ScrollIntoView(options ScrollIntoViewOptions)  {
	js.Value(val).Call("scrollIntoView", options)
}
func (val HTMLElement) ScrollTo(options ScrollOptions)  {
	js.Value(val).Call("scrollTo", options)
}
func (val HTMLElement) SetAttribute(name string, value string)  {
	js.Value(val).Call("setAttribute", name, value)
}
func (val HTMLElement) SetAttributeNS(namespace string, name string, value string)  {
	js.Value(val).Call("setAttributeNS", namespace, name, value)
}
func (val HTMLElement) SetPointerCapture(pointerId int)  {
	js.Value(val).Call("setPointerCapture", pointerId)
}
func (val HTMLElement) ToggleAttribute(name string)  {
	js.Value(val).Call("toggleAttribute", name)
}
func (val HTMLElement) Blur()  {
	js.Value(val).Call("blur")
}
func (val HTMLElement) Click()  {
	js.Value(val).Call("click")
}
func (val HTMLElement) Focus()  {
	js.Value(val).Call("focus")
}
func (val HTMLElement) FocusWithOptions(options FocusOptions)  {
	js.Value(val).Call("focus", options)
}

type CSSStyleDeclaration js.Value

func wrapCSSStyleDeclaration(value js.Value) CSSStyleDeclaration {
	return CSSStyleDeclaration(value)
}
func (val CSSStyleDeclaration) CssText() string {
	return js.Value(val).Call("cssText").String()
}
func (val CSSStyleDeclaration) Length() int {
	return js.Value(val).Length()
}
func (val CSSStyleDeclaration) GetPropertyPriority(property string) string {
	return js.Value(val).Call("getPropertyPriority", property).String()
}
func (val CSSStyleDeclaration) GetPropertyValue(property string) string {
	return js.Value(val).Call("getPropertyValue", property).String()
}
func (val CSSStyleDeclaration) Item(index int) string {
	return js.Value(val).Call("item", index).String()
}
func (val CSSStyleDeclaration) RemoveProperty(property string)  {
	js.Value(val).Call("removeProperty", property)
}
func (val CSSStyleDeclaration) SetProperty(property string, value string)  {
	js.Value(val).Call("setProperty", property, value)
}
func (val CSSStyleDeclaration) SetPropertyWithPriority(property string, value string, priority string)  {
	js.Value(val).Call("setProperty", property, value, priority)
}

type DocumentFragment js.Value

func wrapDocumentFragment(value js.Value) DocumentFragment {
	return DocumentFragment(value)
}
func (val DocumentFragment) BaseURI() string {
	return js.Value(val).Call("baseURI").String()
}
func (val DocumentFragment) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Call("childNodes"))
}
func (val DocumentFragment) FirstChild() Node {
	return wrapNode(js.Value(val).Call("firstChild"))
}
func (val DocumentFragment) IsConnected() bool {
	return js.Value(val).Call("isConnected").Bool()
}
func (val DocumentFragment) LastChild() Node {
	return wrapNode(js.Value(val).Call("lastChild"))
}
func (val DocumentFragment) NextSibling() Node {
	return wrapNode(js.Value(val).Call("nextSibling"))
}
func (val DocumentFragment) NodeName() string {
	return js.Value(val).Call("nodeName").String()
}
func (val DocumentFragment) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Call("nodeType"))
}
func (val DocumentFragment) NodeValue() Node {
	return wrapNode(js.Value(val).Call("nodeValue"))
}
func (val DocumentFragment) ParentNode() Node {
	return wrapNode(js.Value(val).Call("parentNode"))
}
func (val DocumentFragment) ParentElement() Element {
	return wrapElement(js.Value(val).Call("parentElement"))
}
func (val DocumentFragment) PreviousSibling() Node {
	return wrapNode(js.Value(val).Call("previousSibling"))
}
func (val DocumentFragment) TextContent() string {
	return js.Value(val).Call("textContent").String()
}
func (val DocumentFragment) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", aChild))
}
func (val DocumentFragment) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val DocumentFragment) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", otherNode))
}
func (val DocumentFragment) Contains(node Node) bool {
	return js.Value(val).Call("contains", node).Bool()
}
func (val DocumentFragment) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", options))
}
func (val DocumentFragment) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val DocumentFragment) InsertBefore(newNode Node, referenceNode Node) bool {
	return js.Value(val).Call("insertBefore", newNode, referenceNode).Bool()
}
func (val DocumentFragment) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val DocumentFragment) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", otherNode).Bool()
}
func (val DocumentFragment) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", otherNode).Bool()
}
func (val DocumentFragment) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val DocumentFragment) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val DocumentFragment) Normalize()  {
	js.Value(val).Call("normalize")
}
func (val DocumentFragment) RemoveChild(child Node)  {
	js.Value(val).Call("removeChild", child)
}
func (val DocumentFragment) ReplaceChild(newChild Node, oldChild Node)  {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}

type SVGElement js.Value

func wrapSVGElement(value js.Value) SVGElement {
	return SVGElement(value)
}
func (val SVGElement) BaseURI() string {
	return js.Value(val).Call("baseURI").String()
}
func (val SVGElement) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Call("childNodes"))
}
func (val SVGElement) FirstChild() Node {
	return wrapNode(js.Value(val).Call("firstChild"))
}
func (val SVGElement) IsConnected() bool {
	return js.Value(val).Call("isConnected").Bool()
}
func (val SVGElement) LastChild() Node {
	return wrapNode(js.Value(val).Call("lastChild"))
}
func (val SVGElement) NextSibling() Node {
	return wrapNode(js.Value(val).Call("nextSibling"))
}
func (val SVGElement) NodeName() string {
	return js.Value(val).Call("nodeName").String()
}
func (val SVGElement) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Call("nodeType"))
}
func (val SVGElement) NodeValue() Node {
	return wrapNode(js.Value(val).Call("nodeValue"))
}
func (val SVGElement) ParentNode() Node {
	return wrapNode(js.Value(val).Call("parentNode"))
}
func (val SVGElement) ParentElement() Element {
	return wrapElement(js.Value(val).Call("parentElement"))
}
func (val SVGElement) PreviousSibling() Node {
	return wrapNode(js.Value(val).Call("previousSibling"))
}
func (val SVGElement) TextContent() string {
	return js.Value(val).Call("textContent").String()
}
func (val SVGElement) ChildElementCount() int {
	return js.Value(val).Call("childElementCount").Int()
}
func (val SVGElement) Children() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("children"))
}
func (val SVGElement) ClassList() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Call("classList"))
}
func (val SVGElement) ClassName() string {
	return js.Value(val).Call("className").String()
}
func (val SVGElement) ClientHeight() int {
	return js.Value(val).Call("clientHeight").Int()
}
func (val SVGElement) ClientLeft() int {
	return js.Value(val).Call("clientLeft").Int()
}
func (val SVGElement) ClientTop() int {
	return js.Value(val).Call("clientTop").Int()
}
func (val SVGElement) ClientWidth() int {
	return js.Value(val).Call("clientWidth").Int()
}
func (val SVGElement) FirstElementChild() Element {
	return wrapElement(js.Value(val).Call("firstElementChild"))
}
func (val SVGElement) ID() string {
	return js.Value(val).Call("id").String()
}
func (val SVGElement) InnerHTML() string {
	return js.Value(val).Call("innerHTML").String()
}
func (val SVGElement) LocalName() string {
	return js.Value(val).Call("localName").String()
}
func (val SVGElement) NamespaceURI() string {
	return js.Value(val).Call("namespaceURI").String()
}
func (val SVGElement) NextElementSibling() Element {
	return wrapElement(js.Value(val).Call("nextElementSibling"))
}
func (val SVGElement) OuterHTML() string {
	return js.Value(val).Call("outerHTML").String()
}
func (val SVGElement) Part() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Call("part"))
}
func (val SVGElement) Prefix() string {
	return js.Value(val).Call("prefix").String()
}
func (val SVGElement) PreviousElementSibling() Element {
	return wrapElement(js.Value(val).Call("previousElementSibling"))
}
func (val SVGElement) ScrollHeight() int {
	return js.Value(val).Call("scrollHeight").Int()
}
func (val SVGElement) ScrollLeft() int {
	return js.Value(val).Call("scrollLeft").Int()
}
func (val SVGElement) ScrollTop() int {
	return js.Value(val).Call("scrollTop").Int()
}
func (val SVGElement) ScrollWidth() int {
	return js.Value(val).Call("scrollWidth").Int()
}
func (val SVGElement) Slot() string {
	return js.Value(val).Call("slot").String()
}
func (val SVGElement) TagName() string {
	return js.Value(val).Call("tagName").String()
}
func (val SVGElement) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", aChild))
}
func (val SVGElement) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val SVGElement) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", otherNode))
}
func (val SVGElement) Contains(node Node) bool {
	return js.Value(val).Call("contains", node).Bool()
}
func (val SVGElement) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", options))
}
func (val SVGElement) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val SVGElement) InsertBefore(newNode Node, referenceNode Node) bool {
	return js.Value(val).Call("insertBefore", newNode, referenceNode).Bool()
}
func (val SVGElement) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val SVGElement) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", otherNode).Bool()
}
func (val SVGElement) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", otherNode).Bool()
}
func (val SVGElement) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val SVGElement) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val SVGElement) Normalize()  {
	js.Value(val).Call("normalize")
}
func (val SVGElement) RemoveChild(child Node)  {
	js.Value(val).Call("removeChild", child)
}
func (val SVGElement) ReplaceChild(newChild Node, oldChild Node)  {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}
func (val SVGElement) AddEventListener(eventType string, listener func(event Event), options AddEventListenerOptions, useCapture bool)  {
	js.Value(val).Call("addEventListener", eventType, listener, options, useCapture)
}
func (val SVGElement) After(node Node)  {
	js.Value(val).Call("after", node)
}
func (val SVGElement) Append(params Node)  {
	js.Value(val).Call("append", params)
}
func (val SVGElement) Before(params Node)  {
	js.Value(val).Call("before", params)
}
func (val SVGElement) Closest(selector string) Element {
	return wrapElement(js.Value(val).Call("closest", selector))
}
func (val SVGElement) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val SVGElement) GetAttribute(name string) string {
	return js.Value(val).Call("getAttribute", name).String()
}
func (val SVGElement) GetAttributeNames() string {
	return js.Value(val).Call("getAttributeNames").String()
}
func (val SVGElement) GetBoundingClientRect() DOMRect {
	return wrapDOMRect(js.Value(val).Call("getBoundingClientRect"))
}
func (val SVGElement) GetElementsByClassName(names string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByClassName", names))
}
func (val SVGElement) GetElementsByTagName(tagName string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByTagName", tagName))
}
func (val SVGElement) GetElementsByTagNameNS(namespaceURI string, localName string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByTagNameNS", namespaceURI, localName))
}
func (val SVGElement) HasAttribute(name string) bool {
	return js.Value(val).Call("hasAttribute", name).Bool()
}
func (val SVGElement) HasAttributeNS(namespaceURI string, localName string) bool {
	return js.Value(val).Call("hasAttributeNS", namespaceURI, localName).Bool()
}
func (val SVGElement) HasAttributes() bool {
	return js.Value(val).Call("hasAttributes").Bool()
}
func (val SVGElement) HasPointerCapture(pointerId int) bool {
	return js.Value(val).Call("hasPointerCapture", pointerId).Bool()
}
func (val SVGElement) InsertAdjacentElement(position AdjacentPosition, element Element) Element {
	return wrapElement(js.Value(val).Call("insertAdjacentElement", position, element))
}
func (val SVGElement) InsertAdjacentHTML(position AdjacentPosition, text string)  {
	js.Value(val).Call("insertAdjacentHTML", position, text)
}
func (val SVGElement) InsertAdjacentText(position AdjacentPosition, text string)  {
	js.Value(val).Call("insertAdjacentText", position, text)
}
func (val SVGElement) Matches(selectors string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("matches", selectors))
}
func (val SVGElement) Prepend(params Node)  {
	js.Value(val).Call("prepend", params)
}
func (val SVGElement) QuerySelector(selectors string) Element {
	return wrapElement(js.Value(val).Call("querySelector", selectors))
}
func (val SVGElement) QuerySelectorAll(selectors string) NodeList {
	return wrapNodeList(js.Value(val).Call("querySelectorAll", selectors))
}
func (val SVGElement) ReleasePointerCapture(pointerId int)  {
	js.Value(val).Call("releasePointerCapture", pointerId)
}
func (val SVGElement) Remove()  {
	js.Value(val).Call("remove")
}
func (val SVGElement) RemoveAttribute(attrName string)  {
	js.Value(val).Call("removeAttribute", attrName)
}
func (val SVGElement) RemoveAttributeNS(namespaceURI string, localName string)  {
	js.Value(val).Call("removeAttributeNS", namespaceURI, localName)
}
func (val SVGElement) RemoveEventListener(eventType string, listener func(event Event), options AddEventListenerOptions, useCapture bool)  {
	js.Value(val).Call("removeEventListener", eventType, listener, options, useCapture)
}
func (val SVGElement) ReplaceChildren(node Node)  {
	js.Value(val).Call("replaceChildren", node)
}
func (val SVGElement) ReplaceWith(node Node)  {
	js.Value(val).Call("replaceWith", node)
}
func (val SVGElement) Scroll(options ScrollOptions)  {
	js.Value(val).Call("scroll", options)
}
func (val SVGElement) ScrollBy(options ScrollOptions)  {
	js.Value(val).Call("scrollBy", options)
}
func (val SVGElement) ScrollIntoView(options ScrollIntoViewOptions)  {
	js.Value(val).Call("scrollIntoView", options)
}
func (val SVGElement) ScrollTo(options ScrollOptions)  {
	js.Value(val).Call("scrollTo", options)
}
func (val SVGElement) SetAttribute(name string, value string)  {
	js.Value(val).Call("setAttribute", name, value)
}
func (val SVGElement) SetAttributeNS(namespace string, name string, value string)  {
	js.Value(val).Call("setAttributeNS", namespace, name, value)
}
func (val SVGElement) SetPointerCapture(pointerId int)  {
	js.Value(val).Call("setPointerCapture", pointerId)
}
func (val SVGElement) ToggleAttribute(name string)  {
	js.Value(val).Call("toggleAttribute", name)
}

package dom

import (
	"syscall/js"
	"time"
)

type Event js.Value

func wrapEvent(value js.Value) Event {
	return Event(value)
}
func (val Event) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val Event) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val Event) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val Event) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val Event) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val Event) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val Event) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val Event) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val Event) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val Event) Type() string {
	return js.Value(val).Get("type").String()
}
func (val Event) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val Event) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val Event) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}

type EventTarget interface {
	AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	DispatchEvent(event Event) bool
}
type GenericEvent js.Value

func wrapGenericEvent(value js.Value) GenericEvent {
	return GenericEvent(value)
}
func (val GenericEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val GenericEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val GenericEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val GenericEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val GenericEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val GenericEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val GenericEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val GenericEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val GenericEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val GenericEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val GenericEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val GenericEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val GenericEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}

type UIEvent js.Value

func wrapUIEvent(value js.Value) UIEvent {
	return UIEvent(value)
}
func (val UIEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val UIEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val UIEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val UIEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val UIEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val UIEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val UIEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val UIEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val UIEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val UIEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val UIEvent) View() js.Value {
	return js.Value(val).Get("view")
}
func (val UIEvent) Detail() int {
	return js.Value(val).Get("detail").Int()
}
func (val UIEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val UIEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val UIEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}

type MouseEvent js.Value

func wrapMouseEvent(value js.Value) MouseEvent {
	return MouseEvent(value)
}
func (val MouseEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val MouseEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val MouseEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val MouseEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val MouseEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val MouseEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val MouseEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val MouseEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val MouseEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val MouseEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val MouseEvent) View() js.Value {
	return js.Value(val).Get("view")
}
func (val MouseEvent) Detail() int {
	return js.Value(val).Get("detail").Int()
}
func (val MouseEvent) AltKey() bool {
	return js.Value(val).Get("altKey").Bool()
}
func (val MouseEvent) Button() int {
	return js.Value(val).Get("button").Int()
}
func (val MouseEvent) Buttons() int {
	return js.Value(val).Get("buttons").Int()
}
func (val MouseEvent) ClientX() float64 {
	return js.Value(val).Get("clientX").Float()
}
func (val MouseEvent) ClientY() float64 {
	return js.Value(val).Get("clientY").Float()
}
func (val MouseEvent) CtrlKey() bool {
	return js.Value(val).Get("ctrlKey").Bool()
}
func (val MouseEvent) MetaKey() bool {
	return js.Value(val).Get("metaKey").Bool()
}
func (val MouseEvent) MovementX() int {
	return js.Value(val).Get("movementX").Int()
}
func (val MouseEvent) MovementY() int {
	return js.Value(val).Get("movementY").Int()
}
func (val MouseEvent) OffsetX() float64 {
	return js.Value(val).Get("offsetX").Float()
}
func (val MouseEvent) OffsetY() float64 {
	return js.Value(val).Get("offsetY").Float()
}
func (val MouseEvent) PageX() float64 {
	return js.Value(val).Get("pageX").Float()
}
func (val MouseEvent) PageY() float64 {
	return js.Value(val).Get("pageY").Float()
}
func (val MouseEvent) RelatedTarget() js.Value {
	return js.Value(val).Get("relatedTarget")
}
func (val MouseEvent) ShiftKey() bool {
	return js.Value(val).Get("shiftKey").Bool()
}
func (val MouseEvent) X() float64 {
	return js.Value(val).Get("x").Float()
}
func (val MouseEvent) Y() float64 {
	return js.Value(val).Get("y").Float()
}
func (val MouseEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val MouseEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val MouseEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}
func (val MouseEvent) GetModifierState(key int) bool {
	return js.Value(val).Call("getModifierState", key).Bool()
}

type FocusEvent js.Value

func wrapFocusEvent(value js.Value) FocusEvent {
	return FocusEvent(value)
}
func (val FocusEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val FocusEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val FocusEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val FocusEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val FocusEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val FocusEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val FocusEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val FocusEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val FocusEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val FocusEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val FocusEvent) View() js.Value {
	return js.Value(val).Get("view")
}
func (val FocusEvent) Detail() int {
	return js.Value(val).Get("detail").Int()
}
func (val FocusEvent) RelatedTarget() js.Value {
	return js.Value(val).Get("relatedTarget")
}
func (val FocusEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val FocusEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val FocusEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}

type InputEvent js.Value

func wrapInputEvent(value js.Value) InputEvent {
	return InputEvent(value)
}
func (val InputEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val InputEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val InputEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val InputEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val InputEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val InputEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val InputEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val InputEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val InputEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val InputEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val InputEvent) View() js.Value {
	return js.Value(val).Get("view")
}
func (val InputEvent) Detail() int {
	return js.Value(val).Get("detail").Int()
}
func (val InputEvent) Data() string {
	return js.Value(val).Get("data").String()
}
func (val InputEvent) DataTransfer() string {
	return js.Value(val).Get("dataTransfer").String()
}
func (val InputEvent) InputType() string {
	return js.Value(val).Get("inputType").String()
}
func (val InputEvent) IsComposing() bool {
	return js.Value(val).Get("isComposing").Bool()
}
func (val InputEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val InputEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val InputEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}
func (val InputEvent) GetTargetRanges() []StaticRange {
	return makeStaticRangeSlice(js.Value(val).Call("getTargetRanges"))
}

type MessageEvent js.Value

func wrapMessageEvent(value js.Value) MessageEvent {
	return MessageEvent(value)
}
func (val MessageEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val MessageEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val MessageEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val MessageEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val MessageEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val MessageEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val MessageEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val MessageEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val MessageEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val MessageEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val MessageEvent) Data() string {
	return js.Value(val).Get("data").String()
}
func (val MessageEvent) Origin() string {
	return js.Value(val).Get("origin").String()
}
func (val MessageEvent) LastEventID() string {
	return js.Value(val).Get("lastEventId").String()
}
func (val MessageEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val MessageEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val MessageEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}

type KeyboardEvent js.Value

func wrapKeyboardEvent(value js.Value) KeyboardEvent {
	return KeyboardEvent(value)
}
func (val KeyboardEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val KeyboardEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val KeyboardEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val KeyboardEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val KeyboardEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val KeyboardEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val KeyboardEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val KeyboardEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val KeyboardEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val KeyboardEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val KeyboardEvent) View() js.Value {
	return js.Value(val).Get("view")
}
func (val KeyboardEvent) Detail() int {
	return js.Value(val).Get("detail").Int()
}
func (val KeyboardEvent) AltKey() bool {
	return js.Value(val).Get("altKey").Bool()
}
func (val KeyboardEvent) Code() string {
	return js.Value(val).Get("code").String()
}
func (val KeyboardEvent) CtrlKey() bool {
	return js.Value(val).Get("ctrlKey").Bool()
}
func (val KeyboardEvent) IsComposing() bool {
	return js.Value(val).Get("isComposing").Bool()
}
func (val KeyboardEvent) Key() string {
	return js.Value(val).Get("key").String()
}
func (val KeyboardEvent) Location() string {
	return js.Value(val).Get("location").String()
}
func (val KeyboardEvent) MetaKey() bool {
	return js.Value(val).Get("metaKey").Bool()
}
func (val KeyboardEvent) Repeat() bool {
	return js.Value(val).Get("repeat").Bool()
}
func (val KeyboardEvent) ShiftKey() bool {
	return js.Value(val).Get("shiftKey").Bool()
}
func (val KeyboardEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val KeyboardEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val KeyboardEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
}
func (val KeyboardEvent) GetModifierState() bool {
	return js.Value(val).Call("getModifierState").Bool()
}

type AbstractRange interface {
	Collapsed() bool
	EndContainer() Node
	EndOffset() int
	StartContainer() Node
	StartOffset() int
}
type StaticRange js.Value

func wrapStaticRange(value js.Value) StaticRange {
	return StaticRange(value)
}
func (val StaticRange) Collapsed() bool {
	return js.Value(val).Get("collapsed").Bool()
}
func (val StaticRange) EndContainer() Node {
	return wrapNode(js.Value(val).Get("endContainer"))
}
func (val StaticRange) EndOffset() int {
	return js.Value(val).Get("endOffset").Int()
}
func (val StaticRange) StartContainer() Node {
	return wrapNode(js.Value(val).Get("startContainer"))
}
func (val StaticRange) StartOffset() int {
	return js.Value(val).Get("startOffset").Int()
}

type SecurityPolicyViolationEvent js.Value

func wrapSecurityPolicyViolationEvent(value js.Value) SecurityPolicyViolationEvent {
	return SecurityPolicyViolationEvent(value)
}
func (val SecurityPolicyViolationEvent) Bubbles() bool {
	return js.Value(val).Get("bubbles").Bool()
}
func (val SecurityPolicyViolationEvent) Cancelable() bool {
	return js.Value(val).Get("cancelable").Bool()
}
func (val SecurityPolicyViolationEvent) Composed() bool {
	return js.Value(val).Get("composed").Bool()
}
func (val SecurityPolicyViolationEvent) CurrentTarget() js.Value {
	return js.Value(val).Get("currentTarget")
}
func (val SecurityPolicyViolationEvent) DefaultPrevented() bool {
	return js.Value(val).Get("defaultPrevented").Bool()
}
func (val SecurityPolicyViolationEvent) EventPhase() EventPhase {
	return wrapEventPhase(js.Value(val).Get("eventPhase"))
}
func (val SecurityPolicyViolationEvent) IsTrusted() bool {
	return js.Value(val).Get("isTrusted").Bool()
}
func (val SecurityPolicyViolationEvent) Target() js.Value {
	return js.Value(val).Get("target")
}
func (val SecurityPolicyViolationEvent) TimeStamp() time.Time {
	return millisecondsSinceEpocToTime(js.Value(val).Get("timeStamp"))
}
func (val SecurityPolicyViolationEvent) Type() string {
	return js.Value(val).Get("type").String()
}
func (val SecurityPolicyViolationEvent) BlockedURI() string {
	return js.Value(val).Get("blockedURI").String()
}
func (val SecurityPolicyViolationEvent) ColumnNumber() int {
	return js.Value(val).Get("columnNumber").Int()
}
func (val SecurityPolicyViolationEvent) Disposition() string {
	return js.Value(val).Get("disposition").String()
}
func (val SecurityPolicyViolationEvent) DocumentURI() string {
	return js.Value(val).Get("documentURI").String()
}
func (val SecurityPolicyViolationEvent) EffectiveDirective() string {
	return js.Value(val).Get("effectiveDirective").String()
}
func (val SecurityPolicyViolationEvent) LineNumber() int {
	return js.Value(val).Get("lineNumber").Int()
}
func (val SecurityPolicyViolationEvent) OriginalPolicy() string {
	return js.Value(val).Get("originalPolicy").String()
}
func (val SecurityPolicyViolationEvent) Referrer() string {
	return js.Value(val).Get("referrer").String()
}
func (val SecurityPolicyViolationEvent) SourceFile() string {
	return js.Value(val).Get("sourceFile").String()
}
func (val SecurityPolicyViolationEvent) StatusCode() int {
	return js.Value(val).Get("statusCode").Int()
}
func (val SecurityPolicyViolationEvent) ViolatedDirective() string {
	return js.Value(val).Get("violatedDirective").String()
}
func (val SecurityPolicyViolationEvent) PreventDefault() {
	js.Value(val).Call("preventDefault")
}
func (val SecurityPolicyViolationEvent) StopImmediatePropagation() {
	js.Value(val).Call("stopImmediatePropagation")
}
func (val SecurityPolicyViolationEvent) StopPropagation() {
	js.Value(val).Call("stopPropagation")
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
	SetTextContent(textContent string)
	AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	DispatchEvent(event Event) bool
	AppendChild(aChild Node) Node
	CloneNode(deep bool) Node
	CompareDocumentPosition(otherNode Node) DocumentPosition
	Contains(node Node) bool
	GetRootNode(options GetRootNodeOptions) Node
	HasChildNodes() bool
	InsertBefore(newNode Node, referenceNode Node) Element
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
	SetTextContent(textContent string)
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
	SetInnerHTML(innerHTML string)
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
	AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	DispatchEvent(event Event) bool
	AppendChild(aChild Node) Node
	CloneNode(deep bool) Node
	CompareDocumentPosition(otherNode Node) DocumentPosition
	Contains(node Node) bool
	GetRootNode(options GetRootNodeOptions) Node
	HasChildNodes() bool
	InsertBefore(newNode Node, referenceNode Node) Element
	IsDefaultNamespace(namespaceURL string) bool
	IsEqualNode(otherNode Node) bool
	IsOtherNode(otherNode Node) bool
	LookupPrefix(prefix string) string
	LookupNamespaceURI(prefix string) string
	Normalize()
	RemoveChild(child Node)
	ReplaceChild(newChild Node, oldChild Node)
	After(node ...Node)
	Append(params ...Node)
	Before(params ...Node)
	Closest(selector string) Element
	GetAttribute(name string) string
	GetAttributeNames() string
	GetBoundingClientRect() DOMRect
	GetElementsByClassName(names string) HTMLCollection
	GetElementByID(id string) Element
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
	Prepend(nodes ...Node)
	QuerySelector(selectors string) Element
	QuerySelectorAll(selectors string) NodeList
	ReleasePointerCapture(pointerId int)
	Remove()
	RemoveAttribute(attrName string)
	RemoveAttributeNS(namespaceURI string, localName string)
	ReplaceChildren(node ...Node)
	ReplaceWith(node ...Node)
	Scroll(options ScrollOptions)
	ScrollBy(options ScrollOptions)
	ScrollIntoView(options ScrollIntoViewOptions)
	ScrollTo(options ScrollOptions)
	SetAttribute(name string, value string)
	SetAttributeNS(namespace string, name string, value string)
	SetPointerCapture(pointerId int)
	ToggleAttribute(name string)
}
type HTMLIFrameElement js.Value

func wrapHTMLIFrameElement(value js.Value) HTMLIFrameElement {
	return HTMLIFrameElement(value)
}
func (val HTMLIFrameElement) BaseURI() string {
	return js.Value(val).Get("baseURI").String()
}
func (val HTMLIFrameElement) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Get("childNodes"))
}
func (val HTMLIFrameElement) FirstChild() Node {
	return wrapNode(js.Value(val).Get("firstChild"))
}
func (val HTMLIFrameElement) IsConnected() bool {
	return js.Value(val).Get("isConnected").Bool()
}
func (val HTMLIFrameElement) LastChild() Node {
	return wrapNode(js.Value(val).Get("lastChild"))
}
func (val HTMLIFrameElement) NextSibling() Node {
	return wrapNode(js.Value(val).Get("nextSibling"))
}
func (val HTMLIFrameElement) NodeName() string {
	return js.Value(val).Get("nodeName").String()
}
func (val HTMLIFrameElement) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Get("nodeType"))
}
func (val HTMLIFrameElement) NodeValue() Node {
	return wrapNode(js.Value(val).Get("nodeValue"))
}
func (val HTMLIFrameElement) ParentNode() Node {
	return wrapNode(js.Value(val).Get("parentNode"))
}
func (val HTMLIFrameElement) ParentElement() Element {
	return wrapElement(js.Value(val).Get("parentElement"))
}
func (val HTMLIFrameElement) PreviousSibling() Node {
	return wrapNode(js.Value(val).Get("previousSibling"))
}
func (val HTMLIFrameElement) TextContent() string {
	return js.Value(val).Get("textContent").String()
}
func (val HTMLIFrameElement) SetTextContent(textContent string) {
	js.Value(val).Set("textContent", textContent)
}
func (val HTMLIFrameElement) ChildElementCount() int {
	return js.Value(val).Get("childElementCount").Int()
}
func (val HTMLIFrameElement) Children() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("children"))
}
func (val HTMLIFrameElement) ClassList() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Get("classList"))
}
func (val HTMLIFrameElement) ClassName() string {
	return js.Value(val).Get("className").String()
}
func (val HTMLIFrameElement) ClientHeight() int {
	return js.Value(val).Get("clientHeight").Int()
}
func (val HTMLIFrameElement) ClientLeft() int {
	return js.Value(val).Get("clientLeft").Int()
}
func (val HTMLIFrameElement) ClientTop() int {
	return js.Value(val).Get("clientTop").Int()
}
func (val HTMLIFrameElement) ClientWidth() int {
	return js.Value(val).Get("clientWidth").Int()
}
func (val HTMLIFrameElement) FirstElementChild() Element {
	return wrapElement(js.Value(val).Get("firstElementChild"))
}
func (val HTMLIFrameElement) ID() string {
	return js.Value(val).Get("id").String()
}
func (val HTMLIFrameElement) InnerHTML() string {
	return js.Value(val).Get("innerHTML").String()
}
func (val HTMLIFrameElement) SetInnerHTML(innerHTML string) {
	js.Value(val).Set("innerHTML", innerHTML)
}
func (val HTMLIFrameElement) LocalName() string {
	return js.Value(val).Get("localName").String()
}
func (val HTMLIFrameElement) NamespaceURI() string {
	return js.Value(val).Get("namespaceURI").String()
}
func (val HTMLIFrameElement) NextElementSibling() Element {
	return wrapElement(js.Value(val).Get("nextElementSibling"))
}
func (val HTMLIFrameElement) OuterHTML() string {
	return js.Value(val).Get("outerHTML").String()
}
func (val HTMLIFrameElement) Part() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Get("part"))
}
func (val HTMLIFrameElement) Prefix() string {
	return js.Value(val).Get("prefix").String()
}
func (val HTMLIFrameElement) PreviousElementSibling() Element {
	return wrapElement(js.Value(val).Get("previousElementSibling"))
}
func (val HTMLIFrameElement) ScrollHeight() int {
	return js.Value(val).Get("scrollHeight").Int()
}
func (val HTMLIFrameElement) ScrollLeft() int {
	return js.Value(val).Get("scrollLeft").Int()
}
func (val HTMLIFrameElement) ScrollTop() int {
	return js.Value(val).Get("scrollTop").Int()
}
func (val HTMLIFrameElement) ScrollWidth() int {
	return js.Value(val).Get("scrollWidth").Int()
}
func (val HTMLIFrameElement) Slot() string {
	return js.Value(val).Get("slot").String()
}
func (val HTMLIFrameElement) TagName() string {
	return js.Value(val).Get("tagName").String()
}
func (val HTMLIFrameElement) AccessKey() string {
	return js.Value(val).Get("accessKey").String()
}
func (val HTMLIFrameElement) AccessKeyLabel() string {
	return js.Value(val).Get("accessKeyLabel").String()
}
func (val HTMLIFrameElement) IsContentEditable() bool {
	return js.Value(val).Get("isContentEditable").Bool()
}
func (val HTMLIFrameElement) Dataset() StringMap {
	return wrapStringMap(js.Value(val).Get("dataset"))
}
func (val HTMLIFrameElement) Dir() string {
	return js.Value(val).Get("dir").String()
}
func (val HTMLIFrameElement) Draggable() bool {
	return js.Value(val).Get("draggable").Bool()
}
func (val HTMLIFrameElement) EnterKeyHint() string {
	return js.Value(val).Get("enterKeyHint").String()
}
func (val HTMLIFrameElement) Hidden() bool {
	return js.Value(val).Get("hidden").Bool()
}
func (val HTMLIFrameElement) SetHidden(hidden bool) {
	js.Value(val).Set("hidden", hidden)
}
func (val HTMLIFrameElement) InnerText() string {
	return js.Value(val).Get("innerText").String()
}
func (val HTMLIFrameElement) SetInnerText(innerText string) {
	js.Value(val).Set("innerText", innerText)
}
func (val HTMLIFrameElement) Lang() string {
	return js.Value(val).Get("lang").String()
}
func (val HTMLIFrameElement) SetLang(lang string) {
	js.Value(val).Set("lang", lang)
}
func (val HTMLIFrameElement) OffsetHeight() float64 {
	return js.Value(val).Get("offsetHeight").Float()
}
func (val HTMLIFrameElement) OffsetLeft() float64 {
	return js.Value(val).Get("offsetLeft").Float()
}
func (val HTMLIFrameElement) OffsetParent() float64 {
	return js.Value(val).Get("offsetParent").Float()
}
func (val HTMLIFrameElement) OffsetTop() float64 {
	return js.Value(val).Get("offsetTop").Float()
}
func (val HTMLIFrameElement) OffsetWidth() float64 {
	return js.Value(val).Get("offsetWidth").Float()
}
func (val HTMLIFrameElement) OuterText() string {
	return js.Value(val).Get("outerText").String()
}
func (val HTMLIFrameElement) SetOuterText(outerText string) {
	js.Value(val).Set("outerText", outerText)
}
func (val HTMLIFrameElement) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(js.Value(val).Get("style"))
}
func (val HTMLIFrameElement) TabIndex() int {
	return js.Value(val).Get("tabIndex").Int()
}
func (val HTMLIFrameElement) SetTabIndex(tabIndex int) {
	js.Value(val).Set("tabIndex", tabIndex)
}
func (val HTMLIFrameElement) Title() string {
	return js.Value(val).Get("title").String()
}
func (val HTMLIFrameElement) SetTitle(title string) {
	js.Value(val).Set("title", title)
}
func (val HTMLIFrameElement) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val HTMLIFrameElement) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val HTMLIFrameElement) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val HTMLIFrameElement) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", convertNodeToValue(aChild)))
}
func (val HTMLIFrameElement) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val HTMLIFrameElement) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", convertNodeToValue(otherNode)))
}
func (val HTMLIFrameElement) Contains(node Node) bool {
	return js.Value(val).Call("contains", convertNodeToValue(node)).Bool()
}
func (val HTMLIFrameElement) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", wrapOptions(options)))
}
func (val HTMLIFrameElement) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val HTMLIFrameElement) InsertBefore(newNode Node, referenceNode Node) Element {
	return wrapElement(js.Value(val).Call("insertBefore", convertNodeToValue(newNode), convertNodeToValue(referenceNode)))
}
func (val HTMLIFrameElement) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val HTMLIFrameElement) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", convertNodeToValue(otherNode)).Bool()
}
func (val HTMLIFrameElement) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", convertNodeToValue(otherNode)).Bool()
}
func (val HTMLIFrameElement) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val HTMLIFrameElement) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val HTMLIFrameElement) Normalize() {
	js.Value(val).Call("normalize")
}
func (val HTMLIFrameElement) RemoveChild(child Node) {
	js.Value(val).Call("removeChild", convertNodeToValue(child))
}
func (val HTMLIFrameElement) ReplaceChild(newChild Node, oldChild Node) {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}
func (val HTMLIFrameElement) After(node ...Node) {
	js.Value(val).Call("after", nodesToValues(node)...)
}
func (val HTMLIFrameElement) Append(params ...Node) {
	js.Value(val).Call("append", nodesToValues(params)...)
}
func (val HTMLIFrameElement) Before(params ...Node) {
	js.Value(val).Call("before", nodesToValues(params)...)
}
func (val HTMLIFrameElement) Closest(selector string) Element {
	return wrapElement(js.Value(val).Call("closest", selector))
}
func (val HTMLIFrameElement) GetAttribute(name string) string {
	return js.Value(val).Call("getAttribute", name).String()
}
func (val HTMLIFrameElement) GetAttributeNames() string {
	return js.Value(val).Call("getAttributeNames").String()
}
func (val HTMLIFrameElement) GetBoundingClientRect() DOMRect {
	return wrapDOMRect(js.Value(val).Call("getBoundingClientRect"))
}
func (val HTMLIFrameElement) GetElementsByClassName(names string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByClassName", names))
}
func (val HTMLIFrameElement) GetElementByID(id string) Element {
	return wrapElement(js.Value(val).Call("getElementById", id))
}
func (val HTMLIFrameElement) GetElementsByTagName(tagName string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByTagName", tagName))
}
func (val HTMLIFrameElement) GetElementsByTagNameNS(namespaceURI string, localName string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByTagNameNS", namespaceURI, localName))
}
func (val HTMLIFrameElement) HasAttribute(name string) bool {
	return js.Value(val).Call("hasAttribute", name).Bool()
}
func (val HTMLIFrameElement) HasAttributeNS(namespaceURI string, localName string) bool {
	return js.Value(val).Call("hasAttributeNS", namespaceURI, localName).Bool()
}
func (val HTMLIFrameElement) HasAttributes() bool {
	return js.Value(val).Call("hasAttributes").Bool()
}
func (val HTMLIFrameElement) HasPointerCapture(pointerId int) bool {
	return js.Value(val).Call("hasPointerCapture", pointerId).Bool()
}
func (val HTMLIFrameElement) InsertAdjacentElement(position AdjacentPosition, element Element) Element {
	return wrapElement(js.Value(val).Call("insertAdjacentElement", string(position), element))
}
func (val HTMLIFrameElement) InsertAdjacentHTML(position AdjacentPosition, text string) {
	js.Value(val).Call("insertAdjacentHTML", string(position), text)
}
func (val HTMLIFrameElement) InsertAdjacentText(position AdjacentPosition, text string) {
	js.Value(val).Call("insertAdjacentText", string(position), text)
}
func (val HTMLIFrameElement) Matches(selectors string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("matches", selectors))
}
func (val HTMLIFrameElement) Prepend(nodes ...Node) {
	js.Value(val).Call("prepend", nodesToValues(nodes)...)
}
func (val HTMLIFrameElement) QuerySelector(selectors string) Element {
	return wrapElement(js.Value(val).Call("querySelector", selectors))
}
func (val HTMLIFrameElement) QuerySelectorAll(selectors string) NodeList {
	return wrapNodeList(js.Value(val).Call("querySelectorAll", selectors))
}
func (val HTMLIFrameElement) ReleasePointerCapture(pointerId int) {
	js.Value(val).Call("releasePointerCapture", pointerId)
}
func (val HTMLIFrameElement) Remove() {
	js.Value(val).Call("remove")
}
func (val HTMLIFrameElement) RemoveAttribute(attrName string) {
	js.Value(val).Call("removeAttribute", attrName)
}
func (val HTMLIFrameElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Value(val).Call("removeAttributeNS", namespaceURI, localName)
}
func (val HTMLIFrameElement) ReplaceChildren(node ...Node) {
	js.Value(val).Call("replaceChildren", nodesToValues(node)...)
}
func (val HTMLIFrameElement) ReplaceWith(node ...Node) {
	js.Value(val).Call("replaceWith", nodesToValues(node)...)
}
func (val HTMLIFrameElement) Scroll(options ScrollOptions) {
	js.Value(val).Call("scroll", wrapOptions(options))
}
func (val HTMLIFrameElement) ScrollBy(options ScrollOptions) {
	js.Value(val).Call("scrollBy", wrapOptions(options))
}
func (val HTMLIFrameElement) ScrollIntoView(options ScrollIntoViewOptions) {
	js.Value(val).Call("scrollIntoView", wrapOptions(options))
}
func (val HTMLIFrameElement) ScrollTo(options ScrollOptions) {
	js.Value(val).Call("scrollTo", wrapOptions(options))
}
func (val HTMLIFrameElement) SetAttribute(name string, value string) {
	js.Value(val).Call("setAttribute", name, value)
}
func (val HTMLIFrameElement) SetAttributeNS(namespace string, name string, value string) {
	js.Value(val).Call("setAttributeNS", namespace, name, value)
}
func (val HTMLIFrameElement) SetPointerCapture(pointerId int) {
	js.Value(val).Call("setPointerCapture", pointerId)
}
func (val HTMLIFrameElement) ToggleAttribute(name string) {
	js.Value(val).Call("toggleAttribute", name)
}
func (val HTMLIFrameElement) Blur() {
	js.Value(val).Call("blur")
}
func (val HTMLIFrameElement) Click() {
	js.Value(val).Call("click")
}
func (val HTMLIFrameElement) Focus() {
	js.Value(val).Call("focus")
}
func (val HTMLIFrameElement) FocusWithOptions(options FocusOptions) {
	js.Value(val).Call("focus", wrapOptions(options))
}

type DOMRect js.Value

func wrapDOMRect(value js.Value) DOMRect {
	return DOMRect(value)
}
func (val DOMRect) X() int {
	return js.Value(val).Get("x").Int()
}
func (val DOMRect) Y() int {
	return js.Value(val).Get("y").Int()
}
func (val DOMRect) Width() int {
	return js.Value(val).Get("width").Int()
}
func (val DOMRect) Height() int {
	return js.Value(val).Get("height").Int()
}
func (val DOMRect) Top() int {
	return js.Value(val).Get("top").Int()
}
func (val DOMRect) Right() int {
	return js.Value(val).Get("right").Int()
}
func (val DOMRect) Bottom() int {
	return js.Value(val).Get("bottom").Int()
}
func (val DOMRect) Left() int {
	return js.Value(val).Get("left").Int()
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
	return js.Value(val).Get("value").String()
}
func (val DOMTokenList) Item(index int) string {
	return js.Value(val).Call("item", index).String()
}
func (val DOMTokenList) Contains(token string) bool {
	return js.Value(val).Call("contains", token).Bool()
}
func (val DOMTokenList) Add(token ...string) {
	js.Value(val).Call("add", stringsToAny(token)...)
}
func (val DOMTokenList) Remove(token ...string) {
	js.Value(val).Call("remove", stringsToAny(token)...)
}
func (val DOMTokenList) Replace(oldToken string, newToken string) {
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
	return js.Value(val).Get("baseURI").String()
}
func (val HTMLElement) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Get("childNodes"))
}
func (val HTMLElement) FirstChild() Node {
	return wrapNode(js.Value(val).Get("firstChild"))
}
func (val HTMLElement) IsConnected() bool {
	return js.Value(val).Get("isConnected").Bool()
}
func (val HTMLElement) LastChild() Node {
	return wrapNode(js.Value(val).Get("lastChild"))
}
func (val HTMLElement) NextSibling() Node {
	return wrapNode(js.Value(val).Get("nextSibling"))
}
func (val HTMLElement) NodeName() string {
	return js.Value(val).Get("nodeName").String()
}
func (val HTMLElement) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Get("nodeType"))
}
func (val HTMLElement) NodeValue() Node {
	return wrapNode(js.Value(val).Get("nodeValue"))
}
func (val HTMLElement) ParentNode() Node {
	return wrapNode(js.Value(val).Get("parentNode"))
}
func (val HTMLElement) ParentElement() Element {
	return wrapElement(js.Value(val).Get("parentElement"))
}
func (val HTMLElement) PreviousSibling() Node {
	return wrapNode(js.Value(val).Get("previousSibling"))
}
func (val HTMLElement) TextContent() string {
	return js.Value(val).Get("textContent").String()
}
func (val HTMLElement) SetTextContent(textContent string) {
	js.Value(val).Set("textContent", textContent)
}
func (val HTMLElement) ChildElementCount() int {
	return js.Value(val).Get("childElementCount").Int()
}
func (val HTMLElement) Children() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("children"))
}
func (val HTMLElement) ClassList() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Get("classList"))
}
func (val HTMLElement) ClassName() string {
	return js.Value(val).Get("className").String()
}
func (val HTMLElement) ClientHeight() int {
	return js.Value(val).Get("clientHeight").Int()
}
func (val HTMLElement) ClientLeft() int {
	return js.Value(val).Get("clientLeft").Int()
}
func (val HTMLElement) ClientTop() int {
	return js.Value(val).Get("clientTop").Int()
}
func (val HTMLElement) ClientWidth() int {
	return js.Value(val).Get("clientWidth").Int()
}
func (val HTMLElement) FirstElementChild() Element {
	return wrapElement(js.Value(val).Get("firstElementChild"))
}
func (val HTMLElement) ID() string {
	return js.Value(val).Get("id").String()
}
func (val HTMLElement) InnerHTML() string {
	return js.Value(val).Get("innerHTML").String()
}
func (val HTMLElement) SetInnerHTML(innerHTML string) {
	js.Value(val).Set("innerHTML", innerHTML)
}
func (val HTMLElement) LocalName() string {
	return js.Value(val).Get("localName").String()
}
func (val HTMLElement) NamespaceURI() string {
	return js.Value(val).Get("namespaceURI").String()
}
func (val HTMLElement) NextElementSibling() Element {
	return wrapElement(js.Value(val).Get("nextElementSibling"))
}
func (val HTMLElement) OuterHTML() string {
	return js.Value(val).Get("outerHTML").String()
}
func (val HTMLElement) Part() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Get("part"))
}
func (val HTMLElement) Prefix() string {
	return js.Value(val).Get("prefix").String()
}
func (val HTMLElement) PreviousElementSibling() Element {
	return wrapElement(js.Value(val).Get("previousElementSibling"))
}
func (val HTMLElement) ScrollHeight() int {
	return js.Value(val).Get("scrollHeight").Int()
}
func (val HTMLElement) ScrollLeft() int {
	return js.Value(val).Get("scrollLeft").Int()
}
func (val HTMLElement) ScrollTop() int {
	return js.Value(val).Get("scrollTop").Int()
}
func (val HTMLElement) ScrollWidth() int {
	return js.Value(val).Get("scrollWidth").Int()
}
func (val HTMLElement) Slot() string {
	return js.Value(val).Get("slot").String()
}
func (val HTMLElement) TagName() string {
	return js.Value(val).Get("tagName").String()
}
func (val HTMLElement) AccessKey() string {
	return js.Value(val).Get("accessKey").String()
}
func (val HTMLElement) AccessKeyLabel() string {
	return js.Value(val).Get("accessKeyLabel").String()
}
func (val HTMLElement) IsContentEditable() bool {
	return js.Value(val).Get("isContentEditable").Bool()
}
func (val HTMLElement) Dataset() StringMap {
	return wrapStringMap(js.Value(val).Get("dataset"))
}
func (val HTMLElement) Dir() string {
	return js.Value(val).Get("dir").String()
}
func (val HTMLElement) Draggable() bool {
	return js.Value(val).Get("draggable").Bool()
}
func (val HTMLElement) EnterKeyHint() string {
	return js.Value(val).Get("enterKeyHint").String()
}
func (val HTMLElement) Hidden() bool {
	return js.Value(val).Get("hidden").Bool()
}
func (val HTMLElement) SetHidden(hidden bool) {
	js.Value(val).Set("hidden", hidden)
}
func (val HTMLElement) InnerText() string {
	return js.Value(val).Get("innerText").String()
}
func (val HTMLElement) SetInnerText(innerText string) {
	js.Value(val).Set("innerText", innerText)
}
func (val HTMLElement) Lang() string {
	return js.Value(val).Get("lang").String()
}
func (val HTMLElement) SetLang(lang string) {
	js.Value(val).Set("lang", lang)
}
func (val HTMLElement) OffsetHeight() float64 {
	return js.Value(val).Get("offsetHeight").Float()
}
func (val HTMLElement) OffsetLeft() float64 {
	return js.Value(val).Get("offsetLeft").Float()
}
func (val HTMLElement) OffsetParent() float64 {
	return js.Value(val).Get("offsetParent").Float()
}
func (val HTMLElement) OffsetTop() float64 {
	return js.Value(val).Get("offsetTop").Float()
}
func (val HTMLElement) OffsetWidth() float64 {
	return js.Value(val).Get("offsetWidth").Float()
}
func (val HTMLElement) OuterText() string {
	return js.Value(val).Get("outerText").String()
}
func (val HTMLElement) SetOuterText(outerText string) {
	js.Value(val).Set("outerText", outerText)
}
func (val HTMLElement) Style() CSSStyleDeclaration {
	return wrapCSSStyleDeclaration(js.Value(val).Get("style"))
}
func (val HTMLElement) TabIndex() int {
	return js.Value(val).Get("tabIndex").Int()
}
func (val HTMLElement) SetTabIndex(tabIndex int) {
	js.Value(val).Set("tabIndex", tabIndex)
}
func (val HTMLElement) Title() string {
	return js.Value(val).Get("title").String()
}
func (val HTMLElement) SetTitle(title string) {
	js.Value(val).Set("title", title)
}
func (val HTMLElement) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val HTMLElement) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val HTMLElement) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val HTMLElement) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", convertNodeToValue(aChild)))
}
func (val HTMLElement) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val HTMLElement) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", convertNodeToValue(otherNode)))
}
func (val HTMLElement) Contains(node Node) bool {
	return js.Value(val).Call("contains", convertNodeToValue(node)).Bool()
}
func (val HTMLElement) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", wrapOptions(options)))
}
func (val HTMLElement) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val HTMLElement) InsertBefore(newNode Node, referenceNode Node) Element {
	return wrapElement(js.Value(val).Call("insertBefore", convertNodeToValue(newNode), convertNodeToValue(referenceNode)))
}
func (val HTMLElement) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val HTMLElement) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", convertNodeToValue(otherNode)).Bool()
}
func (val HTMLElement) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", convertNodeToValue(otherNode)).Bool()
}
func (val HTMLElement) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val HTMLElement) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val HTMLElement) Normalize() {
	js.Value(val).Call("normalize")
}
func (val HTMLElement) RemoveChild(child Node) {
	js.Value(val).Call("removeChild", convertNodeToValue(child))
}
func (val HTMLElement) ReplaceChild(newChild Node, oldChild Node) {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}
func (val HTMLElement) After(node ...Node) {
	js.Value(val).Call("after", nodesToValues(node)...)
}
func (val HTMLElement) Append(params ...Node) {
	js.Value(val).Call("append", nodesToValues(params)...)
}
func (val HTMLElement) Before(params ...Node) {
	js.Value(val).Call("before", nodesToValues(params)...)
}
func (val HTMLElement) Closest(selector string) Element {
	return wrapElement(js.Value(val).Call("closest", selector))
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
func (val HTMLElement) GetElementByID(id string) Element {
	return wrapElement(js.Value(val).Call("getElementById", id))
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
	return wrapElement(js.Value(val).Call("insertAdjacentElement", string(position), element))
}
func (val HTMLElement) InsertAdjacentHTML(position AdjacentPosition, text string) {
	js.Value(val).Call("insertAdjacentHTML", string(position), text)
}
func (val HTMLElement) InsertAdjacentText(position AdjacentPosition, text string) {
	js.Value(val).Call("insertAdjacentText", string(position), text)
}
func (val HTMLElement) Matches(selectors string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("matches", selectors))
}
func (val HTMLElement) Prepend(nodes ...Node) {
	js.Value(val).Call("prepend", nodesToValues(nodes)...)
}
func (val HTMLElement) QuerySelector(selectors string) Element {
	return wrapElement(js.Value(val).Call("querySelector", selectors))
}
func (val HTMLElement) QuerySelectorAll(selectors string) NodeList {
	return wrapNodeList(js.Value(val).Call("querySelectorAll", selectors))
}
func (val HTMLElement) ReleasePointerCapture(pointerId int) {
	js.Value(val).Call("releasePointerCapture", pointerId)
}
func (val HTMLElement) Remove() {
	js.Value(val).Call("remove")
}
func (val HTMLElement) RemoveAttribute(attrName string) {
	js.Value(val).Call("removeAttribute", attrName)
}
func (val HTMLElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Value(val).Call("removeAttributeNS", namespaceURI, localName)
}
func (val HTMLElement) ReplaceChildren(node ...Node) {
	js.Value(val).Call("replaceChildren", nodesToValues(node)...)
}
func (val HTMLElement) ReplaceWith(node ...Node) {
	js.Value(val).Call("replaceWith", nodesToValues(node)...)
}
func (val HTMLElement) Scroll(options ScrollOptions) {
	js.Value(val).Call("scroll", wrapOptions(options))
}
func (val HTMLElement) ScrollBy(options ScrollOptions) {
	js.Value(val).Call("scrollBy", wrapOptions(options))
}
func (val HTMLElement) ScrollIntoView(options ScrollIntoViewOptions) {
	js.Value(val).Call("scrollIntoView", wrapOptions(options))
}
func (val HTMLElement) ScrollTo(options ScrollOptions) {
	js.Value(val).Call("scrollTo", wrapOptions(options))
}
func (val HTMLElement) SetAttribute(name string, value string) {
	js.Value(val).Call("setAttribute", name, value)
}
func (val HTMLElement) SetAttributeNS(namespace string, name string, value string) {
	js.Value(val).Call("setAttributeNS", namespace, name, value)
}
func (val HTMLElement) SetPointerCapture(pointerId int) {
	js.Value(val).Call("setPointerCapture", pointerId)
}
func (val HTMLElement) ToggleAttribute(name string) {
	js.Value(val).Call("toggleAttribute", name)
}
func (val HTMLElement) Blur() {
	js.Value(val).Call("blur")
}
func (val HTMLElement) Click() {
	js.Value(val).Call("click")
}
func (val HTMLElement) Focus() {
	js.Value(val).Call("focus")
}
func (val HTMLElement) FocusWithOptions(options FocusOptions) {
	js.Value(val).Call("focus", wrapOptions(options))
}

type CSSStyleDeclaration js.Value

func wrapCSSStyleDeclaration(value js.Value) CSSStyleDeclaration {
	return CSSStyleDeclaration(value)
}
func (val CSSStyleDeclaration) CSSText() string {
	return js.Value(val).Get("cssText").String()
}
func (val CSSStyleDeclaration) SetCSSText(cssText string) {
	js.Value(val).Set("cssText", cssText)
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
func (val CSSStyleDeclaration) RemoveProperty(property string) {
	js.Value(val).Call("removeProperty", property)
}
func (val CSSStyleDeclaration) SetProperty(property string, value string) {
	js.Value(val).Call("setProperty", property, value)
}
func (val CSSStyleDeclaration) SetPropertyWithPriority(property string, value string, priority string) {
	js.Value(val).Call("setProperty", property, value, priority)
}

type DocumentFragment js.Value

func wrapDocumentFragment(value js.Value) DocumentFragment {
	return DocumentFragment(value)
}
func (val DocumentFragment) BaseURI() string {
	return js.Value(val).Get("baseURI").String()
}
func (val DocumentFragment) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Get("childNodes"))
}
func (val DocumentFragment) FirstChild() Node {
	return wrapNode(js.Value(val).Get("firstChild"))
}
func (val DocumentFragment) IsConnected() bool {
	return js.Value(val).Get("isConnected").Bool()
}
func (val DocumentFragment) LastChild() Node {
	return wrapNode(js.Value(val).Get("lastChild"))
}
func (val DocumentFragment) NextSibling() Node {
	return wrapNode(js.Value(val).Get("nextSibling"))
}
func (val DocumentFragment) NodeName() string {
	return js.Value(val).Get("nodeName").String()
}
func (val DocumentFragment) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Get("nodeType"))
}
func (val DocumentFragment) NodeValue() Node {
	return wrapNode(js.Value(val).Get("nodeValue"))
}
func (val DocumentFragment) ParentNode() Node {
	return wrapNode(js.Value(val).Get("parentNode"))
}
func (val DocumentFragment) ParentElement() Element {
	return wrapElement(js.Value(val).Get("parentElement"))
}
func (val DocumentFragment) PreviousSibling() Node {
	return wrapNode(js.Value(val).Get("previousSibling"))
}
func (val DocumentFragment) TextContent() string {
	return js.Value(val).Get("textContent").String()
}
func (val DocumentFragment) SetTextContent(textContent string) {
	js.Value(val).Set("textContent", textContent)
}
func (val DocumentFragment) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val DocumentFragment) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val DocumentFragment) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val DocumentFragment) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", convertNodeToValue(aChild)))
}
func (val DocumentFragment) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val DocumentFragment) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", convertNodeToValue(otherNode)))
}
func (val DocumentFragment) Contains(node Node) bool {
	return js.Value(val).Call("contains", convertNodeToValue(node)).Bool()
}
func (val DocumentFragment) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", wrapOptions(options)))
}
func (val DocumentFragment) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val DocumentFragment) InsertBefore(newNode Node, referenceNode Node) Element {
	return wrapElement(js.Value(val).Call("insertBefore", convertNodeToValue(newNode), convertNodeToValue(referenceNode)))
}
func (val DocumentFragment) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val DocumentFragment) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", convertNodeToValue(otherNode)).Bool()
}
func (val DocumentFragment) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", convertNodeToValue(otherNode)).Bool()
}
func (val DocumentFragment) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val DocumentFragment) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val DocumentFragment) Normalize() {
	js.Value(val).Call("normalize")
}
func (val DocumentFragment) RemoveChild(child Node) {
	js.Value(val).Call("removeChild", convertNodeToValue(child))
}
func (val DocumentFragment) ReplaceChild(newChild Node, oldChild Node) {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}

type SVGElement js.Value

func wrapSVGElement(value js.Value) SVGElement {
	return SVGElement(value)
}
func (val SVGElement) BaseURI() string {
	return js.Value(val).Get("baseURI").String()
}
func (val SVGElement) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Get("childNodes"))
}
func (val SVGElement) FirstChild() Node {
	return wrapNode(js.Value(val).Get("firstChild"))
}
func (val SVGElement) IsConnected() bool {
	return js.Value(val).Get("isConnected").Bool()
}
func (val SVGElement) LastChild() Node {
	return wrapNode(js.Value(val).Get("lastChild"))
}
func (val SVGElement) NextSibling() Node {
	return wrapNode(js.Value(val).Get("nextSibling"))
}
func (val SVGElement) NodeName() string {
	return js.Value(val).Get("nodeName").String()
}
func (val SVGElement) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Get("nodeType"))
}
func (val SVGElement) NodeValue() Node {
	return wrapNode(js.Value(val).Get("nodeValue"))
}
func (val SVGElement) ParentNode() Node {
	return wrapNode(js.Value(val).Get("parentNode"))
}
func (val SVGElement) ParentElement() Element {
	return wrapElement(js.Value(val).Get("parentElement"))
}
func (val SVGElement) PreviousSibling() Node {
	return wrapNode(js.Value(val).Get("previousSibling"))
}
func (val SVGElement) TextContent() string {
	return js.Value(val).Get("textContent").String()
}
func (val SVGElement) SetTextContent(textContent string) {
	js.Value(val).Set("textContent", textContent)
}
func (val SVGElement) ChildElementCount() int {
	return js.Value(val).Get("childElementCount").Int()
}
func (val SVGElement) Children() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("children"))
}
func (val SVGElement) ClassList() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Get("classList"))
}
func (val SVGElement) ClassName() string {
	return js.Value(val).Get("className").String()
}
func (val SVGElement) ClientHeight() int {
	return js.Value(val).Get("clientHeight").Int()
}
func (val SVGElement) ClientLeft() int {
	return js.Value(val).Get("clientLeft").Int()
}
func (val SVGElement) ClientTop() int {
	return js.Value(val).Get("clientTop").Int()
}
func (val SVGElement) ClientWidth() int {
	return js.Value(val).Get("clientWidth").Int()
}
func (val SVGElement) FirstElementChild() Element {
	return wrapElement(js.Value(val).Get("firstElementChild"))
}
func (val SVGElement) ID() string {
	return js.Value(val).Get("id").String()
}
func (val SVGElement) InnerHTML() string {
	return js.Value(val).Get("innerHTML").String()
}
func (val SVGElement) SetInnerHTML(innerHTML string) {
	js.Value(val).Set("innerHTML", innerHTML)
}
func (val SVGElement) LocalName() string {
	return js.Value(val).Get("localName").String()
}
func (val SVGElement) NamespaceURI() string {
	return js.Value(val).Get("namespaceURI").String()
}
func (val SVGElement) NextElementSibling() Element {
	return wrapElement(js.Value(val).Get("nextElementSibling"))
}
func (val SVGElement) OuterHTML() string {
	return js.Value(val).Get("outerHTML").String()
}
func (val SVGElement) Part() DOMTokenList {
	return wrapDOMTokenList(js.Value(val).Get("part"))
}
func (val SVGElement) Prefix() string {
	return js.Value(val).Get("prefix").String()
}
func (val SVGElement) PreviousElementSibling() Element {
	return wrapElement(js.Value(val).Get("previousElementSibling"))
}
func (val SVGElement) ScrollHeight() int {
	return js.Value(val).Get("scrollHeight").Int()
}
func (val SVGElement) ScrollLeft() int {
	return js.Value(val).Get("scrollLeft").Int()
}
func (val SVGElement) ScrollTop() int {
	return js.Value(val).Get("scrollTop").Int()
}
func (val SVGElement) ScrollWidth() int {
	return js.Value(val).Get("scrollWidth").Int()
}
func (val SVGElement) Slot() string {
	return js.Value(val).Get("slot").String()
}
func (val SVGElement) TagName() string {
	return js.Value(val).Get("tagName").String()
}
func (val SVGElement) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val SVGElement) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val SVGElement) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val SVGElement) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", convertNodeToValue(aChild)))
}
func (val SVGElement) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val SVGElement) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", convertNodeToValue(otherNode)))
}
func (val SVGElement) Contains(node Node) bool {
	return js.Value(val).Call("contains", convertNodeToValue(node)).Bool()
}
func (val SVGElement) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", wrapOptions(options)))
}
func (val SVGElement) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val SVGElement) InsertBefore(newNode Node, referenceNode Node) Element {
	return wrapElement(js.Value(val).Call("insertBefore", convertNodeToValue(newNode), convertNodeToValue(referenceNode)))
}
func (val SVGElement) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val SVGElement) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", convertNodeToValue(otherNode)).Bool()
}
func (val SVGElement) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", convertNodeToValue(otherNode)).Bool()
}
func (val SVGElement) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val SVGElement) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val SVGElement) Normalize() {
	js.Value(val).Call("normalize")
}
func (val SVGElement) RemoveChild(child Node) {
	js.Value(val).Call("removeChild", convertNodeToValue(child))
}
func (val SVGElement) ReplaceChild(newChild Node, oldChild Node) {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}
func (val SVGElement) After(node ...Node) {
	js.Value(val).Call("after", nodesToValues(node)...)
}
func (val SVGElement) Append(params ...Node) {
	js.Value(val).Call("append", nodesToValues(params)...)
}
func (val SVGElement) Before(params ...Node) {
	js.Value(val).Call("before", nodesToValues(params)...)
}
func (val SVGElement) Closest(selector string) Element {
	return wrapElement(js.Value(val).Call("closest", selector))
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
func (val SVGElement) GetElementByID(id string) Element {
	return wrapElement(js.Value(val).Call("getElementById", id))
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
	return wrapElement(js.Value(val).Call("insertAdjacentElement", string(position), element))
}
func (val SVGElement) InsertAdjacentHTML(position AdjacentPosition, text string) {
	js.Value(val).Call("insertAdjacentHTML", string(position), text)
}
func (val SVGElement) InsertAdjacentText(position AdjacentPosition, text string) {
	js.Value(val).Call("insertAdjacentText", string(position), text)
}
func (val SVGElement) Matches(selectors string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("matches", selectors))
}
func (val SVGElement) Prepend(nodes ...Node) {
	js.Value(val).Call("prepend", nodesToValues(nodes)...)
}
func (val SVGElement) QuerySelector(selectors string) Element {
	return wrapElement(js.Value(val).Call("querySelector", selectors))
}
func (val SVGElement) QuerySelectorAll(selectors string) NodeList {
	return wrapNodeList(js.Value(val).Call("querySelectorAll", selectors))
}
func (val SVGElement) ReleasePointerCapture(pointerId int) {
	js.Value(val).Call("releasePointerCapture", pointerId)
}
func (val SVGElement) Remove() {
	js.Value(val).Call("remove")
}
func (val SVGElement) RemoveAttribute(attrName string) {
	js.Value(val).Call("removeAttribute", attrName)
}
func (val SVGElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Value(val).Call("removeAttributeNS", namespaceURI, localName)
}
func (val SVGElement) ReplaceChildren(node ...Node) {
	js.Value(val).Call("replaceChildren", nodesToValues(node)...)
}
func (val SVGElement) ReplaceWith(node ...Node) {
	js.Value(val).Call("replaceWith", nodesToValues(node)...)
}
func (val SVGElement) Scroll(options ScrollOptions) {
	js.Value(val).Call("scroll", wrapOptions(options))
}
func (val SVGElement) ScrollBy(options ScrollOptions) {
	js.Value(val).Call("scrollBy", wrapOptions(options))
}
func (val SVGElement) ScrollIntoView(options ScrollIntoViewOptions) {
	js.Value(val).Call("scrollIntoView", wrapOptions(options))
}
func (val SVGElement) ScrollTo(options ScrollOptions) {
	js.Value(val).Call("scrollTo", wrapOptions(options))
}
func (val SVGElement) SetAttribute(name string, value string) {
	js.Value(val).Call("setAttribute", name, value)
}
func (val SVGElement) SetAttributeNS(namespace string, name string, value string) {
	js.Value(val).Call("setAttributeNS", namespace, name, value)
}
func (val SVGElement) SetPointerCapture(pointerId int) {
	js.Value(val).Call("setPointerCapture", pointerId)
}
func (val SVGElement) ToggleAttribute(name string) {
	js.Value(val).Call("toggleAttribute", name)
}

type Document js.Value

func wrapDocument(value js.Value) Document {
	return Document(value)
}
func (val Document) BaseURI() string {
	return js.Value(val).Get("baseURI").String()
}
func (val Document) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Get("childNodes"))
}
func (val Document) FirstChild() Node {
	return wrapNode(js.Value(val).Get("firstChild"))
}
func (val Document) IsConnected() bool {
	return js.Value(val).Get("isConnected").Bool()
}
func (val Document) LastChild() Node {
	return wrapNode(js.Value(val).Get("lastChild"))
}
func (val Document) NextSibling() Node {
	return wrapNode(js.Value(val).Get("nextSibling"))
}
func (val Document) NodeName() string {
	return js.Value(val).Get("nodeName").String()
}
func (val Document) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Get("nodeType"))
}
func (val Document) NodeValue() Node {
	return wrapNode(js.Value(val).Get("nodeValue"))
}
func (val Document) ParentNode() Node {
	return wrapNode(js.Value(val).Get("parentNode"))
}
func (val Document) ParentElement() Element {
	return wrapElement(js.Value(val).Get("parentElement"))
}
func (val Document) PreviousSibling() Node {
	return wrapNode(js.Value(val).Get("previousSibling"))
}
func (val Document) TextContent() string {
	return js.Value(val).Get("textContent").String()
}
func (val Document) SetTextContent(textContent string) {
	js.Value(val).Set("textContent", textContent)
}
func (val Document) ActiveElement() Element {
	return wrapElement(js.Value(val).Get("activeElement"))
}
func (val Document) Body() Element {
	return wrapElement(js.Value(val).Get("body"))
}
func (val Document) CharacterSet() string {
	return js.Value(val).Get("characterSet").String()
}
func (val Document) ChildElementCount() int {
	return js.Value(val).Get("childElementCount").Int()
}
func (val Document) Children() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("children"))
}
func (val Document) ContentType() string {
	return js.Value(val).Get("contentType").String()
}
func (val Document) CurrentScript() Element {
	return wrapElement(js.Value(val).Get("currentScript"))
}
func (val Document) DocType() string {
	return js.Value(val).Get("doctype").String()
}
func (val Document) DocumentElement() Element {
	return wrapElement(js.Value(val).Get("documentElement"))
}
func (val Document) DocumentURI() string {
	return js.Value(val).Get("documentURI").String()
}
func (val Document) Forms() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("forms"))
}
func (val Document) Head() Element {
	return wrapElement(js.Value(val).Get("head"))
}
func (val Document) Hidden() bool {
	return js.Value(val).Get("hidden").Bool()
}
func (val Document) Images() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("images"))
}
func (val Document) Links() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("links"))
}
func (val Document) Scripts() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("scripts"))
}
func (val Document) StyleSheets() HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Get("styleSheets"))
}
func (val Document) Cookie() string {
	return js.Value(val).Get("cookie").String()
}
func (val Document) Dir() string {
	return js.Value(val).Get("dir").String()
}
func (val Document) SetDir(dir string) {
	js.Value(val).Set("dir", dir)
}
func (val Document) Domain() string {
	return js.Value(val).Get("domain").String()
}
func (val Document) SetDomain(domain string) {
	js.Value(val).Set("domain", domain)
}
func (val Document) Location() string {
	return js.Value(val).Get("location").String()
}
func (val Document) ReadyState() string {
	return js.Value(val).Get("readyState").String()
}
func (val Document) Referrer() string {
	return js.Value(val).Get("referrer").String()
}
func (val Document) Title() string {
	return js.Value(val).Get("title").String()
}
func (val Document) SetTitle(title string) {
	js.Value(val).Set("title", title)
}
func (val Document) URL() string {
	return js.Value(val).Get("URL").String()
}
func (val Document) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val Document) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val Document) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val Document) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", convertNodeToValue(aChild)))
}
func (val Document) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val Document) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", convertNodeToValue(otherNode)))
}
func (val Document) Contains(node Node) bool {
	return js.Value(val).Call("contains", convertNodeToValue(node)).Bool()
}
func (val Document) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", wrapOptions(options)))
}
func (val Document) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val Document) InsertBefore(newNode Node, referenceNode Node) Element {
	return wrapElement(js.Value(val).Call("insertBefore", convertNodeToValue(newNode), convertNodeToValue(referenceNode)))
}
func (val Document) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val Document) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", convertNodeToValue(otherNode)).Bool()
}
func (val Document) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", convertNodeToValue(otherNode)).Bool()
}
func (val Document) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val Document) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val Document) Normalize() {
	js.Value(val).Call("normalize")
}
func (val Document) RemoveChild(child Node) {
	js.Value(val).Call("removeChild", convertNodeToValue(child))
}
func (val Document) ReplaceChild(newChild Node, oldChild Node) {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}
func (val Document) Append(params ...Node) {
	js.Value(val).Call("append", nodesToValues(params)...)
}
func (val Document) AdoptNode(externalNode Node) {
	js.Value(val).Call("adoptNode", convertNodeToValue(externalNode))
}
func (val Document) CreateElement(tagName string) Element {
	return wrapElement(js.Value(val).Call("createElement", tagName))
}
func (val Document) CreateElementWithOptions(tagName string, options CreateElementOptions) Element {
	return wrapElement(js.Value(val).Call("createElement", tagName, wrapOptions(options)))
}
func (val Document) CreateElementNS(namespaceURI string, qualifiedName string, options CreateElementOptions) Element {
	return wrapElement(js.Value(val).Call("createElementNS", namespaceURI, qualifiedName, wrapOptions(options)))
}
func (val Document) CreateTextNode(text string) Text {
	return wrapText(js.Value(val).Call("createTextNode", text))
}
func (val Document) GetElementByID(id string) Element {
	return wrapElement(js.Value(val).Call("getElementById", id))
}
func (val Document) GetElementsByClassName(id string) HTMLCollection {
	return wrapHTMLCollection(js.Value(val).Call("getElementsByClassName", id))
}
func (val Document) Prepend(nodes ...Node) {
	js.Value(val).Call("prepend", nodesToValues(nodes)...)
}
func (val Document) QuerySelector(selectors string) Element {
	return wrapElement(js.Value(val).Call("querySelector", selectors))
}
func (val Document) QuerySelectorAll(selectors string) NodeList {
	return wrapNodeList(js.Value(val).Call("querySelectorAll", selectors))
}

type Window js.Value

func wrapWindow(value js.Value) Window {
	return Window(value)
}
func (val Window) Document() Document {
	return wrapDocument(js.Value(val).Get("document"))
}
func (val Window) FrameElement() Element {
	return wrapElement(js.Value(val).Get("frameElement"))
}
func (val Window) InnerHeight() int {
	return js.Value(val).Get("innerHeight").Int()
}
func (val Window) InnerWidth() int {
	return js.Value(val).Get("innerWidth").Int()
}
func (val Window) IsSecureContext() bool {
	return js.Value(val).Get("isSecureContext").Bool()
}
func (val Window) Name() string {
	return js.Value(val).Get("name").String()
}
func (val Window) Parent() Window {
	return wrapWindow(js.Value(val).Get("parent"))
}
func (val Window) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val Window) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val Window) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val Window) Alert(message string) {
	js.Value(val).Call("alert", message)
}
func (val Window) Blur() {
	js.Value(val).Call("blur")
}
func (val Window) Close() {
	js.Value(val).Call("close")
}
func (val Window) Confirm(message string) bool {
	return js.Value(val).Call("confirm", message).Bool()
}
func (val Window) Focus() {
	js.Value(val).Call("focus")
}
func (val Window) GetComputedStyle(el Element) {
	js.Value(val).Call("getComputedStyle", el)
}
func (val Window) Open(url string, target string, windowFeatures string) {
	js.Value(val).Call("open", url, target, windowFeatures)
}
func (val Window) PostMessage(message any, postMessage string) {
	js.Value(val).Call("postMessage", message, postMessage)
}
func (val Window) Prompt(message string, defaultValue string) string {
	return js.Value(val).Call("prompt", message, defaultValue).String()
}

type CharacterData interface {
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
	SetTextContent(textContent string)
	Data() string
	Length() int
	NextElementSibling() Element
	PreviousElementSibling() Element
	AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool)
	DispatchEvent(event Event) bool
	AppendChild(aChild Node) Node
	CloneNode(deep bool) Node
	CompareDocumentPosition(otherNode Node) DocumentPosition
	Contains(node Node) bool
	GetRootNode(options GetRootNodeOptions) Node
	HasChildNodes() bool
	InsertBefore(newNode Node, referenceNode Node) Element
	IsDefaultNamespace(namespaceURL string) bool
	IsEqualNode(otherNode Node) bool
	IsOtherNode(otherNode Node) bool
	LookupPrefix(prefix string) string
	LookupNamespaceURI(prefix string) string
	Normalize()
	RemoveChild(child Node)
	ReplaceChild(newChild Node, oldChild Node)
}
type Text js.Value

func wrapText(value js.Value) Text {
	return Text(value)
}
func (val Text) BaseURI() string {
	return js.Value(val).Get("baseURI").String()
}
func (val Text) ChildNodes() NodeList {
	return wrapNodeList(js.Value(val).Get("childNodes"))
}
func (val Text) FirstChild() Node {
	return wrapNode(js.Value(val).Get("firstChild"))
}
func (val Text) IsConnected() bool {
	return js.Value(val).Get("isConnected").Bool()
}
func (val Text) LastChild() Node {
	return wrapNode(js.Value(val).Get("lastChild"))
}
func (val Text) NextSibling() Node {
	return wrapNode(js.Value(val).Get("nextSibling"))
}
func (val Text) NodeName() string {
	return js.Value(val).Get("nodeName").String()
}
func (val Text) NodeType() NodeType {
	return wrapNodeType(js.Value(val).Get("nodeType"))
}
func (val Text) NodeValue() Node {
	return wrapNode(js.Value(val).Get("nodeValue"))
}
func (val Text) ParentNode() Node {
	return wrapNode(js.Value(val).Get("parentNode"))
}
func (val Text) ParentElement() Element {
	return wrapElement(js.Value(val).Get("parentElement"))
}
func (val Text) PreviousSibling() Node {
	return wrapNode(js.Value(val).Get("previousSibling"))
}
func (val Text) TextContent() string {
	return js.Value(val).Get("textContent").String()
}
func (val Text) SetTextContent(textContent string) {
	js.Value(val).Set("textContent", textContent)
}
func (val Text) Data() string {
	return js.Value(val).Get("data").String()
}
func (val Text) Length() int {
	return js.Value(val).Length()
}
func (val Text) NextElementSibling() Element {
	return wrapElement(js.Value(val).Get("nextElementSibling"))
}
func (val Text) PreviousElementSibling() Element {
	return wrapElement(js.Value(val).Get("previousElementSibling"))
}
func (val Text) WholeText() string {
	return js.Value(val).Get("wholeText").String()
}
func (val Text) AssignedSlot() Element {
	return wrapElement(js.Value(val).Get("assignedSlot"))
}
func (val Text) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val Text) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val Text) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val Text) AppendChild(aChild Node) Node {
	return wrapNode(js.Value(val).Call("appendChild", convertNodeToValue(aChild)))
}
func (val Text) CloneNode(deep bool) Node {
	return wrapNode(js.Value(val).Call("cloneNode", deep))
}
func (val Text) CompareDocumentPosition(otherNode Node) DocumentPosition {
	return wrapDocumentPosition(js.Value(val).Call("compareDocumentPosition", convertNodeToValue(otherNode)))
}
func (val Text) Contains(node Node) bool {
	return js.Value(val).Call("contains", convertNodeToValue(node)).Bool()
}
func (val Text) GetRootNode(options GetRootNodeOptions) Node {
	return wrapNode(js.Value(val).Call("getRootNode", wrapOptions(options)))
}
func (val Text) HasChildNodes() bool {
	return js.Value(val).Call("hasChildNodes").Bool()
}
func (val Text) InsertBefore(newNode Node, referenceNode Node) Element {
	return wrapElement(js.Value(val).Call("insertBefore", convertNodeToValue(newNode), convertNodeToValue(referenceNode)))
}
func (val Text) IsDefaultNamespace(namespaceURL string) bool {
	return js.Value(val).Call("isDefaultNamespace", namespaceURL).Bool()
}
func (val Text) IsEqualNode(otherNode Node) bool {
	return js.Value(val).Call("isEqualNode", convertNodeToValue(otherNode)).Bool()
}
func (val Text) IsOtherNode(otherNode Node) bool {
	return js.Value(val).Call("isOtherNode", convertNodeToValue(otherNode)).Bool()
}
func (val Text) LookupPrefix(prefix string) string {
	return js.Value(val).Call("lookupPrefix", prefix).String()
}
func (val Text) LookupNamespaceURI(prefix string) string {
	return js.Value(val).Call("lookupNamespaceURI", prefix).String()
}
func (val Text) Normalize() {
	js.Value(val).Call("normalize")
}
func (val Text) RemoveChild(child Node) {
	js.Value(val).Call("removeChild", convertNodeToValue(child))
}
func (val Text) ReplaceChild(newChild Node, oldChild Node) {
	js.Value(val).Call("replaceChild", newChild, oldChild)
}

type EventSource js.Value

func wrapEventSource(value js.Value) EventSource {
	return EventSource(value)
}
func (val EventSource) ReadyState() int {
	return js.Value(val).Get("readyState").Int()
}
func (val EventSource) URL() string {
	return js.Value(val).Get("url").String()
}
func (val EventSource) WithCredentials() bool {
	return js.Value(val).Get("withCredentials").Bool()
}
func (val EventSource) AddEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("addEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val EventSource) RemoveEventListener(eventType string, listener js.Func, options AddEventListenerOptions, useCapture bool) {
	js.Value(val).Call("removeEventListener", eventType, listener, wrapOptions(options), useCapture)
}
func (val EventSource) DispatchEvent(event Event) bool {
	return js.Value(val).Call("dispatchEvent", event).Bool()
}
func (val EventSource) Close() {
	js.Value(val).Call("close")
}

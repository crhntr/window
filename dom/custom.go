package dom

import (
	"fmt"
	"reflect"
	"syscall/js"
	"time"
)

//go:generate go run ../internal/generate ../dom.yml

type AdjacentPosition string

// noinspection SpellCheckingInspection
const (
	PositionBeforeBegin AdjacentPosition = "beforebegin"
	PositionAfterBegin  AdjacentPosition = "afterbegin"
	PositionBeforeEnd   AdjacentPosition = "beforeend"
	PositionAfterEnd    AdjacentPosition = "afterend"
)

type AddEventListenerOptions map[string]any

func (options AddEventListenerOptions) SetCapture(capture bool) AddEventListenerOptions {
	options["capture"] = capture
	return options
}

func (options AddEventListenerOptions) SetPassive(passive bool) AddEventListenerOptions {
	options["passive"] = passive
	return options
}

func (options AddEventListenerOptions) SetSignal(signal js.Value) AddEventListenerOptions {
	options["signal"] = signal
	return options
}

type GetRootNodeOptions map[string]any

func (options AddEventListenerOptions) SetComposed(composed bool) AddEventListenerOptions {
	options["composed"] = composed
	return options
}

type ScrollOptions map[string]any

func (options ScrollOptions) SetTop(value int) ScrollOptions {
	options["top"] = value
	return options
}

func (options ScrollOptions) SetLeft(value int) ScrollOptions {
	options["top"] = value
	return options
}

func (options ScrollOptions) SetBehavior(behavior ScrollBehavior) ScrollOptions {
	options["behavior"] = string(behavior)
	return options
}

type AttachShadowOptions map[string]any

func (options AttachShadowOptions) SetMode(s string) AttachShadowOptions {
	options["mode"] = s
	return options
}

func (options AttachShadowOptions) SetDelegateFocus(focus bool) AttachShadowOptions {
	options["delegatesFocus"] = focus
	return options
}

type NodeType uint8

const (
	NodeTypeElement NodeType = iota + 1
	NodeTypeAttribute
	NodeTypeText
	NodeTypeCdataSection
	NodeTypeProcessingInstruction
	NodeTypeComment
	NodeTypeDocument
	NodeTypeDocumentType
	NodeTypeDocumentFragment
)

func wrapNodeType(v js.Value) NodeType {
	return NodeType(v.Int())
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

func wrapDocumentPosition(v js.Value) DocumentPosition { return DocumentPosition(v.Int()) }

type ScrollAlignment string

const (
	ScrollAlignmentStart   ScrollAlignment = "start"
	ScrollAlignmentCenter  ScrollAlignment = "center"
	ScrollAlignmentEnd     ScrollAlignment = "end"
	ScrollAlignmentNearest ScrollAlignment = "nearest"
)

type ScrollBehavior string

const (
	ScrollBehaviorSmooth  = "smooth"
	ScrollBehaviorInstant = "instant"
)

type ScrollIntoViewOptions map[string]any

func (options FocusOptions) SetBlock(value ScrollAlignment) FocusOptions {
	options["block"] = value
	return options
}

func (options FocusOptions) SetInline(value ScrollAlignment) FocusOptions {
	options["inline"] = value
	return options
}

func (options FocusOptions) SetBehavior(value bool) FocusOptions {
	options["behavior"] = value
	return options
}

type CreateElementOptions map[string]any

func (options CreateElementOptions) SetIs(value string) CreateElementOptions {
	options["is"] = value
	return options
}

type EventPhase int

const (
	EventPhaseNone EventPhase = iota
	EventPhaseCapturing
	EventPhaseAtTarget
	EventPhaseBubbling
)

func wrapEventPhase(v js.Value) EventPhase {
	return EventPhase(v.Int())
}

type FocusOptions map[string]any

func (options FocusOptions) SetPreventScroll(value bool) FocusOptions {
	options["preventScroll"] = value
	return options
}

func (options FocusOptions) SetFocusVisible(value bool) FocusOptions {
	options["focusVisible"] = value
	return options
}

type DOMStringMap js.Value

func wrapDOMStringMap(value js.Value) DOMStringMap { return DOMStringMap(value) }

func (m DOMStringMap) Get(name string) string { return js.Value(m).Get(name).String() }
func (m DOMStringMap) Set(name, value string) { js.Value(m).Set(name, value) }
func (m DOMStringMap) Delete(name string)     { js.Value(m).Delete(name) }

// Node implementations
var (
	textNodePrototype         = js.Global().Get(reflect.TypeOf(Text{}).Name())
	htmlElementPrototype      = js.Global().Get(reflect.TypeOf(HTMLElement{}).Name())
	svgElementPrototype       = js.Global().Get(reflect.TypeOf(SVGElement{}).Name())
	documentFragmentPrototype = js.Global().Get(reflect.TypeOf(DocumentFragment{}).Name())
)

func wrapNode(value js.Value) Node {
	switch {
	case value.InstanceOf(htmlElementPrototype):
		return wrapHTMLElement(value)
	case value.InstanceOf(textNodePrototype):
		return wrapText(value)
	case value.InstanceOf(svgElementPrototype):
		return wrapSVGElement(value)
	case value.InstanceOf(documentFragmentPrototype):
		return wrapDocumentFragment(value)
	default:
		message := fmt.Sprintf("failed to wrap Node %s", js.Global().Get("Object").Call("getPrototypeOf", value))
		js.Global().Get("console").Call("log", value)
		panic(message)
		return nil
	}
}

func wrapElement(value js.Value) Element {
	switch {
	case value.InstanceOf(htmlElementPrototype):
		return wrapHTMLElement(value)
	case value.InstanceOf(svgElementPrototype):
		return wrapSVGElement(value)
	default:
		panic("failed to wrap Element")
	}
}

func IFrameContentWindow(el Element) Window {
	switch e := el.(type) {
	case HTMLElement:
		if e.TagName() != "IFRAME" {
			panic("not an IFRAME got:" + e.TagName())
		}
		return Window(js.Value(e).Get("contentWindow"))
	default:
		panic("not an IFRAME")
	}
}

func GetValue(el Element) string {
	switch e := el.(type) {
	case HTMLElement:
		return js.Value(e).Get("value").String()
	}
	panic(fmt.Sprintf("not an Element: %T", el))
}

func SetValue(el Element, value string) {
	switch e := el.(type) {
	case HTMLElement:
		js.Value(e).Set("value", value)
		return
	}
	panic(fmt.Sprintf("not an Element: %T", el))
}

//// Event implementations
//var (
//	eventPrototype                        = js.Global().Get(reflect.TypeOf(Event{}).Name())
//	uIEventPrototype                      = js.Global().Get(reflect.TypeOf(UIEvent{}).Name())
//	focusEventPrototype                   = js.Global().Get(reflect.TypeOf(FocusEvent{}).Name())
//	inputEventPrototype                   = js.Global().Get(reflect.TypeOf(InputEvent{}).Name())
//	securityPolicyViolationEventPrototype = js.Global().Get(reflect.TypeOf(SecurityPolicyViolationEvent{}).Name())
//)

type EventValue interface {
	UIEvent |
		FocusEvent |
		InputEvent |
		SecurityPolicyViolationEvent |
		MessageEvent |
		GenericEvent
}

func millisecondsSinceEpocToTime(val js.Value) time.Time {
	return time.Unix(0, int64(val.Int()))
}

func makeStaticRangeSlice(val js.Value) []StaticRange {
	return makeValueSlice(val, func(val js.Value) StaticRange { return StaticRange(val) })
}

func makeValueSlice[V any](val js.Value, fn func(val js.Value) V) []V {
	result := make([]V, val.Length())
	for i := 0; i < val.Length(); i++ {
		result[i] = fn(val.Index(i))
	}
	return result
}

func convertNodeToValue(input Node) js.Value {
	switch n := input.(type) {
	case HTMLElement:
		return js.Value(n)
	case SVGElement:
		return js.Value(n)
	case DocumentFragment:
		return js.Value(n)
	case Text:
		return js.Value(n)
	default:
		panic("failed to convert Node to js.Value")
	}
}

func nodesToValues(nodes []Node) []any {
	values := make([]any, 0, len(nodes))
	for _, n := range nodes {
		values = append(values, convertNodeToValue(n))
	}
	return values
}

func stringsToAny(input []string) []any {
	values := make([]any, 0, len(input))
	for _, s := range input {
		values = append(values, s)
	}
	return values
}

func wrapOptions[O interface {
	AddEventListenerOptions |
		GetRootNodeOptions |
		ScrollOptions |
		AttachShadowOptions |
		ScrollIntoViewOptions |
		CreateElementOptions |
		FocusOptions
}](o O) js.Value {
	return js.ValueOf((map[string]any)(o))
}
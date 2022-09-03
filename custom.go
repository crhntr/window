package window

import (
	"reflect"
	"syscall/js"
)

type AdjacentPosition string

// noinspection SpellCheckingInspection
const (
	BeforeBegin AdjacentPosition = "beforebegin"
	AfterBegin  AdjacentPosition = "afterbegin"
	BeforeEnd   AdjacentPosition = "beforeend"
	AfterEnd    AdjacentPosition = "afterend"
)

type AddEventListenerOptions struct {
	Capture bool
	Passive bool
	Signal  js.Value
}

func (options AddEventListenerOptions) value() js.Value {
	return js.ValueOf(map[string]any{
		"capture": options.Capture,
		"passive": options.Passive,
		"signal":  options.Signal,
	})
}

type GetRootNodeOptions struct {
	Composed bool
}

func (options GetRootNodeOptions) value() js.Value {
	return js.ValueOf(map[string]any{
		"composed": options.Composed,
	})
}

type ScrollOptions struct {
	Top, Left int
	Behavior  ScrollBehavior
}

func (options ScrollOptions) value() js.Value {
	return js.ValueOf(map[string]any{
		"top":      options.Top,
		"left":     options.Left,
		"behavior": string(options.Behavior),
	})
}

type AttachShadowOptions struct {
	Mode           string
	DelegatesFocus bool
}

func (options AttachShadowOptions) value() js.Value {
	return js.ValueOf(map[string]any{
		"mode":           options.Mode,
		"delegatesFocus": options.DelegatesFocus,
	})
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

type ScrollIntoViewOptions struct {
	Block, Inline ScrollAlignment
	Behavior      ScrollBehavior
}

func (options ScrollIntoViewOptions) value() js.Value {
	return js.ValueOf(map[string]any{
		"block":    options.Block,
		"inline":   options.Inline,
		"behavior": options.Behavior,
	})
}

type EventPhase int

const (
	EventPhaseNone EventPhase = iota
	EventPhaseCapturing
	EventPhaseAtTarget
	EventPhaseBubbling
)

func warpEventPhase(v js.Value) EventPhase {
	return EventPhase(v.Int())
}

type FocusOptions struct {
	PreventScroll bool
	FocusVisible  bool
}

func (options FocusOptions) value() js.Value {
	return js.ValueOf(map[string]any{
		"preventScroll": options.PreventScroll,
		"focusVisible":  options.FocusVisible,
	})
}

type DOMStringMap js.Value

func wrapDOMStringMap(value js.Value) DOMStringMap { return DOMStringMap(value) }

func (m DOMStringMap) Get(name string) string { return js.Value(m).Get(name).String() }
func (m DOMStringMap) Set(name, value string) { js.Value(m).Set(name, value) }
func (m DOMStringMap) Delete(name string)     { js.Value(m).Delete(name) }

type nodePrototypes struct {
	HTMLElement
	SVGElement
	DocumentFragment
}

var (
	globalHtmlElementPrototype      = js.Global().Get(reflect.TypeOf(HTMLElement{}).Name())
	globalSvgElement                = js.Global().Get(reflect.TypeOf(SVGElement{}).Name())
	globalDocumentFragmentPrototype = js.Global().Get(reflect.TypeOf(DocumentFragment{}).Name())
)

func wrapNode(value js.Value) Node {
	switch {
	case value.InstanceOf(globalHtmlElementPrototype):
		return wrapHTMLElement(value)
	case value.InstanceOf(globalSvgElement):
		return wrapSVGElement(value)
	case value.InstanceOf(globalDocumentFragmentPrototype):
		return wrapDocumentFragment(value)
	default:
		js.Global().Call("console").Call("log", "failed to wrap Node", value)
		return nil
	}
}

func wrapElement(value js.Value) Element {
	switch {
	case value.InstanceOf(globalHtmlElementPrototype):
		return wrapHTMLElement(value)
	case value.InstanceOf(globalSvgElement):
		return wrapSVGElement(value)
	default:
		js.Global().Call("console").Call("log", "failed to wrap Element", value)
		return nil
	}
	return nil
}

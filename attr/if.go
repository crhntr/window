package attr

import (
	"errors"
	"reflect"
	"strings"

	"github.com/crhntr/window/dom"
)

func handleConditionals(node dom.Element, data reflect.Value) error {
	foundTrue, err := handleIfDirective(node, data)
	if err != nil {
		return err
	}
	next, ok := nextElement(node)
	if !ok || !next.HasAttribute(ElseDirective) {
		return nil
	}
	return handleElseDirective(next, foundTrue)
}

func handleIfDirective(node dom.Element, data reflect.Value) (bool, error) {
	ifFieldName := strings.TrimSpace(node.Attribute(IfDirective))

	ifResultIsTrue, err := resolveBooleanForIfDirective(ifFieldName, data)
	if err != nil {
		return false, err
	}

	switch {
	case ifResultIsTrue && node.HasAttribute(HiddenDirective):
		tmpl := node.Attribute(TemplateDirective)
		node.RemoveAttribute(HiddenDirective)
		node.RemoveAttribute(TemplateDirective)
		node.SetInnerHTML(tmpl)
	case !ifResultIsTrue && !node.HasAttribute(HiddenDirective):
		node.SetAttribute(TemplateDirective, node.InnerHTML())
		node.SetAttribute(HiddenDirective, "")
		node.ReplaceChildren()
	}

	return ifResultIsTrue, nil
}

func handleElseDirective(node dom.Element, previousTrue bool) error {
	switch {
	case !previousTrue && node.HasAttribute(HiddenDirective):
		tmpl := node.Attribute(TemplateDirective)
		node.RemoveAttribute(HiddenDirective)
		node.RemoveAttribute(TemplateDirective)
		node.SetInnerHTML(tmpl)
	case previousTrue && !node.HasAttribute(HiddenDirective):
		node.SetAttribute(TemplateDirective, node.InnerHTML())
		node.SetAttribute(HiddenDirective, "")
		node.ReplaceChildren()
	}
	return nil
}

func resolveBooleanForIfDirective(ifExpStr string, data reflect.Value) (bool, error) {
	fieldValue := data.FieldByName(ifExpStr)
	if fieldValue.IsValid() {
		if fieldValue.Kind() != reflect.Bool {
			return false, errors.New("attr: field must have type bool")
		}
		return fieldValue.Bool(), nil
	}
	method := data.MethodByName(ifExpStr)
	if !method.IsValid() {
		return false, errors.New("attr: field or method not found")
	}
	if method.Type().NumIn() != 0 {
		return false, errors.New("attr: boolean method must not receive any parameter")
	}
	if method.Type().NumOut() != 1 || method.Type().Out(0).Kind() != reflect.Bool {
		return false, errors.New("attr: boolean method only return a single boolean")
	}
	return method.Interface().(func() bool)(), nil
}

func nextElement(node dom.Element) (dom.Element, bool) {
	for {
		next := node.NextSibling()
		if next == nil {
			return nil, false
		}
		el, ok := next.(dom.Element)
		if !ok {
			continue
		}
		return el, true
	}
}

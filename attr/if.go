package attr

import (
	"errors"
	"reflect"
	"strings"

	"github.com/crhntr/window/dom"
)

func handleIfDirective(node dom.Element, data reflect.Value) (bool, error) {
	ifFieldName := strings.TrimSpace(node.Attribute(IfDirective))

	ifResultIsTrue, err := resolveBooleanForIfDirective(ifFieldName, data)
	if err != nil {
		return false, err
	}

	if ifResultIsTrue {
		node.RemoveAttribute(IfResultDirective)
		if node.HasAttribute(TemplateDirective) {
			tmpl := node.Attribute(TemplateDirective)
			node.RemoveAttribute(TemplateDirective)
			node.SetInnerHTML(tmpl)
		}
		next := node.NextSibling()
		if next != nil {
			nextEl, ok := next.(dom.Element)
			if ok && nextEl.HasAttribute(ElseDirective) {
				if !nextEl.HasAttribute(TemplateDirective) {
					nextEl.SetAttribute(TemplateDirective, nextEl.InnerHTML())
				}
				nextEl.ReplaceChildren()
			}
		}
	} else {
		if node.Attribute(IfResultDirective) != "false" {
			if !node.HasAttribute(TemplateDirective) {
				node.SetAttribute(TemplateDirective, node.InnerHTML())
			}
			node.SetAttribute(IfResultDirective, "false")
			node.ReplaceChildren()
		}
		next := node.NextSibling()
		if next != nil {
			nextEl, ok := next.(dom.Element)
			if ok && nextEl.HasAttribute(TemplateDirective) {
				tmpl := nextEl.Attribute(TemplateDirective)
				nextEl.RemoveAttribute(TemplateDirective)
				nextEl.SetInnerHTML(tmpl)
			}
		}
	}

	return ifResultIsTrue, nil
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

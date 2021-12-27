package attr

import (
	"errors"
	"reflect"

	"github.com/crhntr/window/dom"
)

const (
	TemplateDirective = "data-attr-template"

	IfDirective       = "data-attr-if"
	IfResultDirective = "data-attr-if-result"
	ElseDirective     = "data-attr-else"

	RangeDirective         = "data-attr-range"
	RangeIndexDirective    = "data-attr-range-index"
	RangeKeyDirective      = "data-attr-range-key"
	RangeKeyValueDirective = "data-attr-range-key-value"
)

func Hydrate(node dom.Element, data interface{}) error {
	return hydrate(node, reflect.ValueOf(data))
}

func hydrate(node dom.Element, data reflect.Value) error {
	if data.Kind() != reflect.Struct {
		return errors.New("data must be a struct")
	}

	if node.HasAttribute(IfDirective) {
		_, err := handleIfDirective(node, data)
		if err != nil {
			return err
		}
	}

	if node.HasAttribute(RangeDirective) {
		err := handleRangeDirective(node, data)
		if err != nil {
			return err
		}
	}

	children := node.Children()
	for i := 0; i < children.Length(); i++ {
		err := hydrate(children.Item(i), data)
		if err != nil {
			continue
		}
	}

	return nil
}

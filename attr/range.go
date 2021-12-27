package attr

import (
	"errors"
	"reflect"

	"github.com/crhntr/window/dom"
)

func handleRangeDirective(node dom.Element, data reflect.Value) error {
	rangeFieldName := node.Attribute(RangeKeyDirective)

	field := data.FieldByName(rangeFieldName)
	if !field.IsValid() {
		return errors.New("attr: range field not found")
	}
	if data.Kind() != reflect.Slice {
		return errors.New("attr: only range over slice kind is allowed")
	}

	var template string
	if node.HasAttribute(TemplateDirective) {
		template = node.Attribute(TemplateDirective)
	} else {
		template = node.InnerHTML()
	}
	node.ReplaceChildren()
	if len(template) == 0 {
		return nil
	}

	for index := 0; index < field.Len(); index++ {

	}

	return nil
}

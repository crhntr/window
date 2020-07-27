// +build js,wasm

package window

import (
	"html/template"
)

func LoadTemplates(tmp *template.Template, selector string) (*template.Template, error) {
	if tmp == nil {
		tmp = template.New("")
	}

	if selector == "" {
		selector = "script[type=text/go-template]"
	}

	for _, el := range QuerySelectorAll(selector) {
		var err error
		tmp, err = tmp.New(el.Attribute("id")).Parse(el.InnerHTML())
		if err != nil {
			return tmp, err
		}
	}

	return tmp, nil
}

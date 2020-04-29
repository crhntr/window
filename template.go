// +build js,wasm

package dom

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

	list := QuerySelectorAll(selector)
	for _, el := range list {
		var err error
		tmp, err = tmp.New(el.Attribute("id")).Parse(el.InnerHTML())
		if err != nil {
			return tmp, err
		}
	}

	return tmp, nil
}
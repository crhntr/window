// +build js,wasm

package window

import (
	"html/template"
	"sync"
)

var (
	once      sync.Once
	templates *template.Template
)

func LoadTemplates(tmp *template.Template, selector string) error {
	var err error

	once.Do(func() {
		if tmp == nil {
			tmp = template.New("")
		}

		if selector == "" {
			selector = `script[type="text/go-template"]`
		}

		for _, el := range Document.QuerySelectorAll(selector).ElementSlice() {
			templateBody := el.InnerHTML()
			templates, err = tmp.New(el.Attribute("id")).Parse(templateBody)
			if err != nil {
				Console.Log("failed parsing template", err.Error(), templateBody)
				break
			}
		}
	})

	return err
}

func Templates() []*template.Template {
	if templates == nil {
		return nil
	}
	return templates.Templates()
}

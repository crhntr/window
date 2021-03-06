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

func LoadTemplates(t *template.Template, query string) (*template.Template, error) {
	if t == nil {
		t = template.New("")
	}
	if query == "" {
		query = `script[type="text/go-template"]`
	}

	for _, el := range Document.QuerySelectorAll(query).ElementSlice() {
		templateBody := el.InnerHTML()
		var err error
		templates, err = t.New(el.Attribute("id")).Parse(templateBody)
		if err != nil {
			Console.Log("failed parsing template", err.Error(), templateBody)
			break
		}
	}

	return templates, nil
}

func SetTemplates(t *template.Template) {
	once.Do(func() {
		templates = t
	})
}

func Templates() []*template.Template {
	if templates == nil {
		return nil
	}
	return templates.Templates()
}

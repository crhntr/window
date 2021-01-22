// +build js,wasm

package window

import (
	"html/template"
	"sync"
)

var (
	once sync.Once
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

		for _, el := range Document.QuerySelectorAll(selector) {
			templates, err = tmp.New(el.Attribute("id")).Parse(el.InnerHTML())
			if err != nil {
				break
			}
		}
	})


	return err
}

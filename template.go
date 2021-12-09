//go:build js && wasm && !tinygo
// +build js,wasm,!tinygo

package window

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"sync"
	"syscall/js"
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

func (document document) NewElementFromTemplate(name string, data interface{}) (Element, error) {
	buf := bytes.NewBuffer(nil)
	err := templates.ExecuteTemplate(buf, name, data)
	if err != nil {
		return Element(js.Null()), err
	}

	div := document.Call("createElement", "div")
	div.Set("innerHTML", strings.TrimSpace(buf.String()))

	v := div.Get("firstChild")
	if !v.Truthy() {
		return Element(js.Null()), fmt.Errorf("could not get created element")
	}

	return Element(v), nil
}

func (document document) CreateDocumentFragmentFromTemplate(name string, data interface{}) (DocumentFragment, error) {
	var buf bytes.Buffer
	err := templates.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return DocumentFragment(js.Null()), err
	}

	div := document.CreateElement("div")
	div.SetInnerHTML(strings.TrimSpace(buf.String()))

	f := NewDocumentFragment()
	f.ReplaceChildren(div.ChildNodes().NodeSlice()...)
	return f, nil
}

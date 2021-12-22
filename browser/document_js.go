package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type Document js.Value

func (d Document) Body() dom.Element { return Element(js.Value(d).Get("body")) }
func (d Document) Head() dom.Element { return Element(js.Value(d).Get("head")) }

func (d Document) NodeType() dom.NodeType         { return v(d).nodeType() }
func (d Document) Contains(other dom.Node) bool   { return v(d).contains(other) }
func (d Document) CloneNode(deep bool) dom.Node   { return v(d).cloneNode(deep) }
func (d Document) TextContent() string            { return v(d).textContent() }
func (d Document) IsSameNode(other dom.Node) bool { return v(d).isSameNode(other) }

func (d Document) Normalize() { v(d).normalize() }
func (d Document) GetElementsByTagName(name string) dom.ElementCollection {
	return v(d).getElementsByTagName(name)
}
func (d Document) GetElementsByClassName(name string) dom.ElementCollection {
	return v(d).getElementsByClassName(name)
}
func (d Document) QuerySelector(query string) dom.Element { return v(d).querySelector(query) }
func (d Document) QuerySelectorAll(query string) dom.NodeList {
	return v(d).querySelectorAll(query)
}
func (d Document) CreateElement(localName string) dom.Element {
	return v(d).createElement(localName)
}
func (d Document) CreateTextNode(text string) dom.Text { return v(d).createTextNode(text) }

func (d Document) Length() int { return v(d).length() }

//func (d Document) CreateDocumentFragment() DocumentFragment { return v(d).createDocumentFragment() }

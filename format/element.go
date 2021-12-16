package format

import "github.com/crhntr/window/dom"

func SetInnerHTML(e dom.Element, format string, args ...interface{}) {
	result := Sprintf(format, args...)
	e.SetInnerHTML(result)
}

func SetInnerText(e dom.Element, format string, args ...interface{}) {
	result := Sprintf(format, args...)
	e.SetInnerText(result)
}

func SetOuterHTML(e dom.Element, format string, args ...interface{}) {
	result := Sprintf(format, args...)
	e.SetOuterHTML(result)
}

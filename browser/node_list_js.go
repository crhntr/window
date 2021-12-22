package browser

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type NodeList js.Value

var _ dom.NodeList = NodeList{}

func (nl NodeList) Item(i int) dom.Node { return v(nl).item(i) }
func (nl NodeList) Length() int         { return v(nl).length() }

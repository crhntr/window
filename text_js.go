//go:build js

package window

import (
	"syscall/js"

	"github.com/crhntr/window/dom"
)

type Text js.Value

func (t Text) Data() string         { return v(t).data() }
func (t Text) Split(n int) dom.Text { return v(t).split(n) }
func (t Text) WholeText() string    { return v(t).wholeText() }

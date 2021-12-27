package attr

import (
	"testing"

	"github.com/crhntr/please"
)

func TestRange(t *testing.T) {
	t.Run("list of strings", func(t *testing.T) {
		var data struct {
			List []string
		}
		data.List = append(data.List, "One", "Two", "Three")

		el := CreateElement(t, `<ul data-attr-range="List"><li data-attr-text=""></li></ul>`)

		err := Hydrate(el, data)
		please.ExpectNilError(t, err)

		please.ExpectEqual(t, el.OuterHTML(), `<ul data-attr-range="List"><li data-attr-range-index="0">One</li><li data-attr-range-index="1">Two</li><li data-attr-range-index="2">Three</li></ul>`)
	})
}

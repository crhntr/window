package attr

import (
	"testing"

	"github.com/crhntr/please"
)

type dataWithBoolField struct {
	Show bool
}

type dataWithBoolMethod struct {
	show bool
}

func (el dataWithBoolMethod) Show() bool { return el.show }

func TestIf(t *testing.T) {
	const (
		trueHTML  = `<div data-attr-if="Show">Hello</div>`
		falseHTML = `<div data-attr-if="Show" data-attr-template="Hello" data-attr-if-result="false"></div>`
	)

	t.Run("field is true", func(t *testing.T) {
		el := CreateElement(t, trueHTML)
		err := Hydrate(el, dataWithBoolField{
			Show: true,
		})
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)
	})

	t.Run("field is false", func(t *testing.T) {
		el := CreateElement(t, trueHTML)
		err := Hydrate(el, dataWithBoolField{
			Show: false,
		})
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)
	})

	t.Run("method returns true", func(t *testing.T) {
		el := CreateElement(t, trueHTML)
		data := dataWithBoolMethod{
			show: true,
		}
		err := Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)
	})

	t.Run("method returns false", func(t *testing.T) {
		el := CreateElement(t, trueHTML)
		data := dataWithBoolMethod{
			show: false,
		}
		err := Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)
	})

	t.Run("toggle", func(t *testing.T) {
		el := CreateElement(t, trueHTML)
		data := dataWithBoolMethod{
			show: false,
		}
		err := Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)

		data.show = false
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)

		data.show = false
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = false
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)
	})

	t.Run("ssr was false", func(t *testing.T) {
		t.Run("field is true", func(t *testing.T) {
			el := CreateElement(t, falseHTML)
			err := Hydrate(el, dataWithBoolField{
				Show: true,
			})
			please.ExpectNilError(t, err)
			please.ExpectEqual(t, el.OuterHTML(), trueHTML)
		})
	})
}

func TestElse(t *testing.T) {
	const (
		templateHTML = `<div><div data-attr-if="Show">Hello</div><div data-attr-else="">Greetings</div></div>`
		trueHTML     = `<div><div data-attr-if="Show">Hello</div><div data-attr-else="" data-attr-template="Greetings"></div></div>`
		falseHTML    = `<div><div data-attr-if="Show" data-attr-template="Hello" data-attr-if-result="false"></div><div data-attr-else="">Greetings</div></div>`
	)

	t.Run("field is true", func(t *testing.T) {
		el := CreateElement(t, templateHTML)
		err := Hydrate(el, dataWithBoolField{
			Show: true,
		})
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)
	})

	t.Run("field is false", func(t *testing.T) {
		el := CreateElement(t, templateHTML)
		err := Hydrate(el, dataWithBoolField{
			Show: false,
		})
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)
	})

	t.Run("method returns true", func(t *testing.T) {
		el := CreateElement(t, templateHTML)
		data := dataWithBoolMethod{
			show: true,
		}
		err := Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)
	})

	t.Run("method returns false", func(t *testing.T) {
		el := CreateElement(t, templateHTML)
		data := dataWithBoolMethod{
			show: false,
		}
		err := Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)
	})

	t.Run("toggle", func(t *testing.T) {
		el := CreateElement(t, templateHTML)
		data := dataWithBoolMethod{
			show: false,
		}
		err := Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)

		data.show = false
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)

		data.show = false
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = false
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), falseHTML)

		data.show = true
		err = Hydrate(el, data)
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, el.OuterHTML(), trueHTML)
	})

	t.Run("ssr was false", func(t *testing.T) {
		t.Run("field is true", func(t *testing.T) {
			el := CreateElement(t, falseHTML)
			err := Hydrate(el, dataWithBoolField{
				Show: true,
			})
			please.ExpectNilError(t, err)
			please.ExpectEqual(t, el.OuterHTML(), trueHTML)
		})
	})
}

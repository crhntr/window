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
		templateHTML = /* language=html */ `<div data-attr-if="Show">Hello</div>`
		trueHTML     = templateHTML
		falseHTML    = /* language=html */ `<div data-attr-if="Show" data-attr-template="Hello" hidden=""></div>`
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

func TestElse(t *testing.T) {
	const (
		templateHTML = /* language=html */ `<div><div data-attr-if="Show">Hello</div><div data-attr-else="">Greetings</div></div>`
		trueHTML     = /* language=html */ `<div><div data-attr-if="Show">Hello</div><div data-attr-else="" data-attr-template="Greetings" hidden=""></div></div>`
		falseHTML    = /* language=html */ `<div><div data-attr-if="Show" data-attr-template="Hello" hidden=""></div><div data-attr-else="">Greetings</div></div>`
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

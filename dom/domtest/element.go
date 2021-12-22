package domtest

import (
	"testing"

	"github.com/crhntr/please"

	"github.com/crhntr/window/dom"
)

type CreateElementFunc func(t *testing.T, input string) dom.Element

func ElementTagName(t *testing.T, create CreateElementFunc) {
	t.Run("tagName", func(t *testing.T) {
		el, ok := create(t, `<div></div>`).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}
		result := el.TagName()
		please.ExpectEqual(t, result, "DIV")
	})
}

func ElementAttribute(t *testing.T, create CreateElementFunc) {
	t.Run("attributes", func(t *testing.T) {
		el, ok := create(t, `<div data-key="initial value"></div>`).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}

		t.Run("hasAttribute", func(t *testing.T) {
			please.ExpectTrue(t, el.HasAttribute("data-key"))
			please.ExpectFalse(t, el.HasAttribute("data-not-set"))
		})

		t.Run("getAttribute", func(t *testing.T) {
			please.ExpectEqual(t,
				el.GetAttribute("data-key"),
				"initial value",
			)
		})

		t.Run("setAttribute update existing", func(t *testing.T) {
			el.SetAttribute("data-key", "second value")
			please.ExpectEqual(t,
				el.GetAttribute("data-key"),
				"second value",
			)
		})

		t.Run("removeAttribute", func(t *testing.T) {
			el.RemoveAttribute("data-key")
			please.ExpectEqual(t,
				el.GetAttribute("data-key"),
				"",
			)
			please.ExpectFalse(t, el.HasAttribute("data-key"))
		})

		t.Run("setAttribute insert new", func(t *testing.T) {
			el.SetAttribute("data-new", "change")
			please.ExpectEqual(t,
				el.GetAttribute("data-new"),
				"change",
			)
		})

		t.Run("toggleAttribute", func(t *testing.T) {
			please.ExpectTrue(t, el.ToggleAttribute("data-boolean"))
			please.ExpectTrue(t, el.HasAttribute("data-boolean"))
			please.ExpectFalse(t, el.ToggleAttribute("data-boolean"))
			please.ExpectFalse(t, el.HasAttribute("data-boolean"))
		})

		t.Run("className", func(t *testing.T) {
			please.ExpectEqual(t, el.ClassName(), "")

			el.SetAttribute("class", "style1")
			please.ExpectEqual(t, el.ClassName(), "style1")

			el.SetAttribute("class", "style1  style2")
			please.ExpectEqual(t, el.ClassName(), "style1  style2")

			el.SetAttribute("class", "style1 style2")
			please.ExpectEqual(t, el.ClassName(), "style1 style2")
		})

		t.Run("id", func(t *testing.T) {
			please.ExpectEqual(t, el.ID(), "")

			el.SetAttribute("id", "id-01")
			please.ExpectEqual(t, el.ID(), "id-01")
		})
	})
}

func ElementInnerHTML(t *testing.T, create CreateElementFunc) {
	makeDiv := func(t *testing.T) dom.Element {
		innerHTML := /* language=html */ `<div><p>loading<span>...</span></p><p>please wait</p></div>`

		ul, ok := create(t, innerHTML).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}
		return ul
	}

	t.Run("innerHTML", func(t *testing.T) {
		t.Run("get", func(t *testing.T) {
			ul := makeDiv(t)
			s := ul.InnerHTML()
			please.ExpectEqual(t, s /* language=html */, `<p>loading<span>...</span></p><p>please wait</p>`)
		})

		t.Run("set", func(t *testing.T) {
			div := makeDiv(t)
			div.SetInnerHTML( /* language=html */ `<h1><a href="/">Hello, world!</a></h1><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>`)
			please.ExpectEqual(t, div.ChildNodes().Length(), 2)

			h1 := div.FirstElementChild()
			please.ExpectEqual(t, h1.TagName(), "H1")

			a := h1.FirstElementChild()
			please.ExpectEqual(t, a.GetAttribute("href"), "/")

			text := a.FirstChild().(dom.Text)
			please.ExpectEqual(t, text.Data(), "Hello, world!")

			p := div.LastElementChild()
			please.ExpectEqual(t, p.TagName(), "P")
			paragraph := p.FirstChild().(dom.Text)
			please.ExpectEqual(t, paragraph.Data(), "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		})
	})
}

func ElementOuterHTML(t *testing.T, create CreateElementFunc) {
	makeDiv := func(t *testing.T) dom.Element {
		innerHTML := /* language=html */ `<div><p>loading<span>...</span></p><p>please wait</p></div>`

		ul, ok := create(t, innerHTML).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}
		return ul
	}

	t.Run("outerHTML", func(t *testing.T) {
		t.Run("get", func(t *testing.T) {
			ul := makeDiv(t)
			s := ul.OuterHTML()
			please.ExpectEqual(t, s /* language=html */, `<div><p>loading<span>...</span></p><p>please wait</p></div>`)
		})

		t.Run("set", func(t *testing.T) {
			t.Run("multiple elements", func(t *testing.T) {
				div := makeDiv(t).(dom.Element)
				p := div.FirstElementChild()
				please.ExpectEqual(t, div.ChildElementCount(), 2)

				p.SetOuterHTML( /* language=html */ `<h1><a href="/">Hello, world!</a></h1><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>`)
				please.ExpectEqual(t, div.ChildElementCount(), 3)
				firstChild := div.FirstElementChild()
				please.ExpectEqual(t, firstChild.OuterHTML(), `<h1><a href="/">Hello, world!</a></h1>`)

				please.ExpectEqual(t, div.OuterHTML(), `<div><h1><a href="/">Hello, world!</a></h1><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p><p>please wait</p></div>`)
			})

			t.Run("a single element", func(t *testing.T) {
				div := makeDiv(t)
				p := div.FirstElementChild()
				p.SetOuterHTML( /* language=html */ `<h1><a href="/">Hello, world!</a></h1>`)

				t.Run("the element is replaced in the document", func(t *testing.T) {
					h1 := div.FirstChild().(dom.Element)
					please.ExpectEqual(t, h1.OuterHTML(), `<h1><a href="/">Hello, world!</a></h1>`)
				})

				t.Run("the variable whose outerHTML property was set will still holds a reference to the original element", func(t *testing.T) {
					please.ExpectEqual(t, p.OuterHTML(), "<p>loading<span>...</span></p>")
				})
			})
		})
	})
}

func ElementQueries(t *testing.T, create CreateElementFunc) {
	makeUL := func(t *testing.T) dom.Element {
		innerHTML := /* language=html */ `<ul>
	<li id="empty"></li>
	<li id="middle">
		<a href="https://example.com">Example</a>
	</li>
	<li id="input">
		<label>
			ome description
			<i class="fa fa-cloud">
			</i><input name="item">	
		</label>
		<ol class="errors">
			<li class="error client-generated latest">Error 1</li>
			<li class="error client-generated">Error 2</li>
			<li class="error server-generated">Error 3</li>
		</ol>
	</li>
</ul>`

		ul, ok := create(t, innerHTML).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}
		return ul
	}

	t.Run("getElementsByTagName", func(t *testing.T) {
		ul := makeUL(t)
		list := ul.GetElementsByTagName("li")
		please.ExpectEqual(t, list.Length(), 6)
	})

	t.Run("getElementsByClassName", func(t *testing.T) {
		ul := makeUL(t)
		list := ul.GetElementsByClassName("client-generated    error")
		please.ExpectEqual(t, list.Length(), 2)
	})
}

func ElementTextContent(t *testing.T, create CreateElementFunc) {
	makeDiv := func(t *testing.T) dom.Element {
		innerHTML := /* language=html */ `<div><p>loading<span>...</span></p><script>alert("loading!")</script></div>`

		ul, ok := create(t, innerHTML).(dom.Element)
		if !ok {
			t.Errorf("result from create is not a dom.Element")
		}
		return ul
	}

	div := makeDiv(t)

	please.ExpectEqual(t, div.TextContent(), `loading...alert("loading!")`)
}

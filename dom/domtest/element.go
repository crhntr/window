package domtest

import (
	"testing"

	"golang.org/x/net/html/atom"

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
		createElement := func(t *testing.T) dom.Element {
			el, ok := create(t, `<div data-key="initial value"></div>`).(dom.Element)
			if !ok {
				t.Errorf("result from create is not a dom.Element")
			}
			return el
		}

		t.Run("hasAttribute", func(t *testing.T) {
			el := createElement(t)
			please.ExpectTrue(t, el.HasAttribute("data-key"))
			please.ExpectFalse(t, el.HasAttribute("data-not-set"))
		})

		t.Run("getAttribute", func(t *testing.T) {
			el := createElement(t)
			please.ExpectEqual(t,
				el.Attribute("data-key"),
				"initial value",
			)
		})

		t.Run("setAttribute update existing", func(t *testing.T) {
			el := createElement(t)
			el.SetAttribute("data-key", "second value")
			please.ExpectEqual(t,
				el.Attribute("data-key"),
				"second value",
			)
		})

		t.Run("setAttribute multiple times", func(t *testing.T) {
			el := createElement(t)
			el.SetAttribute("data-key", "a")
			el.SetAttribute("data-key", "b")
			el.SetAttribute("data-key", "c")
			please.ExpectEqual(t,
				el.Attribute("data-key"),
				"c",
			)
		})

		t.Run("removeAttribute", func(t *testing.T) {
			el := createElement(t)
			el.RemoveAttribute("data-key")
			please.ExpectEqual(t,
				el.Attribute("data-key"),
				"",
			)
			please.ExpectFalse(t, el.HasAttribute("data-key"))
		})

		t.Run("setAttribute insert new", func(t *testing.T) {
			el := createElement(t)
			el.SetAttribute("data-new", "change")
			please.ExpectEqual(t,
				el.Attribute("data-new"),
				"change",
			)
		})

		t.Run("toggleAttribute", func(t *testing.T) {
			el := createElement(t)
			please.ExpectTrue(t, el.ToggleAttribute("data-boolean"))
			please.ExpectTrue(t, el.HasAttribute("data-boolean"))
			please.ExpectFalse(t, el.ToggleAttribute("data-boolean"))
			please.ExpectFalse(t, el.HasAttribute("data-boolean"))
		})

		t.Run("className", func(t *testing.T) {
			el := createElement(t)
			please.ExpectEqual(t, el.ClassName(), "")

			el.SetAttribute("class", "style1")
			please.ExpectEqual(t, el.ClassName(), "style1")

			el.SetAttribute("class", "style1  style2")
			please.ExpectEqual(t, el.ClassName(), "style1  style2")

			el.SetAttribute("class", "style1 style2")
			please.ExpectEqual(t, el.ClassName(), "style1 style2")
		})

		t.Run("id", func(t *testing.T) {
			el := createElement(t)
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
			please.ExpectEqual(t, a.Attribute("href"), "/")

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
		innerHTML := /* language=html */ `<ul id="page-list">
	<li id="empty"></li>
	<li id="middle">
		<a href="https://example.com">Example</a>
	</li>
	<li id="input">
		<label>
			some description
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

	t.Run("getElementById", func(t *testing.T) {
		ul := makeUL(t)
		empty := ul.GetElementByID("empty")
		please.ExpectTrue(t, empty != nil)
		please.ExpectEqual(t, empty.TagName(), "LI")
		please.ExpectEqual(t, empty.ChildElementCount(), 0)

		pageList := ul.GetElementByID("page-list")
		please.ExpectEqual(t, pageList.TagName(), "UL")
		please.ExpectEqual(t, pageList.ChildElementCount(), 3)
	})

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

	t.Run("get", func(t *testing.T) {
		div := makeDiv(t)

		please.ExpectEqual(t, div.TextContent(), `loading...alert("loading!")`)
	})
	t.Run("set", func(t *testing.T) {
		div := makeDiv(t)
		div.SetTextContent("Hello, world!")
		please.ExpectEqual(t, div.OuterHTML(), `<div>Hello, world!</div>`)
	})
}

func ElementParent(t *testing.T, createParent CreateParentNodeFunc, createEl CreateChildNodeFunc) {
	t.Run("ElementParent", func(t *testing.T) {
		t.Run("created parent node is empty", func(t *testing.T) {
			parent := createParent(t).(dom.ElementParent)
			please.ExpectEqual(t, parent.ChildElementCount(), 0)
			please.ExpectEqual(t, parent.FirstElementChild(), nil)
			please.ExpectEqual(t, parent.LastElementChild(), nil)
			please.ExpectEqual(t, parent.Children().Length(), 0)
		})

		t.Run("createEl returns an element", func(t *testing.T) {
			el, ok := createEl(t).(dom.Element)
			please.ExpectTrue(t, ok)

			if el.TagName() == atom.Cite.String() {
				t.Errorf(`expected tag name to not be "cite"`)
			}
		})

		t.Run("add Element", func(t *testing.T) {
			parent := createParent(t)
			child := createEl(t)
			parent.Append(child)

			please.ExpectEqual(t, parent.ChildElementCount(), 1)
			please.ExpectTrue(t, parent.FirstElementChild().IsSameNode(child))
			please.ExpectTrue(t, parent.LastElementChild().IsSameNode(child))
			please.ExpectEqual(t, parent.Children().Length(), 1)
		})

		t.Run("add two Elements", func(t *testing.T) {
			parent := createParent(t)
			child1 := createEl(t).(dom.Element)
			child2 := createEl(t).(dom.Element)
			parent.Append(child1, child2)

			please.ExpectEqual(t, parent.ChildElementCount(), 2)
			please.ExpectTrue(t, parent.FirstElementChild().IsSameNode(child1))
			please.ExpectTrue(t, parent.LastElementChild().IsSameNode(child2))
			please.ExpectTrue(t, child2.PreviousSibling().IsSameNode(child1))
			please.ExpectTrue(t, child1.NextSibling().IsSameNode(child2))
			please.ExpectEqual(t, parent.Children().Length(), 2)
			please.ExpectEqual(t, parent.GetElementsByTagName(child1.TagName()).Length(), 2)
		})

		t.Run("add three Elements with classes", func(t *testing.T) {
			parent := createParent(t)

			child1 := createEl(t).(dom.Element)
			child1.SetAttribute("class", "first child")
			child2 := createEl(t).(dom.Element)
			child2.SetAttribute("class", "child last")

			parent.Append(child1, child2)

			please.ExpectEqual(t, parent.ChildElementCount(), 2)
			please.ExpectTrue(t, parent.FirstElementChild().IsSameNode(child1))
			please.ExpectTrue(t, parent.LastElementChild().IsSameNode(child2))
			please.ExpectTrue(t, child2.PreviousSibling().IsSameNode(child1))
			please.ExpectTrue(t, child1.NextSibling().IsSameNode(child2))
			please.ExpectEqual(t, parent.Children().Length(), 2)

			please.ExpectEqual(t, parent.GetElementsByTagName(child1.TagName()).Length(), 2)
			please.ExpectEqual(t, parent.GetElementsByTagName(atom.Cite.String()).Length(), 0)

			please.ExpectEqual(t, parent.GetElementsByClassName("first child").Length(), 1)
			please.ExpectEqual(t, parent.GetElementsByClassName("child").Length(), 2)
			please.ExpectEqual(t, parent.GetElementsByClassName("random-class-name").Length(), 0)
		})
	})
}

func TestElementInsertAdjacentHTML(t *testing.T, createEl CreateChildNodeFunc) {
	t.Run("InsertAdjacentHTML", func(t *testing.T) {
		setup := func(t *testing.T, createEl CreateChildNodeFunc) (dom.Element, dom.Element, dom.Element) {
			parent := createEl(t).(dom.Element)
			parent.SetInnerHTML(`<div id="a1"></div><div id="a3"><span></span></div><div id="a5"></div>`)
			a1 := parent.FirstChild().(dom.Element)
			a3 := a1.NextSibling().(dom.Element)
			a5 := a3.NextSibling().(dom.Element)

			span := a3.FirstChild()
			please.ExpectEqual(t, span.(dom.Element).TagName(), "SPAN")
			return a1, a3, a5
		}

		t.Run(dom.PositionBeforeBegin.String(), func(t *testing.T) {
			a1, a3, _ := setup(t, createEl)
			a3.InsertAdjacentHTML(dom.PositionBeforeBegin, `<div id="a2"></div>`)
			please.ExpectEqual(t, a3.PreviousSibling().(dom.Element).Attribute("id"), "a2")
			please.ExpectEqual(t, a1.NextSibling().(dom.Element).Attribute("id"), "a2")
		})
		t.Run(dom.PositionAfterBegin.String(), func(t *testing.T) {
			t.Run("el is not empty", func(t *testing.T) {
				_, a3, _ := setup(t, createEl)
				a3.InsertAdjacentHTML(dom.PositionAfterBegin, `<span id="first"></span>`)
				please.ExpectEqual(t, a3.FirstChild().(dom.Element).Attribute("id"), "first")
			})

			t.Run("el isempty", func(t *testing.T) {
				parent := createEl(t).(dom.Element)
				parent.SetInnerHTML(`<div id="empty"></div>`)
				empty := parent.FirstChild().(dom.Element)
				empty.InsertAdjacentHTML(dom.PositionAfterBegin, `<span id="added"></span>`)
				please.ExpectEqual(t, empty.FirstChild().(dom.Element).Attribute("id"), "added")
			})
		})
		t.Run(dom.PositionBeforeEnd.String(), func(t *testing.T) {
			_, a3, _ := setup(t, createEl)
			a3.InsertAdjacentHTML(dom.PositionBeforeEnd, `<span id="last"></span>`)
			please.ExpectEqual(t, a3.LastChild().(dom.Element).Attribute("id"), "last")
		})
		t.Run(dom.PositionAfterEnd.String(), func(t *testing.T) {
			t.Run("el is not the only Element in parent", func(t *testing.T) {
				_, a3, a5 := setup(t, createEl)
				a3.InsertAdjacentHTML(dom.PositionAfterEnd, `<div id="a4"></div>`)
				please.ExpectEqual(t, a3.NextSibling().(dom.Element).Attribute("id"), "a4")
				please.ExpectEqual(t, a5.PreviousSibling().(dom.Element).Attribute("id"), "a4")
			})

			t.Run("el is the only Element in parent", func(t *testing.T) {
				parent := createEl(t).(dom.Element)
				parent.SetInnerHTML(`<div id="only-element"></div>`)
				onlyEl := parent.FirstChild().(dom.Element)
				onlyEl.InsertAdjacentHTML(dom.PositionAfterEnd, `<div id="added"></div>`)
				please.ExpectEqual(t, onlyEl.NextSibling().(dom.Element).Attribute("id"), "added")
				please.ExpectEqual(t, parent.LastChild().(dom.Element).Attribute("id"), "added")
			})
		})
	})
}

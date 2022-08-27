package selector

import (
	"testing"

	"github.com/crhntr/please"
)

func TestParse(t *testing.T) {
	t.SkipNow()
	t.Run("a simple class", func(t *testing.T) {
		input := ".simple"
		matcher, err := Parse(input)
		please.ExpectNilError(t, err)
		m, ok := matcher.(ClassMatcher)
		please.ExpectTrue(t, ok)
		please.ExpectEqual(t, string(m), "simple")
	})
	t.Run("a simple id", func(t *testing.T) {
		input := "#simple"
		matcher, err := Parse(input)
		please.ExpectNilError(t, err)
		m, ok := matcher.(IDMatcher)
		please.ExpectTrue(t, ok)
		please.ExpectEqual(t, string(m), "simple")
	})
	t.Run("a simple tag", func(t *testing.T) {
		input := "simple"
		matcher, err := Parse(input)
		please.ExpectNilError(t, err)
		m, ok := matcher.(TagMatcher)
		please.ExpectTrue(t, ok)
		please.ExpectEqual(t, m.Name, "simple")
	})
	t.Run("a simple namespaced tag", func(t *testing.T) {
		input := "example|simple"
		matcher, err := Parse(input)
		please.ExpectNilError(t, err)
		m, ok := matcher.(TagMatcher)
		please.ExpectTrue(t, ok)
		please.ExpectEqual(t, m.Name, "simple")
		please.ExpectEqual(t, m.Namespace, "example")
	})
	t.Run("any element", func(t *testing.T) {
		input := "*"
		matcher, err := Parse(input)
		please.ExpectNilError(t, err)
		_, ok := matcher.(AnyElementMatcher)
		please.ExpectTrue(t, ok)
	})
	t.Run("compound matcher", func(t *testing.T) {
		t.Run("tag with class", func(t *testing.T) {
			input := "div.simple"
			matcher, err := Parse(input)
			please.ExpectNilError(t, err)

			m, ok := matcher.(CompoundMatcher)
			please.ExpectTrue(t, ok)
			if !please.ExpectEqual(t, len(m), 2) {
				return
			}

			tm, ok := m[0].(TagMatcher)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, tm.Name, "div")

			cm, ok := m[1].(ClassMatcher)
			t.Logf("%T\n", cm)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, string(cm), "simple")
		})
		t.Run("tag with id", func(t *testing.T) {
			input := "div#simple"
			matcher, err := Parse(input)
			please.ExpectNilError(t, err)

			m, ok := matcher.(CompoundMatcher)
			please.ExpectTrue(t, ok)
			if !please.ExpectEqual(t, len(m), 2) {
				return
			}

			tm, ok := m[0].(TagMatcher)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, tm.Name, "div")

			im, ok := m[1].(IDMatcher)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, string(im), "simple")
		})
		t.Run("id and class", func(t *testing.T) {
			input := "#simple.big"
			matcher, err := Parse(input)
			please.ExpectNilError(t, err)

			m, ok := matcher.(CompoundMatcher)
			please.ExpectTrue(t, ok)
			if !please.ExpectEqual(t, len(m), 2) {
				return
			}

			im, ok := m[0].(IDMatcher)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, string(im), "simple")

			cm, ok := m[1].(ClassMatcher)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, string(cm), "big")
		})
		t.Run("class and id", func(t *testing.T) {
			input := ".big#simple"
			matcher, err := Parse(input)
			please.ExpectNilError(t, err)

			m, ok := matcher.(CompoundMatcher)
			please.ExpectTrue(t, ok)
			if !please.ExpectEqual(t, len(m), 2) {
				return
			}

			cm, ok := m[0].(ClassMatcher)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, string(cm), "big")

			im, ok := m[1].(IDMatcher)
			please.ExpectTrue(t, ok)
			please.ExpectEqual(t, string(im), "simple")
		})
	})
	t.Run("ancestor matcher", func(t *testing.T) {
		input := "ul li"
		_, err := Parse(input)
		please.ExpectNilError(t, err)
	})
}

package scanner

import (
	"strings"
	"testing"

	"github.com/crhntr/please"
)

func TestPeek(t *testing.T) {
	t.Run("any", func(t *testing.T) {
		s := New(strings.NewReader("*"))
		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeSymbol)
		please.ExpectEqual(t, tok.Data, "*")
	})
	t.Run("div", func(t *testing.T) {
		s := New(strings.NewReader("div"))
		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "div")
	})
	t.Run("class dot", func(t *testing.T) {
		s := New(strings.NewReader(".div"))
		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeSymbol)
		please.ExpectEqual(t, tok.Data, ".")
		please.ExpectEqual(t, len(s.peekedRunes), 0)
	})
	t.Run("tag class", func(t *testing.T) {
		s := New(strings.NewReader("div.greeting"))
		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "div")
		please.ExpectEqual(t, s.peekedRunes, []rune{'.'})
	})
	t.Run("tag id", func(t *testing.T) {
		s := New(strings.NewReader("div#greeting"))
		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "div")
		please.ExpectEqual(t, s.peekedRunes, []rune{'#'})
	})
	t.Run("peeked ident rune", func(t *testing.T) {
		s := New(strings.NewReader("ag"))
		s.peekedRunes = []rune{'t'}
		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "tag")
		please.ExpectEqual(t, s.peekedRunes, []rune{})
	})
	t.Run("peeked symbol rune", func(t *testing.T) {
		s := New(strings.NewReader("class"))
		s.peekedRunes = []rune{'.'}
		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeSymbol)
		please.ExpectEqual(t, tok.Data, ".")
		please.ExpectEqual(t, s.peekedRunes, []rune{})
	})
	t.Run("namespaced element with class and id", func(t *testing.T) {
		s := New(strings.NewReader("svg|circle.red#first"))

		tok, err := s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "svg")
		please.ExpectEqual(t, s.peekedRunes, []rune{'|'})

		tok, err = s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeSymbol)
		please.ExpectEqual(t, tok.Data, "|")
		please.ExpectEqual(t, len(s.peekedRunes), 0)

		tok, err = s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "circle")
		please.ExpectEqual(t, s.peekedRunes, []rune{'.'})

		tok, err = s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeSymbol)
		please.ExpectEqual(t, tok.Data, ".")
		please.ExpectEqual(t, len(s.peekedRunes), 0)

		tok, err = s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "red")
		please.ExpectEqual(t, s.peekedRunes, []rune{'#'})

		tok, err = s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeSymbol)
		please.ExpectEqual(t, tok.Data, "#")
		please.ExpectEqual(t, len(s.peekedRunes), 0)

		tok, err = s.Scan()
		please.ExpectNilError(t, err)
		please.ExpectEqual(t, tok.Type, TokenTypeIdent)
		please.ExpectEqual(t, tok.Data, "first")
		please.ExpectEqual(t, len(s.peekedRunes), 0)
	})
}

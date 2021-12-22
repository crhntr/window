package selector

import (
	"strings"
)

func scan(s string) ([]token, error) {

	return nil, nil
}

type tokenType int

const (
	identToken tokenType = iota
	whitespaceToken
)

type token struct {
	t tokenType
	v string
}

const (
	newlineCharacter     = "\u000A"
	whitespaceCharacters = "\u0009\u0020" + newlineCharacter
)

func consumeWhitespace(s string) (token, string, bool) {
	var (
		i int
		c rune
	)
	for i, c = range s {
		if !strings.ContainsRune(whitespaceCharacters, c) {
			break
		}
	}
	if i == 0 {
		return token{}, s, false
	}
	return token{
		t: whitespaceToken,
		v: s[:i],
	}, s, true
}

func consumeIdent(s string) (token, string, bool) {
	var (
		v  string
		si = 0
	)

	if len(s) == 0 {
		return token{}, s, false
	}

	if !strings.HasPrefix(s, "--") {
		if s[si] == '-' {
			v += "-"
			si++
		}
		if s[si] == '\\' {
			v = "\\" + readHexCharacters(s[:5])
			si = len(v)
		} else {

			isAZazOrNonAscii(s[si:])
		}
	}

	if || fc == '\\' {

	if fc == '\\' {
	v = "\\" + readHexCharacters(s[:5])
	} else {
	v += s[:runeLen(fc)]
	}
	}

	consumed := len(v)

	for i, c := range s[consumed:] {

	}

	return token{
	t: identToken,
	v: v,
	}, s, true
}

func isAZazOrNonAscii(c rune) bool {
	return (c > 'a' && c < 'z') && (c > 'A' && c < 'Z') || c > '\u0080'
}

func firstRune(s string) rune {
	for _, c := range s {
		return c
	}
	return 0
}

func runeLen(c rune) int {
	if c == 0 {
		return 0
	}
	if (c & 0x7F) == c {
		return 1
	}
	if (c & 0xDFBF) == c {
		return 2
	}
	if (c & 0xEF_BF_BF) == c {
		return 3
	}
	if (c & 0x77_BF_BF_BF) == c { // rune is signed, so it should be 0xF7_BF_BF_BF but this works
		return 3
	}
	return 4
}

func readHexCharacters(s string) string {
	var (
		i int
		c rune
	)
	for i, c = range s {
		if !isHexCharacter(c) {
			break
		}
	}
	return s[:i]
}

func isHexCharacter(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

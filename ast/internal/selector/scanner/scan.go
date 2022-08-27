package scanner

import (
	"io"
)

type Scanner struct {
	r            io.RuneReader
	peekedRunes  []rune
	peekedTokens []Token
}

func New(rr io.RuneReader) *Scanner {
	return &Scanner{r: rr}
}

type TokenType uint8

const (
	TokenTypeError TokenType = iota
	TokenTypeEOF
	TokenTypeSymbol
	TokenTypeWhiteSpace
	TokenTypeIdent
)

type Token struct {
	Type TokenType
	//Pos  int
	Data string
}

func (tt TokenType) IsErr() bool {
	return tt == TokenTypeError || tt == TokenTypeEOF
}

func (sc *Scanner) Scan() (Token, error) {
	if len(sc.peekedTokens) > 0 {
		t := sc.peekedTokens[len(sc.peekedTokens)-1]
		sc.peekedTokens = sc.peekedTokens[:len(sc.peekedTokens)-1]
		return t, nil
	}
	return sc.scan()
}

func (sc *Scanner) scan() (Token, error) {
	type stateID uint8
	const (
		initial stateID = iota
		whitespace
		ident
	)

	var (
		state    stateID
		consumed string
	)

	for {
		ch, err := sc.next()
		if err != nil && err != io.EOF {
			return Token{}, err
		}
		switch state {
		case initial:
			if err == io.EOF {
				return Token{Type: TokenTypeEOF}, err
			}
			switch ch {
			case '.', '#', ':', '[', ']', '|', '*':
				return Token{Type: TokenTypeSymbol, Data: string(ch)}, nil
			case ' ', '\t', '\n':
				state = whitespace
				consumed += string(ch)
			default:
				state = ident
				consumed += string(ch)
			}
		case whitespace:
			if !isWhitespaceRune(ch) || err == io.EOF {
				sc.put(ch)
				return Token{Type: TokenTypeWhiteSpace, Data: consumed}, nil
			}
			consumed += string(ch)
		case ident:
			if !isIdentRune(ch) || err == io.EOF {
				sc.put(ch)
				return Token{Type: TokenTypeIdent, Data: consumed}, nil
			}
			consumed += string(ch)
		}
	}
}

func (sc *Scanner) Peek() (Token, error) {
	p, err := sc.scan()
	sc.peekedTokens = append(sc.peekedTokens, p)
	return p, err
}

func (sc *Scanner) put(c rune) {
	if c != 0 {
		sc.peekedRunes = append(sc.peekedRunes, c)
	}
}

func (sc *Scanner) next() (rune, error) {
	if len(sc.peekedRunes) > 0 {
		ch := sc.peekedRunes[len(sc.peekedRunes)-1]
		sc.peekedRunes = sc.peekedRunes[:len(sc.peekedRunes)-1]
		return ch, nil
	}
	ch, _, err := sc.r.ReadRune()
	return ch, err
}

func isWhitespaceRune(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func isIdentRune(c rune) bool {
	return c == 0 || (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '1') ||
		c == '-'
}

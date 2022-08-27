package selector

import (
	"errors"
	"io"
	"strings"

	"golang.org/x/net/html"

	"github.com/crhntr/window/ast/internal/selector/scanner"
)

func Parse(s string) (Matcher, error) {
	sc := scanner.New(strings.NewReader(s))

	var result CompoundMatcher

	for {
		m, err := parseNodeMatcher(sc)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		result = append(result, m)
	}

	if len(result) == 1 {
		return result[0], nil
	}

	return result, nil
}

func parseNodeMatcher(sc *scanner.Scanner) (Matcher, error) {
	tok, err := sc.Scan()
	if err != nil {
		return nil, err
	}

	switch tok.Type {
	case scanner.TokenTypeError:
		break
	case scanner.TokenTypeSymbol:
		switch tok.Data {
		case "*":
			return AnyElementMatcher{}, nil
		case "#":
			nameTok, scanErr := sc.Scan()
			if scanErr != nil {
				return nil, scanErr
			}
			return IDMatcher(nameTok.Data), nil
		case ".":
			nameTok, scanErr := sc.Scan()
			if scanErr != nil {
				return nil, scanErr
			}
			return ClassMatcher(nameTok.Data), nil
		}
	case scanner.TokenTypeIdent:
		p, err := sc.Peek()
		if err != nil && err != io.EOF {
			return nil, err
		}
		if p.Type == scanner.TokenTypeSymbol && p.Data == "|" {
			t2, err := sc.Scan()
			if err != nil {
				return nil, err
			}
			if t2.Type != scanner.TokenTypeIdent {
				return nil, errors.New("expected an identifier")
			}
			return TagMatcher{Namespace: t2.Data, Name: tok.Data}, nil
		}
		sc.Scan()
		return TagMatcher{Name: tok.Data}, nil
	}

	return nil, errors.New("failed to parse a matcher")
}

type Matcher interface {
	Matches(node *html.Node) bool
}

type AnyElementMatcher struct{}

func (m AnyElementMatcher) Matches(node *html.Node) bool {
	return node.Type == html.ElementNode
}

type ClassMatcher string

func (m ClassMatcher) Matches(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}
	for _, a := range node.Attr {
		if a.Key != "class" {
			continue
		}
		classes := strings.Fields(a.Val)
		for _, c := range classes {
			if c == string(m) {
				return true
			}
		}
	}
	return false
}

type IDMatcher string

func (m IDMatcher) Matches(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}
	for _, a := range node.Attr {
		if a.Key != "id" {
			continue
		}
		id := strings.TrimSpace(a.Val)
		return id == string(m)
	}
	return false
}

type TagMatcher struct {
	Namespace,
	Name string
}

func (m TagMatcher) Matches(node *html.Node) bool {
	return node.Type == html.ElementNode && strings.EqualFold(node.Data, m.Name) && strings.EqualFold(node.Namespace, m.Namespace)
}

type CompoundMatcher []Matcher

func (m CompoundMatcher) Matches(node *html.Node) bool {
	for _, matcher := range m {
		if !matcher.Matches(node) {
			return false
		}
	}
	return true
}

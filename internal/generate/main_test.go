package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strings"
	"testing"
)

func TestMethod(t *testing.T) {
	method := Method{
		Identifier: Identifier{
			Name: "prepend",
		},
		Parameters: []Parameter{
			{
				ID: Identifier{
					Name: "nodes",
				},
				Type:       "Node",
				IsVariadic: true,
			},
		},
	}

	result := method.functionDeclaration(ast.NewIdent("Element"))

	var buf bytes.Buffer
	err := format.Node(&buf, &token.FileSet{}, result)
	if err != nil {
		t.Fatal(err)
	}
	code := buf.String()
	code = strings.ReplaceAll(code, "\n", "")
	expected, _ := format.Source([]byte(`func (val Element) Prepend(nodes ...Node) {js.Value(val).Call("prepend", nodes...)}`))
	codeBuf, _ := format.Source([]byte(code))
	if strings.TrimSpace(string(codeBuf)) != strings.TrimSpace(string(expected)) {
		t.Logf("\n%q\n!=\n%q", code, expected)
	}
}

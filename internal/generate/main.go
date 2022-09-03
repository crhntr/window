package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func main() {
	generationInput, readFileErr := os.ReadFile("dom.yml")
	if readFileErr != nil {
		panic(readFileErr)
	}

	d := yaml.NewDecoder(bytes.NewReader(generationInput))
	d.KnownFields(true)
	var specs SpecificationList
	for {
		var enc encodedSpec
		err := d.Decode(&enc)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		spec, err := enc.declarationProviderFactory()
		if err != nil {
			panic(err)
		}
		specs = append(specs, spec)
	}

	err := specs.resolveInheritance()
	if err != nil {
		panic(err)
	}

	fileNode, err := generate(specs)
	if err != nil {
		panic(err)
	}

	fileSet := new(token.FileSet)
	err = ast.Print(fileSet, fileNode)
	if err != nil {
		panic(err)
	}

	generatedFile, err := os.Create("generated.go")
	if err != nil {
		panic(err)
	}
	defer closeAndIgnoreError(generatedFile)
	err = format.Node(generatedFile, fileSet, fileNode)
	if err != nil {
		panic(err)
	}
}

func closeAndIgnoreError(c io.Closer) { _ = c.Close() }

type Specification interface {
	ID() Identifier
}

type encodedSpec struct {
	Identifier `yaml:",inline"`
	Kind       string    `yaml:"kind"`
	Parent     string    `yaml:"parent"`
	Spec       yaml.Node `yaml:"spec"`
}

func (encoded encodedSpec) ID() Identifier { return encoded.Identifier }

func (encoded encodedSpec) declarationProviderFactory() (Specification, error) {
	var spec Specification
	switch encoded.Kind {
	case "class":
		spec = &Class{encodedSpec: encoded}
	case "interface":
		spec = &Interface{encodedSpec: encoded}
	}
	err := encoded.Spec.Decode(spec)
	if err != nil {
		return nil, err
	}
	return spec, nil
}

type Identifier struct {
	Name         string `yaml:"name"`
	NameOverride string `yaml:"name_override"`
}

func (n Identifier) String() string {
	if n.NameOverride != "" {
		return n.NameOverride
	}
	return n.Name
}

func (n Identifier) Exported() string {
	return titleCase.String(n.String())
}

func (n Identifier) ExportedIdentifier() *ast.Ident {
	return ast.NewIdent(n.Exported())
}

type Interface struct {
	encodedSpec `yaml:"-"`
	Properties  []Property `yaml:"properties"`
	Methods     []Method   `yaml:"methods"`
}

type Class struct {
	encodedSpec `yaml:"-"`
	Properties  []Property `yaml:"properties"`
	Methods     []Method   `yaml:"methods"`
}

type Property struct {
	definition string
	Identifier `yaml:",inline"`
	Type       string `yaml:"type"`
	IsArray    bool   `yaml:"isArray"`
	WrapResult string `yaml:"wrap"`
}

type Method struct {
	definition string
	Identifier `yaml:",inline"`
	Parameters []Parameter `yaml:"params"`
	Result     struct {
		Type string `yaml:"type"`
		Wrap string `yaml:"wrap"`
	} `yaml:"result"`
}

type Parameter struct {
	ID         Identifier `yaml:",inline"`
	Type       string     `yaml:"type"`
	IsVariadic bool       `yaml:"IsVariadic"`
}

var titleCase = cases.Title(language.AmericanEnglish, cases.NoLower)

func (property Property) functionDeclaration(receiverIdentifier *ast.Ident) *ast.FuncDecl {
	receiver := ast.NewIdent("val")

	retExprPrefix, retExpSuffix := wrapResult(property.Type, property.WrapResult)

	var exprString string
	switch property.Name {
	case "length":
		if retExprPrefix != "" {
			retExpSuffix = ")"
		}
		exprString = fmt.Sprintf(`js.Value(%s).Length()`, receiver.Name)
	default:
		exprString = fmt.Sprintf(`%sjs.Value(%s).Call(%q)%s`, retExprPrefix, receiver.Name, property.Name, retExpSuffix)
	}
	expr, err := parser.ParseExpr(exprString)
	if err != nil {
		panic(err)
	}

	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{expr},
			},
		},
	}
	return &ast.FuncDecl{
		Name: property.ExportedIdentifier(),
		Recv: &ast.FieldList{List: []*ast.Field{
			{
				Names: []*ast.Ident{receiver},
				Type:  receiverIdentifier,
			},
		}},
		Type: property.funcType(),
		Body: body,
	}
}

func wrapResult(typeName, wrapResult string) (string, string) {
	typeMethodPrefix, typeMethodSuffix := "", ""

	if ast.IsExported(typeName) {
		typeMethodSuffix += ")"
		typeMethodPrefix = "wrap" + typeName + "(" + typeMethodPrefix
	} else {
		switch typeName {
		case "int":
			typeMethodSuffix = ".Int()"
		case "string":
			typeMethodSuffix = ".String()"
		case "bool":
			typeMethodSuffix = ".Bool()"
		case "float64":
			typeMethodSuffix = ".Float()"
		}
	}
	if wrapResult != "" {
		typeMethodSuffix += ")"
		typeMethodPrefix = wrapResult + "(" + typeMethodPrefix
	}
	return typeMethodPrefix, typeMethodSuffix
}

func (method Method) functionDeclaration(receiverIdentifier *ast.Ident) *ast.FuncDecl {
	receiver := ast.NewIdent("val")

	retExprPrefix, retExpSuffix := wrapResult(method.Result.Type, method.Result.Wrap)

	params := ""
	for _, param := range method.Parameters {
		params += ", " + param.ID.Name
	}

	expr, err := parser.ParseExpr(fmt.Sprintf(`%sjs.Value(val).Call(%q%s)%s`, retExprPrefix, method.Name, params, retExpSuffix))
	if err != nil {
		panic(err)
	}

	body := &ast.BlockStmt{
		List: []ast.Stmt{},
	}

	if method.Result.Type != "" {
		body.List = append(body.List, &ast.ReturnStmt{
			Results: []ast.Expr{expr},
		})
	} else {
		body.List = append(body.List, &ast.ExprStmt{X: expr})
	}
	return &ast.FuncDecl{
		Name: method.ExportedIdentifier(),
		Recv: &ast.FieldList{List: []*ast.Field{
			{
				Names: []*ast.Ident{receiver},
				Type:  receiverIdentifier,
			},
		}},
		Type: method.funcType(),
		Body: body,
	}
}

func (property Property) funcType() *ast.FuncType {
	funcType := new(ast.FuncType)
	funcType.Params = new(ast.FieldList)

	funcType.Results = &ast.FieldList{
		List: []*ast.Field{{
			Type: ast.NewIdent(property.Type),
		}},
	}

	return funcType
}

func (method Method) funcType() *ast.FuncType {
	funcType := new(ast.FuncType)
	funcType.Params = new(ast.FieldList)

	for _, param := range method.Parameters {
		funcType.Params.List = append(funcType.Params.List, &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(param.ID.Name)},
			Type:  ast.NewIdent(param.Type),
		})
	}
	funcType.Results = &ast.FieldList{
		List: []*ast.Field{{
			Type: ast.NewIdent(method.Result.Type),
		}},
	}
	return funcType
}

func (property Property) interfaceMethodSignature() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{property.ExportedIdentifier()},
		Type:  property.funcType(),
	}
}

func (method Method) interfaceMethodSignature() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{method.ExportedIdentifier()},
		Type:  method.funcType(),
	}
}

func (spec Interface) interfaceDeclaration() *ast.GenDecl {
	interfaceType := new(ast.InterfaceType)
	interfaceType.Methods = new(ast.FieldList)

	for _, property := range spec.Properties {
		interfaceType.Methods.List = append(interfaceType.Methods.List, property.interfaceMethodSignature())
	}

	for _, method := range spec.Methods {
		interfaceType.Methods.List = append(interfaceType.Methods.List, method.interfaceMethodSignature())
	}

	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: spec.ID().ExportedIdentifier(),
				Type: interfaceType,
			},
		},
	}
}

func (spec Class) classDeclarations() []ast.Decl {
	var declarations []ast.Decl

	classDeclaration := new(ast.TypeSpec)
	classDeclaration.Name = spec.ID().ExportedIdentifier()
	classDeclaration.Type = &ast.SelectorExpr{
		X:   ast.NewIdent("js"),
		Sel: ast.NewIdent("Value"),
	}
	declarations = append(declarations, &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{classDeclaration},
	})

	wrapDeclaration := &ast.FuncDecl{
		Name: ast.NewIdent("wrap" + classDeclaration.Name.Name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{{
					Names: []*ast.Ident{ast.NewIdent("value")},
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent("js"),
						Sel: ast.NewIdent("Value"),
					},
				}},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{{
					Type: spec.ExportedIdentifier(),
				}},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.CallExpr{
							Fun: spec.ExportedIdentifier(),
							Args: []ast.Expr{
								ast.NewIdent("value"),
							},
						},
					},
				},
			},
		},
	}
	declarations = append(declarations, wrapDeclaration)

	for _, property := range spec.Properties {
		declarations = append(declarations, property.functionDeclaration(classDeclaration.Name))
	}
	for _, method := range spec.Methods {
		declarations = append(declarations, method.functionDeclaration(classDeclaration.Name))
	}

	return declarations
}

type SpecificationList []Specification

func (list SpecificationList) Find(name string) (Specification, bool) {
	for _, spec := range list {
		if spec.ID().Name == name {
			return spec, true
		}
	}
	return nil, false
}

func generate(list SpecificationList) (ast.Node, error) {
	file := new(ast.File)

	//file.Doc = &ast.CommentGroup{
	//	List: []*ast.Comment{{Text: "GENERATED CODE. DO NOT EDIT"}},
	//}

	file.Name = ast.NewIdent("window")

	file.Decls = append(file.Decls, &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"syscall/js"`}},
			&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"time"`}},
		},
	})

	for _, s := range list {
		switch spec := s.(type) {
		case *Class:
			file.Decls = append(file.Decls, spec.classDeclarations()...)
		case *Interface:
			file.Decls = append(file.Decls, spec.interfaceDeclaration())
		}
	}

	return file, nil
}

func (list SpecificationList) resolveInheritance() error {
	for i, s := range list {
		resolved, err := resolveInheritance(s, list)
		if err != nil {
			return err
		}
		list[i] = resolved
	}
	return nil
}

func resolveInheritance(s Specification, list SpecificationList) (Specification, error) {
	switch spec := s.(type) {
	case *Class:
		methods, properties, err := resolveMethodsAndProperties(spec.Parent, list)
		if err != nil {
			return nil, err
		}
		spec.Methods = joinNew(methods, spec.Methods)
		spec.Properties = joinNew(properties, spec.Properties)
		return spec, nil
	case *Interface:
		methods, properties, err := resolveMethodsAndProperties(spec.Parent, list)
		if err != nil {
			return nil, err
		}
		spec.Methods = joinNew(methods, spec.Methods)
		spec.Properties = joinNew(properties, spec.Properties)
		return spec, nil
	default:
		panic("not implemented")
	}
}

func resolveMethodsAndProperties(parentName string, list SpecificationList) ([]Method, []Property, error) {
	if parentName == "" {
		return nil, nil, nil
	}
	p, found := list.Find(parentName)
	if !found {
		return nil, nil, fmt.Errorf("failed to find %q", parentName)
	}
	p, err := resolveInheritance(p, list)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to resolve inheritance for %q: %w", parentName, err)
	}
	switch parent := p.(type) {
	case *Class:
		return parent.Methods, parent.Properties, nil
	case *Interface:
		return parent.Methods, parent.Properties, nil
	default:
		panic("not implemented")
	}
}

func joinNew[T interface {
	ExportedIdentifier() *ast.Ident
}](lists ...[]T) []T {
	length := 0
	for _, list := range lists {
		length += len(list)
	}
	joined := make([]T, 0, length)
	for _, list := range lists {
	nextDecl:
		for _, declaration := range list {
			for _, existingDecl := range joined {
				if existingDecl.ExportedIdentifier().Name == declaration.ExportedIdentifier().Name {
					continue nextDecl
				}
			}
			joined = append(joined, declaration)
		}

	}
	return joined
}

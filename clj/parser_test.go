package clj

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParseAST(t *testing.T) {
	goSrc := `package main

func Square(x int) int { return x * x }`

	f, err := parse(goSrc)
	if err != nil {
		t.Fatal("unexpectd error: ", err)
	}

	cljast := Parse(f.Decls[0])

	if cljast.Text() != "(defn Square [x] (* x x))" {
		t.Errorf(`parse ast:
expect

       (defn Square [x] (* x x))

but got

       %s`, cljast.Text())
	}
}

func parse(src string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "src.go", src, 0)
}

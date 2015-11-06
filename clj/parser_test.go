package clj

import (
	"strings"
	"testing"
)

func TestParseFile(t *testing.T) {
	goSrc := `package main

func Square(x int) int { return x * x }`

	cljFile, err := ParseFile(strings.NewReader(goSrc))
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	if cljFile.Text() != "(defn Square [x] (* x x))" {
		t.Errorf(`parse ast:
expect

       (defn Square [x] (* x x))

but got

       %s`, cljFile.Text())
	}
}

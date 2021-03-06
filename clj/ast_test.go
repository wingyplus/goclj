package clj

import "testing"

func TestSymbol(t *testing.T) {
	if sym := Symbol("x").Text(); sym != "x" {
		t.Error("symbol: expect x but got %s", sym)
	}
}

func TestVector(t *testing.T) {
	cases := []struct {
		vector *Vector
		text   string
	}{
		{&Vector{1, 2, 3, 4}, "[1 2 3 4]"},
		{&Vector{1, 2, 3}, "[1 2 3]"},
		{&Vector{}, "[]"},
		{&Vector{Symbol("x")}, "[x]"},
	}

	for _, tc := range cases {
		if text := tc.vector.Text(); text != tc.text {
			t.Errorf("test vector: expect %s but got %s", tc.text, text)
		}
	}
}

func TestMacro(t *testing.T) {
	m := &Macro{
		Name: &Ident{
			Name: "defn",
		},
		List: []Expr{
			Vector{Symbol("x")},
		},
	}

	cljSrc := `(defn [x])`
	if m.Text() != cljSrc {
		t.Error("macro: not match (defn [x]) return", m.Text())
	}
}

func TestIf(t *testing.T) {
	testExpr(t, If(NewBinaryExpr("=", "x", "y"), NewBinaryExpr("*", "x", "y"), nil), "(if (= x y) (* x y))")
	testExpr(t, If(NewBinaryExpr("=", "x", "y"), NewBinaryExpr("*", "x", "y"), NewBinaryExpr("+", "x", "y")), "(if (= x y) (* x y) (+ x y))")
}

func testExpr(t *testing.T, expr Expr, cljSrc string) {
	if text := expr.Text(); text != cljSrc {
		t.Errorf(`parse ast:
expect

       %s

but got

       %s`, cljSrc, text)
	}
}

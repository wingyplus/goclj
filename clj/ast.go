package clj

import (
	"fmt"
	"strings"
)

type Expr interface {
	Text() string
}

type Symbol string

func (s Symbol) Text() string {
	return string(s)
}

type Vector []interface{}

func (vect Vector) Text() string {
	vectstr := make([]string, len(vect))
	for i, elem := range vect {
		switch elem.(type) {
		case int:
			vectstr[i] = fmt.Sprintf("%d", elem)
		case Symbol:
			vectstr[i] = elem.(Symbol).Text()
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(vectstr, " "))
}

type Ident struct {
	Name string
}

func (ident *Ident) Text() string {
	return ident.Name
}

type Macro struct {
	Name Expr
	List []Expr
}

func (m *Macro) Text() string {
	return fmt.Sprintf("(%s %s)", m.Name.Text(), join(m.List))
}

func join(list []Expr) string {
	l := make([]string, len(list))
	for i, expr := range list {
		l[i] = expr.Text()
	}
	return strings.Join(l, " ")
}

type BinaryExpr struct {
	Macro
}

func NewBinaryExpr(op, x, y string) *BinaryExpr {
	return &BinaryExpr{
		Macro: Macro{
			Name: &Ident{
				Name: op,
			},
			List: []Expr{
				Symbol(x), Symbol(y),
			},
		},
	}
}

type Func struct {
	Macro
}

func NewFunc(name string) *Func {
	return &Func{
		Macro: Macro{
			Name: &Ident{
				Name: "defn",
			},
			List: []Expr{
				Symbol(name),
			},
		},
	}
}

func (fn *Func) SetParams(params Vector) {
	fn.List = append(fn.List, params)
}

func (fn *Func) SetBody(stmt []Expr) {
	fn.List = append(fn.List, stmt...)
}

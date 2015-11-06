package clj

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func ParseExpr(decl ast.Decl) Expr {
	fn := decl.(*ast.FuncDecl)

	paramsList := fn.Type.Params.List
	params := make(Vector, len(paramsList))
	for i, param := range paramsList {
		params[i] = Symbol(param.Names[0].Name)
	}

	biexpr := fn.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.BinaryExpr)

	return &Macro{
		Name: &Ident{
			Name: "defn",
		},
		List: []Expr{
			Symbol(fn.Name.Name),
			params,
			&Macro{
				Name: &Ident{
					Name: biexpr.Op.String(),
				},
				List: []Expr{
					Symbol(biexpr.X.(*ast.Ident).Name),
					Symbol(biexpr.Y.(*ast.Ident).Name),
				},
			},
		},
	}
}

type File struct {
	Exprs []Expr
}

func (f *File) Text() string {
	return f.Exprs[0].Text()
}

func ParseFile(src io.Reader) (*File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), "src.go", src, 0)
	if err != nil {
		return nil, err
	}

	file := &File{
		Exprs: []Expr{
			ParseExpr(f.Decls[0]),
		},
	}
	return file, nil
}

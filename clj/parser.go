package clj

import "go/ast"

func Parse(decl ast.Decl) Expr {
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

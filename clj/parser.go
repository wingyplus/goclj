package clj

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func ParseExpr(decl ast.Decl) Expr {
	fnDecl := decl.(*ast.FuncDecl)

	paramsList := fnDecl.Type.Params.List
	params := make(Vector, len(paramsList))
	for i, param := range paramsList {
		params[i] = Symbol(param.Names[0].Name)
	}

	expr := fnDecl.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.BinaryExpr)

	fn := NewFunc(fnDecl.Name.Name)
	fn.SetParams(params)
	fn.SetBody([]Expr{parseBinaryExpr(expr)})

	return fn
}

func parseBinaryExpr(expr *ast.BinaryExpr) *BinaryExpr {
	return NewBinaryExpr(expr.Op.String(), expr.X.(*ast.Ident).Name, expr.Y.(*ast.Ident).Name)
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

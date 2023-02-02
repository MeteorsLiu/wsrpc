package builder

import (
	"go/ast"
	"log"
)

func checkExport(n string) {
	if ast.IsExported(n) {
		return
	}
	log.Fatalf("Warning: %s is Not Exported, cannot generate a non-exported function", n)
}

func parseFuncParamsType(p []*ast.Field) []string {
	if p == nil {
		return nil
	}
	tys := make([]string, len(p))

	for _, ty := range p {
		switch t := ty.Type.(type) {
		case *ast.StarExpr:
			tys = append(tys, "*"+t.X.(*ast.Ident).Name)
		case *ast.Ident:
			tys = append(tys, t.Name)

		}
	}

	return tys
}

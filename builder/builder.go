package builder

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type AST struct {
	filename     string
	exportStruct string
}

func NewAST(f, export string) *AST {
	return &AST{
		filename:     f,
		exportStruct: export,
	}
}

func (a *AST) isExportMember(da *ast.Field) bool {
	if da != nil {
		if das, ok := da.Type.(*ast.StarExpr); ok {
			if dasx, ok := das.X.(*ast.Ident); ok {
				if dasx.Name == a.exportStruct {
					return true
				}
			}
		} else if dai, ok := da.Type.(*ast.Ident); ok {
			fmt.Println(dai.Name)
		}
	}
	return false
}

func parseFuncParams(fd *ast.FuncType) {

}

func (a *AST) Parse() {
	fset := token.NewFileSet()
	fs, err := parser.ParseFile(fset, a.filename, nil, 0)
	if err != nil {
		log.Fatal("Parse Error: ", err)
	}
	for _, decl := range fs.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			for _, da := range d.Recv.List {
				if a.isExportMember(da) {
					fmt.Println("is Export Method: ", d.Name, d.Type)
				}
			}
		}
	}
}

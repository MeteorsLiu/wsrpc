package builder

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"strings"
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
func (a *AST) Filter(s string) bool {
	return strings.Contains(s, a.exportStruct)
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
				fmt.Println(reflect.TypeOf(da.Type).Elem())
			}
		}
	}
}

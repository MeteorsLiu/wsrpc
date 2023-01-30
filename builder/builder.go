package builder

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
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
	for _, v := range fs.Decls {
		fmt.Println(v)
	}
}

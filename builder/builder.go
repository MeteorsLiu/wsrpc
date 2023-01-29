package builder

import (
	"flag"
	"go/ast"
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
	fs, err := parser.ParseFile(fset, a.filename, "", 0)
	if err != nil {
		log.Fatal("Syntax Error")
	}
	ast.FilterFile(fs, a.Filter)
	ast.Print(fset, fs)
}
func main() {
	var parseFile string
	var exportSuffix string
	var exportStruct string
	flag.StringVar(&parseFile, "file", "", "RPC Struct Golang File")
	flag.StringVar(&exportSuffix, "suffix", "_gen.go", "The Suffix Name of Exported File")
	flag.StringVar(&exportStruct, "struct", "", "The RPC Struct to be exported.")
	flag.Parse()

	if parseFile == "" || exportStruct == "" {
		log.Fatal("No specific file")
	}
	NewAST(parseFile, exportStruct).Parse()
}

package builder

import (
	"flag"
	"log"
)

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

package builder

import (
	"testing"
)

func TestBuider(t *testing.T) {
	NewAST("testparse.go", "ExportStruct").Parse()
}

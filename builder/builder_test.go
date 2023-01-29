package builder

import (
	"testing"
)

func TestBuider(t *testing.T) {
	NewAST("test_parse.go", "TestExport").Parse()
}

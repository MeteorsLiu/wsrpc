package builder

import (
	"testing"
)

func TestBuider(t *testing.T) {
	NewAST("_test_parse.go", "TestExport").Parse()
}

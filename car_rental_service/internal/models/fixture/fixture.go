package fixture

import (
	"os"
	"path/filepath"
	"testing"
)

func LoadFixture(t *testing.T, filename string) []byte {
	data, err := os.ReadFile(filepath.Join("../testdata", filename))
	if err != nil {
		t.Fatalf("can't read test file %v", filename)
	}
	return data
}
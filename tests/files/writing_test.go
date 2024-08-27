package files_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vossenwout/ai-code-review/internal/files"
)

// Tests the functionality to save the project string to a file.
func TestSaveStringToFile(t *testing.T) {
	content := "This is an example project."
	tempFile := filepath.Join(t.TempDir(), "testfdile.txt")

	err := files.SaveStringToFile(content, tempFile)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	savedContent, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	if string(savedContent) != content {
		t.Errorf("expected content %s, got %s", content, string(savedContent))
	}
}

package files_test

import (
	"testing"

	"github.com/vossenwout/ai-code-review/internal/files"
)

// Test the filtering of file paths.
func TestFilterFilePaths(t *testing.T) {
	tests := []struct {
		name             string
		filePaths        []string
		prefixesToFilter []string
		expected         []string
	}{
		{
			name:             "No filter match",
			filePaths:        []string{"main.go", "internal/files/reading.go", "internal/formatting/format.go"},
			prefixesToFilter: []string{"test"},
			expected:         []string{"main.go", "internal/files/reading.go", "internal/formatting/format.go"},
		},
		{
			name:             "Single filter match",
			filePaths:        []string{"main.go", "internal/files/reading.go", "internal/formatting/format.go"},
			prefixesToFilter: []string{"internal"},
			expected:         []string{"main.go"},
		},
		{
			name: "Multiple filter matches",
			filePaths: []string{"main.go", "internal/files/reading.go",
				"cmd/ai-code-review/main.go", ".gitignore"},
			prefixesToFilter: []string{"internal", "cmd", "."},
			expected:         []string{"main.go"},
		},
		{
			name:             "No files left after filtering",
			filePaths:        []string{"internal/files/reading.go", "cmd/ai-code-review/main.go"},
			prefixesToFilter: []string{"internal", "cmd"},
			expected:         []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := files.FilterFilePaths(tt.filePaths, tt.prefixesToFilter)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d files, got %d", len(tt.expected), len(result))
			}
			for i, path := range result {
				if path != tt.expected[i] {
					t.Errorf("expected %s, got %s", tt.expected[i], path)
				}
			}
		})
	}
}

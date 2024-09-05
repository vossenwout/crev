package formatting_test

import (
	"os"
	"strings"
	"testing"

	"github.com/vossenwout/crev/internal/formatting"
)

func TestGeneratePathTree(t *testing.T) {
	paths := []string{
		"cmd",
		"cmd/ai-code-review",
		"cmd/ai-code-review/main.go",
		"internal",
		"internal/files",
		"internal/files/filtering.go",
		"internal/formatting",
		"internal/formatting/format.go",
		"go.mod",
	}
	expected, err := os.ReadFile("../test_data/expected_tree_1.txt")
	if err != nil {
		t.Errorf("error reading test file: %v", err)
	}

	result := formatting.GeneratePathTree(paths)
	// Normalize both expected and result strings
	expectedStr := strings.TrimSpace(string(expected))
	resultStr := strings.TrimSpace(result)

	if resultStr != expectedStr {
		t.Errorf("expected \n%s\n, got \n%s\n", expectedStr, resultStr)
	}
}

func TestCreateProjectString(t *testing.T) {
	projectTree, err := os.ReadFile("../test_data/expected_tree_1.txt")
	if err != nil {
		t.Errorf("error reading test file: %v", err)
	}
	fileContentMap := map[string]string{
		"cmd/ai-code-review/main.go":    "package main\n",
		"internal/files/filtering.go":   "package files\n",
		"internal/formatting/format.go": "package formatting\n",
		"go.mod":                        "go mod\n",
	}
	expectedProjectString, err := os.ReadFile("../test_data/expected_project_string_1.txt")
	if err != nil {
		t.Errorf("error reading test file: %v", err)
	}
	expected := string(expectedProjectString)
	result := formatting.CreateProjectString(string(projectTree), fileContentMap)
	expectedStr := strings.TrimSpace(string(expected))
	resultStr := strings.TrimSpace(result)
	if resultStr != expectedStr {
		t.Errorf("expected \n%s\n, got \n%s\n", expected, result)
	}
}

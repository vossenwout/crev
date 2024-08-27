package files_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vossenwout/ai-code-review/internal/files"
)

// Tests the functionality to get all file paths starting from a root path.
func TestGetAllFilePaths(t *testing.T) {
	rootDir := t.TempDir()

	subDir := filepath.Join(rootDir, "subdir")
	os.Mkdir(subDir, 0755)
	os.WriteFile(filepath.Join(rootDir, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(subDir, "file2.txt"), []byte("content2"), 0644)

	expected := []string{
		filepath.Join(rootDir, "file1.txt"),
		subDir,
		filepath.Join(subDir, "file2.txt"),
	}

	filePaths, err := files.GetAllFilePaths(rootDir)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(filePaths) != len(expected) {
		t.Fatalf("expected %d files, got %d", len(expected), len(filePaths))
	}

	for _, exp := range expected {
		found := false
		for _, fp := range filePaths {
			if fp == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected path %s not found in result", exp)
		}
	}
}

// Tests the functionality to read all the content of files and directories and create a
// file content map.
func TestGetContentMapOfFiles(t *testing.T) {
	rootDir := t.TempDir()

	subDir1 := filepath.Join(rootDir, "subdir_1")
	subDir2 := filepath.Join(rootDir, "subdir_2")
	os.Mkdir(subDir1, 0755)
	os.Mkdir(subDir2, 0755)
	os.WriteFile(filepath.Join(rootDir, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(subDir1, "file2.txt"), []byte("content2"), 0644)

	filePaths := []string{
		filepath.Join(rootDir, "file1.txt"),
		subDir1,
		filepath.Join(subDir1, "file2.txt"),
		subDir2,
	}

	fileContentMap, err := files.GetContentMapOfFiles(filePaths)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if fileContentMap[filepath.Join(rootDir, "file1.txt")] != "content1" {
		t.Errorf("expected content1, got %s", fileContentMap[filepath.Join(rootDir, "file1.txt")])
	}

	if fileContentMap[filepath.Join(subDir1, "file2.txt")] != "content2" {
		t.Errorf("expected content2, got %s", fileContentMap[filepath.Join(subDir1, "file2.txt")])
	}

	// subDir1 is a directory with at least 1 file, so it should not be present in the map
	if _, ok := fileContentMap[subDir1]; ok {
		t.Errorf("directory with at least 1 file should not be present %s", fileContentMap[subDir1])
	}

	// subDir2 is an empty directory, so it should be present in the map
	if fileContentMap[subDir2] != "empty directory" {
		t.Errorf("expected empty directory, got %s", fileContentMap[subDir2])
	}
}

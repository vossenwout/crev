package files_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vossenwout/crev/internal/files"
)

// Tests the functionality to get all file paths starting from a root path.
func TestGetAllFilePaths(t *testing.T) {
	rootDir := t.TempDir()

	subDir := filepath.Join(rootDir, "subdir")
	err := os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(rootDir, "file1.txt"), []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir, "file2.txt"), []byte("content2"), 0644)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []string{
		filepath.Join(rootDir, "file1.txt"),
		subDir,
		filepath.Join(subDir, "file2.txt"),
	}

	filePaths, err := files.GetAllFilePaths(rootDir, nil, nil)
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

// tests the prefix filtering functionality with a full path prefix.
func TestGetAllFilePathsWithPrefixFilter(t *testing.T) {
	rootDir := t.TempDir()

	subDir1 := filepath.Join(rootDir, "subdir_1")
	err := os.Mkdir(subDir1, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	subDir2 := filepath.Join(rootDir, "subdir_2")
	err = os.Mkdir(subDir2, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(rootDir, "file1.go"), []byte("content1"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir1, "file2.go"), []byte("content2"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir2, "file3.go"), []byte("content3"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []string{
		filepath.Join(rootDir, "file1.go"),
		subDir2,
		filepath.Join(subDir2, "file3.go"),
	}
	// filter out full path prefix subDir1
	filePaths, err := files.GetAllFilePaths(rootDir, []string{subDir1}, nil)

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

// tests the prefix filtering functionality when the prefix is a nested directory.
func TestGetAllFilePathsWithPrefixFilterNestedDir(t *testing.T) {
	rootDir := t.TempDir()

	subDir1 := filepath.Join(rootDir, "subdir_1")
	err := os.Mkdir(subDir1, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	subDir2 := filepath.Join(subDir1, ".subdir_2")
	err = os.Mkdir(subDir2, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir2, "file1.go"), []byte("content3"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	subDir3 := filepath.Join(rootDir, "subdir_3")
	err = os.Mkdir(subDir3, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir3, "file2.go"), []byte("content3"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []string{
		subDir1,
		subDir3,
		filepath.Join(subDir3, "file2.go"),
	}

	filePaths, err := files.GetAllFilePaths(rootDir, []string{"."}, nil)

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

// Tests the functionality to filter out text extensions.
func TestGetAllFilePathsWithExtensionFilter(t *testing.T) {
	rootDir := t.TempDir()

	subDir1 := filepath.Join(rootDir, "subdir_1")
	err := os.Mkdir(subDir1, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	subDir2 := filepath.Join(rootDir, "subdir_2")
	err = os.Mkdir(subDir2, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(rootDir, "file1.go"), []byte("content1"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir1, "file2.go"), []byte("content2"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir2, "file3.txt"), []byte("content3"), 064)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []string{
		filepath.Join(rootDir, "file1.go"),
		subDir1,
		filepath.Join(subDir1, "file2.go"),
		subDir2,
	}

	filePaths, err := files.GetAllFilePaths(rootDir, nil, []string{".go"})
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
	err := os.Mkdir(subDir1, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.Mkdir(subDir2, 0755)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(rootDir, "file1.txt"), []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = os.WriteFile(filepath.Join(subDir1, "file2.txt"), []byte("content2"), 0644)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	filePaths := []string{
		filepath.Join(rootDir, "file1.txt"),
		subDir1,
		filepath.Join(subDir1, "file2.txt"),
		subDir2,
	}

	fileContentMap, err := files.GetContentMapOfFiles(filePaths, 10)
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

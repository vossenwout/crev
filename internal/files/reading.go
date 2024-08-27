// Contains code to read the content of files and directories.
package files

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Given a file path, getFileContent returns the content of the file.
func getFileContent(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

// Given a root path returns all the file paths in the root directory
// and its subdirectories.
func GetAllFilePaths(root string) ([]string, error) {
	var filePaths []string
	err := filepath.WalkDir(root, func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path != root {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}

// Given a list of file paths, getContentMapOfFiles returns a map of file paths to their content.
func GetContentMapOfFiles(filePaths []string) (map[string]string, error) {
	fileContentMap := make(map[string]string)
	for _, path := range filePaths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			fileContent, err := getFileContent(path)
			if err != nil {
				return nil, err
			}
			fileContentMap[path] = fileContent
		} else {
			dirEntries, err := os.ReadDir(path)
			if err != nil {
				return nil, err
			}
			if len(dirEntries) == 0 {
				fileContentMap[path] = "empty directory"
			}
		}
	}
	return fileContentMap, nil
}

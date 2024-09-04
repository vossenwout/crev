// Contains code to read the content of files and directories.
package files

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Given a root path returns all the file paths in the root directory
// and its subdirectories.
func GetAllFilePaths(root string, prefixesToFilter []string, extensionsToKeep []string) ([]string, error) {
	var filePaths []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Skip the root directory.
		if path == root {
			return nil
		}
		// First filter out the paths that contain any of the prefixes in prefixesToFilter.
		for _, prefixToFilter := range prefixesToFilter {
			// check if the full path itself contains the prefix
			if strings.HasPrefix(path, prefixToFilter) {
				return nil
			}
			// check if any of the components of the path contain the prefix
			components := strings.Split(filepath.Clean(path), string(filepath.Separator))
			for _, component := range components {
				if strings.HasPrefix(component, prefixToFilter) {
					return nil
				}
			}
		}

		// if no specific extensions to include are provided, include all files
		if len(extensionsToKeep) == 0 {
			filePaths = append(filePaths, path)
		} else {
			// if path is a directory, include it
			if d.IsDir() {
				filePaths = append(filePaths, path)
			} else {
				// else, only include files with the specified extensions
				for _, extensionToKeep := range extensionsToKeep {
					if filepath.Ext(path) == extensionToKeep {
						filePaths = append(filePaths, path)
						break
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

// Given a file path, getFileContent returns the content of the file.
func getFileContent(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

// Given a list of file paths, getContentMapOfFiles returns a map of file paths to their content.
func GetContentMapOfFiles(filePaths []string, maxConcurrency int) (map[string]string, error) {
	var fileContentMap sync.Map
	var wg sync.WaitGroup
	errChan := make(chan error, len(filePaths))
	semaphore := make(chan struct{}, maxConcurrency)

	for _, path := range filePaths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			info, err := os.Stat(p)
			if err != nil {
				errChan <- err
				return
			}
			if !info.IsDir() {
				fileContent, err := getFileContent(p)
				if err != nil {
					errChan <- err
					return
				}
				fileContentMap.Store(p, fileContent)
			} else {
				dirEntries, err := os.ReadDir(p)
				if err != nil {
					errChan <- err
					return
				}
				if len(dirEntries) == 0 {
					fileContentMap.Store(p, "empty directory")
				}
			}
		}(path)
	}
	wg.Wait()
	close(errChan)
	if len(errChan) > 0 {
		return nil, <-errChan
	}

	resultMap := make(map[string]string)
	fileContentMap.Range(func(key, value interface{}) bool {
		resultMap[key.(string)] = value.(string)
		return true
	})

	return resultMap, nil
}

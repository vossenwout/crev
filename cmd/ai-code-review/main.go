// Contains code to flatten out your directory structure into a single file so that you can
// give it to an llm for code review.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Given a path, getAllFilePaths returns all the file paths in the directory
// and its subdirectories.
func getAllFilePaths(path string) ([]string, error) {
	var filePaths []string
	err := filepath.WalkDir(path, func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path != "." {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}

// Given a list of paths, generatePathTree returns a string representation of the
// directory structure.
func generatePathTree(paths []string) string {
	var treeString strings.Builder
	levelPrefix := make(map[int]string)
	for i, path := range paths {
		parts := strings.Split(path, string(os.PathSeparator))
		level := len(parts) - 1
		isLast := i == len(paths)-1 ||
			len(strings.Split(paths[i+1], string(os.PathSeparator))) <= level
		var prefix string
		for l := 0; l < level; l++ {
			prefix += levelPrefix[l]
		}
		branch := "├── "
		if isLast {
			branch = "└── "
			levelPrefix[level] = "    "
		} else {
			levelPrefix[level] = "│   "
		}
		treeString.WriteString(prefix + branch + parts[len(parts)-1] + "\n")
	}
	return treeString.String()
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
func getContentMapOfFiles(filePaths []string) (map[string]string, error) {
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

// Filter out file paths that contain any of the prefixes in prefixesToFilter.
func filterFilePaths(filePaths []string, prefixesToFilter []string) []string {
	var retainedFilePaths []string
	for _, filePath := range filePaths {
		isPrefixFound := false
		for _, prefixToFilter := range prefixesToFilter {
			if strings.HasPrefix(filePath, prefixToFilter) {
				isPrefixFound = true
				break
			}
		}
		if !isPrefixFound {
			retainedFilePaths = append(retainedFilePaths, filePath)
		}
	}
	return retainedFilePaths
}

// Creates a string representation of the project.
func createProjectString(projectTree string, fileContentMap map[string]string) string {
	var projectString strings.Builder
	projectString.WriteString("Project Directory Structure:" + "\n")
	projectString.WriteString(projectTree + "\n\n")
	for fileName := range fileContentMap {
		fileContent := fileContentMap[fileName]
		projectString.WriteString("File: " + "\n")
		projectString.WriteString(fileName + "\n")
		projectString.WriteString("Content: " + "\n")
		projectString.WriteString(fileContent + "\n\n")
	}
	return projectString.String()
}

// Saves a string to a file.
func saveStringToFile(content string, path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// https://trstringer.com/golang-deferred-function-error-handling/
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err == nil {
				err = fmt.Errorf("failed to close file: %w", closeErr)
			}
		}
	}()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	root := "."
	filePaths, err := getAllFilePaths(root)
	if err != nil {
		log.Fatal(err)
		return
	}
	patternsToFilter := []string{".", "readme"}
	filePaths = filterFilePaths(filePaths, patternsToFilter)
	projectTree := generatePathTree(filePaths)
	fileContentMap, err := getContentMapOfFiles(filePaths)
	if err != nil {
		log.Fatal(err)
	}
	projectString := createProjectString(projectTree, fileContentMap)
	err = saveStringToFile(projectString, ".project.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Project structure saved to .project.txt")

}

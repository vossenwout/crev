// Contains code to flatten out your directory structure into a single file so that you can
// give it to an llm for code review.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Also add some metadata at the top of the file maybe?

// getAllFilePaths recursively gets all file paths in a directory.
func getAllFilePaths(path string) ([]string, error) {
	var filePaths []string

	// Walk through the directory.
	err := filepath.Walk(path, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip the root directory itself.
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

// generatePathTree constructs a tree-like structure from file paths.
func generatePathTree(paths []string) string {
	var treeString strings.Builder
	levelPrefix := make(map[int]string) // To keep track of prefix at each level

	for i, path := range paths {
		parts := strings.Split(path, string(os.PathSeparator))
		level := len(parts) - 1

		// Determine if this is the last item at its level.
		isLast := i == len(paths)-1 || len(strings.Split(paths[i+1], string(os.PathSeparator))) <= level

		// Build the prefix based on the level and whether it’s the last item.
		var prefix string
		for l := 0; l < level; l++ {
			prefix += levelPrefix[l]
		}

		// Determine the correct branch.
		branch := "├── "
		if isLast {
			branch = "└── "
			levelPrefix[level] = "    " // No vertical line after the last item
		} else {
			levelPrefix[level] = "│   " // Continue the vertical line
		}

		// Append the current file or directory to the tree string.
		treeString.WriteString(prefix + branch + parts[len(parts)-1] + "\n")
	}
	return treeString.String()
}

func getFileContent(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func getContentMapOfFiles(filePaths []string) map[string]string {
	fileContentMap := make(map[string]string)
	for _, path := range filePaths {
		info, err := os.Stat(path)
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			fileContent, err := getFileContent(path)
			if err != nil {
				panic(err)
			}
			fileContentMap[path] = fileContent
		}
	}
	return fileContentMap
}

// TODO implement a better pattern filter
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

func createProjectString(projectTree string, fileContentMap map[string]string) string {
	var projectString strings.Builder

	projectString.WriteString(projectTree + "/n")

	for fileName := range fileContentMap {
		fileContent := fileContentMap[fileName]
		projectString.WriteString(fileName + "/n")
		projectString.WriteString(fileContent + "/n")
	}

	return projectString.String()

}

func main() {
	root := "." // Replace with your root directory

	// Get all file paths.
	filePaths, err := getAllFilePaths(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	patternsToFilter := []string{".", "readme"}
	filePaths = filterFilePaths(filePaths, patternsToFilter)

	// Build the tree structure from file paths.
	projectTree := generatePathTree(filePaths)
	fmt.Println(projectTree)
	fileContentMap := getContentMapOfFiles(filePaths)
	//fmt.Println(fileContentMap)
	for key := range fileContentMap {
		fmt.Println(key)
	}
	createProjectString(projectTree, fileContentMap)
}

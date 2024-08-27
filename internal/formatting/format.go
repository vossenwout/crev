// Contains code to format the project structure into a string.
package formatting

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Given a path, calculatePathLevel returns the level of the path in the directory structure.
func calculatePathLevel(path string) int {
	return len(strings.Split(path, string(os.PathSeparator))) - 1
}

// Given a list of paths, an index, and a level, checkIfLastPathAtLevel returns true if the path
func checkIfLastPathAtLevel(paths []string, i, level int) bool {
	return i == len(paths)-1 || len(strings.Split(paths[i+1], string(os.PathSeparator))) <= level
}

// Given a level prefix, level, and a boolean indicating if the path is the last at the level,
func buildTreeBranch(levelPrefix map[int]string, level int, isLast bool) string {
	var branchPrefix strings.Builder
	for l := 0; l < level; l++ {
		branchPrefix.WriteString(levelPrefix[l])
	}

	branch := "├── "
	if isLast {
		branch = "└── "
		levelPrefix[level] = "    "
	} else {
		levelPrefix[level] = "│   "
	}

	return branchPrefix.String() + branch
}

// Given a list of paths, generatePathTree returns a string representation of the
// directory structure.
func GeneratePathTree(paths []string) string {
	// Sort the paths lexicographically to ensure correct tree structure
	sort.Strings(paths)
	var treeBuilder strings.Builder
	levelPrefix := make(map[int]string)

	for i, path := range paths {
		level := calculatePathLevel(path)
		isLast := checkIfLastPathAtLevel(paths, i, level)
		treeBuilder.WriteString(buildTreeBranch(levelPrefix, level, isLast) + filepath.Base(path) + "\n")
	}

	return treeBuilder.String()
}

// Creates a string representation of the project.
func CreateProjectString(projectTree string, fileContentMap map[string]string) string {
	var projectString strings.Builder
	projectString.WriteString("Project Directory Structure:" + "\n")
	projectString.WriteString(projectTree + "\n\n")

	// Collect and sort the file paths lexicographically to make the function deterministic
	filePaths := make([]string, 0, len(fileContentMap))
	for filePath := range fileContentMap {
		filePaths = append(filePaths, filePath)
	}
	sort.Strings(filePaths)

	for _, fileName := range filePaths {
		fileContent := fileContentMap[fileName]
		projectString.WriteString("File: " + "\n")
		projectString.WriteString(fileName + "\n")
		projectString.WriteString("Content: " + "\n")
		projectString.WriteString(fileContent + "\n\n")
	}
	return projectString.String()
}

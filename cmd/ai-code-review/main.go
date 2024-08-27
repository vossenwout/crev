// Contains code to flatten out your directory structure into a single file so that you can
// give it to an llm for code review.
package main

import (
	"log"

	"github.com/vossenwout/ai-code-review/internal/files"
	"github.com/vossenwout/ai-code-review/internal/formatting"
)

func main() {
	root := "."
	filePaths, err := files.GetAllFilePaths(root)
	if err != nil {
		log.Fatal(err)
		return
	}
	patternsToFilter := []string{".", "readme"}
	filePaths = files.FilterFilePaths(filePaths, patternsToFilter)

	projectTree := formatting.GeneratePathTree(filePaths)
	fileContentMap, err := files.GetContentMapOfFiles(filePaths)
	if err != nil {
		log.Fatal(err)
	}
	projectString := formatting.CreateProjectString(projectTree, fileContentMap)
	err = files.SaveStringToFile(projectString, ".project.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Project structure saved to .project.txt")

}

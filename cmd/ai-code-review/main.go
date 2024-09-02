// Contains code to flatten out your directory structure into a single file so that you can
// give it to an llm for code review.
package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/vossenwout/ai-code-review/internal/files"
	"github.com/vossenwout/ai-code-review/internal/formatting"
	"github.com/vossenwout/ai-code-review/internal/logging"
)

func parseCommandlineArgs() (string, int, []string) {
	// Parse command line arguments
	patternsPtr := flag.String("ignore", ".", "Comma separated list of perfixes of file and directory names to ingore. Ex. .,tests,readme")
	concurrencyPtr := flag.Int("concurrency", 1000, "Maximum concurrent file reads")
	flag.Parse()
	if len(flag.Args()) < 1 {
		logging.DisplayUsageAndExit()
	}
	root := flag.Arg(0)
	maxConcurrency := *concurrencyPtr
	patternsToFilter := strings.Split(*patternsPtr, ",")
	for i, pattern := range patternsToFilter {
		patternsToFilter[i] = strings.TrimSpace(pattern)
	}
	return root, maxConcurrency, patternsToFilter
}

func main() {
	start := time.Now()
	root, maxConcurrency, patternsToFilter := parseCommandlineArgs()
	//dummy change
	filePaths, err := files.GetAllFilePaths(root)
	if err != nil {
		log.Fatal(err)
		return
	}

	filePaths = files.FilterFilePaths(filePaths, patternsToFilter)
	projectTree := formatting.GeneratePathTree(filePaths)
	fileContentMap, err := files.GetContentMapOfFiles(filePaths, maxConcurrency)

	if err != nil {
		log.Fatal(err)
	}

	projectString := formatting.CreateProjectString(projectTree, fileContentMap)
	err = files.SaveStringToFile(projectString, ".project.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Project structure successfully saved to .project.txt")
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

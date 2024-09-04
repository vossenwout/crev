/*
Copyright © 2024 Wout Vossen <vossen.w@hotmail.com>
*/
package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/vossenwout/ai-code-review/internal/files"
	"github.com/vossenwout/ai-code-review/internal/formatting"
)

var ignoreDirs []string
var extensionsToKeep []string
var maxConcurrency int

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a textual representation of your code",
	Long: `Generates a textual representation of your code.

Example usage:
ai-code-review generate --path /path/to/code
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// start timer
		start := time.Now()

		// get all file paths from the root directory
		rootDir := "."
		filePaths, err := files.GetAllFilePaths(rootDir, ignoreDirs, extensionsToKeep)
		if err != nil {
			log.Fatal(err)
			return
		}

		// generate the project tree
		projectTree := formatting.GeneratePathTree(filePaths)

		// get the content of all files
		fileContentMap, err := files.GetContentMapOfFiles(filePaths, maxConcurrency)
		if err != nil {
			log.Fatal(err)
		}

		// create the project string
		projectString := formatting.CreateProjectString(projectTree, fileContentMap)

		// save the project string to a file
		err = files.SaveStringToFile(projectString, ".project.txt")
		if err != nil {
			log.Fatal(err)
		}

		// log success
		log.Println("Project structure successfully saved to .project.txt")
		elapsed := time.Since(start)
		log.Printf("Execution time: %s", elapsed)

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringSliceVar(&ignoreDirs, "ignore", []string{"."}, "Comma seperated prefixes of paths to ignore")
	generateCmd.Flags().StringSliceVar(&extensionsToKeep, "extensions", []string{}, "Comma seperated file extensions to keep. (default: all files)")
	generateCmd.Flags().IntVar(&maxConcurrency, "max-concurrency", 1000, "Maximum number of concurrent file reads")
}

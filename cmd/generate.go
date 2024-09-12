// Description: This file contains the generate command which generates a textual representation of the project structure.
package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vossenwout/crev/internal/files"
	"github.com/vossenwout/crev/internal/formatting"
)

var standardPrefixesToFilter = []string{
	// ignore .git, .idea, .vscode, etc.
	".",
	// ignore crev specific files
	"crev",
	// ignore go.mod, go.sum, etc.
	"go",
	"license",
	// poetry
	"pyproject.toml",
	"poetry.lock",
	"venv",
	// output files
	"build",
	"dist",
	"out",
	"target",
	"bin",
	// javascript
	"node_modules",
	"coverage",
	"public",
	"static",
	"Thumbs.db",
	"package",
	"yarn.lock",
	"package",
	"tsconfig",
	// next.js
	"next.config",
	"next-env",
	// python
	"__pycache__",
	"logs",
	// java
	"gradle",
	// c++
	"CMakeLists",
	// ruby
	"vendor",
	"Gemfile",
	// php
	"composer",
	// rust
	"target",
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a textual representation of your code",
	Long: `Generates a textual representation of your code starting from the directory you execute
	this command in. By default files starting with "." are ignored.

Example usage:
crev generate
crev generate --ignore=tests,readme.md --extensions=go,py,js
`,
	Args: cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		// start timer
		start := time.Now()

		// get all file paths from the root directory
		rootDir := "."
		prefixesToFilter := viper.GetStringSlice("ignore")
		// always ignore directories starting with "."
		prefixesToFilter = append(prefixesToFilter, standardPrefixesToFilter...)
		extensionsToKeep := viper.GetStringSlice("extensions")
		filePaths, err := files.GetAllFilePaths(rootDir, prefixesToFilter, extensionsToKeep)
		if err != nil {
			log.Fatal(err)
			return
		}

		// generate the project tree
		projectTree := formatting.GeneratePathTree(filePaths)

		maxConcurrency := 100
		// get the content of all files
		fileContentMap, err := files.GetContentMapOfFiles(filePaths, maxConcurrency)
		if err != nil {
			log.Fatal(err)
		}

		// create the project string
		projectString := formatting.CreateProjectString(projectTree, fileContentMap)

		outputFile := "crev-project.txt"
		// save the project string to a file
		err = files.SaveStringToFile(projectString, outputFile)
		if err != nil {
			log.Fatal(err)
		}

		// log success
		log.Println("Project overview succesfully saved to: " + outputFile)

		// estimate number of tokens
		log.Printf("Estimated token count: %d - %d tokens",
			len(projectString)/4, len(projectString)/3)

		elapsed := time.Since(start)
		log.Printf("Execution time: %s", elapsed)

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringSlice("ignore", []string{}, "Comma-separated prefixes of paths to ignore")
	generateCmd.Flags().StringSlice("extensions", []string{}, "Comma-separated file extensions to include. (default: all files)")
	err := viper.BindPFlag("ignore", generateCmd.Flags().Lookup("ignore"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("extensions", generateCmd.Flags().Lookup("extensions"))
	if err != nil {
		log.Fatal(err)
	}
}

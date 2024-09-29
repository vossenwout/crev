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

var standardPrefixesToIgnore = []string{
	// ignore .git, .idea, .vscode, etc.
	".",
	// ignore crev specific files
	"crev",
	// ignore go.mod, go.sum, etc.
	"go",
	"license",
	// readme
	"readme",
	"README",
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
	// tailwind
	"tailwind",
	"postcss",
}

var standardExtensionsToIgnore = []string{
	".jpeg",
	".jpg",
	".png",
	".gif",
	".pdf",
	".svg",
	".ico",
	".woff",
	".woff2",
	".eot",
	".ttf",
	".otf",
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Bundle your project into a single file",
	Long: `Bundle your project into a single file, starting from the directory you are in.
By default common configuration and setup files (ex. .vscode, .venv, package.lock) are ignored as well as non-text extensions like .jpeg, .png, .pdf. 

For more information see: https://crevcli.com/docs

Example usage:
crev bundle
crev bundle --ignore-pre=tests,readme --ignore-ext=.txt 
crev bundle --ignore-pre=tests,readme --include-ext=.go,.py,.js
`,
	Args: cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		// start timer
		start := time.Now()

		// get all file paths from the root directory
		rootDir := "."

		prefixesToIgnore := viper.GetStringSlice("ignore-pre")
		prefixesToIgnore = append(prefixesToIgnore, standardPrefixesToIgnore...)

		extensionsToIgnore := viper.GetStringSlice("ignore-ext")
		extensionsToIgnore = append(extensionsToIgnore, standardExtensionsToIgnore...)

		extensionsToInclude := viper.GetStringSlice("include-ext")

		filePaths, err := files.GetAllFilePaths(rootDir, prefixesToIgnore,
			extensionsToInclude, extensionsToIgnore)
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
	// TODO Fix description with defaults
	generateCmd.Flags().StringSlice("ignore-pre", []string{}, "Comma-separated prefixes of file and dir names to ignore. Ex tests,readme")
	generateCmd.Flags().StringSlice("ignore-ext", []string{}, "Comma-separated file extensions to ignore. Ex .txt,.md")
	generateCmd.Flags().StringSlice("include-ext", []string{}, "Comma-separated file extensions to include. Ex .go,.py,.js")
	err := viper.BindPFlag("ignore-pre", generateCmd.Flags().Lookup("ignore-pre"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("ignore-ext", generateCmd.Flags().Lookup("ignore-ext"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("include-ext", generateCmd.Flags().Lookup("include-ext"))
	if err != nil {
		log.Fatal(err)
	}
}

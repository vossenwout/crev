// Description: This file contains the generate command which generates
// a textual representation of the project structure.
package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vossenwout/crev/internal/files"
	"github.com/vossenwout/crev/internal/formatting"
)

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
	Run: func(cmd *cobra.Command, args []string) {
		// start timer
		start := time.Now()

		// get all file paths from the root directory
		rootDir := "."
		prefixesToFilter := viper.GetStringSlice("ignore")
		// always ignore directories starting with "."
		prefixesToFilter = append(prefixesToFilter, ".")
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

		output_file := ".crev-project-overview.txt"
		// save the project string to a file
		err = files.SaveStringToFile(projectString, output_file)
		if err != nil {
			log.Fatal(err)
		}

		// log success
		log.Println("Project structure successfully saved to " + output_file)
		elapsed := time.Since(start)
		log.Printf("Execution time: %s", elapsed)

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringSlice("ignore", []string{}, "Comma seperated prefixes of paths to ignore")
	generateCmd.Flags().StringSlice("extensions", []string{}, "Comma seperated file extensions to include. (default: all files)")
	viper.BindPFlag("ignore", generateCmd.Flags().Lookup("ignore"))
	viper.BindPFlag("extensions", generateCmd.Flags().Lookup("extensions"))
}

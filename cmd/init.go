// Description: This file contains the implementation of the init command which is used to create
// a default configuration file for the crev tool.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Define a default template configuration
var defaultConfig = []byte(`
# Configuration for the crev tool
# specify your Code AI Review API key (necessary for review command)
api-key: "TODO"
# specify the prefixes of files and directories to ignore (paths starting with . are always ignored)
ignore: # ex. [tests, build, readme.md]
# specify the extensions of files to include (by default all files are included)
extensions: # ex. [.go, .py, .js]
`)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a default configuration file",
	Long: `Generates a default configuration file (.crev-config.yaml) in the current directory.

The configuration file includes:
- API key for accessing the Code AI Review service (required for the "review" command)
- File and directory ignore patterns when generating the project overview
- File extensions to include when generating the project overview

You can modify this file as needed to suit your project's structure.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set the filename for the config
		configFileName := ".crev-config.yaml"

		// Get the current directory
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Unable to get current directory")
			os.Exit(1)
		}

		// Define the full path for the config file
		configFilePath := filepath.Join(currentDir, configFileName)

		// Check if config file already exists
		if _, err := os.Stat(configFilePath); err == nil {
			fmt.Println("Config file already exists at", configFilePath)
			os.Exit(1)
		}

		// Write the default config to the file
		err = os.WriteFile(configFilePath, defaultConfig, 0644)
		if err != nil {
			fmt.Println("Unable to write config file")
			os.Exit(1)
		}

		// Load the config using Viper
		viper.SetConfigFile(configFilePath)
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Unable to read config file")
			os.Exit(1)
		}

		fmt.Println("Config file created at", configFilePath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

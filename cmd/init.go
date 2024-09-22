// Description: This file implements the "init" command, which generates a default configuration file in the current directory.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Define a default template configuration
var defaultConfig = []byte(`# Configuration for the crev tool

# specify your CREV API key (necessary for review command) ! this overwrites the value you specify in the environment variable
crev_api_key: # ex. csk_8e796a8f6fdb15f0902eee0d4138b9d5975e244e6cc61ef302feaf37af24c7cb
# specify the prefixes of files and directories to ignore (by default common configuration files are ignored)
ignore-pre: # ex. [tests, readme.md, scripts]
# specify the extensions of files to ignore 
ignore-ext: # ex. [.go, .py, .js]
# specify the extensions of files to include 
include-ext: # ex. [.go, .py, .js]
`)

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
		configFileName := ".crev-config.yaml"

		// Check if the config file already exists
		if viper.ConfigFileUsed() != "" {
			fmt.Println("Config file already exists at ", viper.ConfigFileUsed())
			os.Exit(1)
		}

		// Write the default config using Viper
		err := os.WriteFile(configFileName, defaultConfig, 0644)
		if err != nil {
			fmt.Println("Unable to write config file: ", err)
			os.Exit(1)
		}

		// Inform the user
		fmt.Println("Config file created at:", configFileName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

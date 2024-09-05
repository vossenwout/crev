/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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
# specify the extensions of files to include
extensions: # ex. [.go, .py, .js]
`)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

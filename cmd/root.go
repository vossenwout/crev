// Description: This file contains the root command for the CLI tool.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Region string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crev",
	Short: "Initialize",
	Long: `Allows you to generate a textual representation of 
your code and let it be reviewed by an AI.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// otherwise the completion command will be available
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Viper will search for config in the following order:
	// First search in current directory
	viper.SetConfigType("yaml")
	viper.SetConfigName(".crev-config")
	//Second search in current directory
	viper.AddConfigPath(".")
	// Finally search home directory
	home, err := os.UserHomeDir()

	cobra.CheckErr(err)
	viper.AddConfigPath(home)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

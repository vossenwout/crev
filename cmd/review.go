package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if apiKey := viper.GetString("api-key"); apiKey == "" {
			log.Fatal(`API key is required for this command. Set the API key under the 'api-key' key in your .crev-config.yaml or provide it as a flag --api-key.`)
		}
		log.Println("Reviewing code...")
		dat, err := os.ReadFile(".crev-project-overview.txt")
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(dat))
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)
	reviewCmd.Flags().String("api-key", "", "Your Code AI Review API key")
	err := viper.BindPFlag("api-key", reviewCmd.Flags().Lookup("api-key"))
	if err != nil {
		log.Fatal(err)
	}
}

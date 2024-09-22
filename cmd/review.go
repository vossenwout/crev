package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vossenwout/crev/internal/review"
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
		apiKey := viper.GetString("crev_api_key")
		if apiKey == "" {
			log.Fatal(`Api key is required for review. Generate yours on: ... and set it as CREV_API_KEY env var or specify it under 'crev_api_key' key in your .crev-config.yaml. For more information see: ...`)
		}
		dat, err := os.ReadFile("crev-project.txt")
		if err != nil {
			log.Fatal("Could not find crev-project.txt. Did you forget to run the \"crev bundle\" command?")
		}
		review.Review(string(dat), apiKey)
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)
	reviewCmd.Flags().String("crev_api_key", "", "Your Code AI Review API key ")
	err := viper.BindPFlag("crev_api_key", reviewCmd.Flags().Lookup("crev_api_key"))
	if err != nil {
		log.Fatal(err)
	}
}

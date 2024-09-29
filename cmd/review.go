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
	Short: "Let an AI review your crev-project.txt",
	Long: `Let an AI review the crev-project.txt you generated with the crev bundle command. 

This command requires a CREV_API_KEY to be set as an environment variable or in your .crev-config.yaml.
You can generate an CREV_API_KEY on the crev website. For more information see: https://crevcli.com/docs`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("crev_api_key")
		if apiKey == "" {
			log.Fatal(`Api key is required for review. Get yours on: https://crevcli.com/api-key and set it as CREV_API_KEY env var or specify it under 'crev_api_key' key in your .crev-config.yaml. For more information see: https://crevcli.com/docs`)
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

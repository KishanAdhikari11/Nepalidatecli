package cmd

import (
	scrapper "nepalidate/Nepalidatescrapper"

	"github.com/spf13/cobra"
)

// nepalidateCmd represents the nepalidate command
var nepalidateCmd = &cobra.Command{
	Use:   "nepalidate",
	Short: "Nepali date CLI",
	Long:  `A simple cli app written in golang that displays nepalidate, time, festivals, panchang, thithi and english date.`,
	Run: func(cmd *cobra.Command, args []string) {
		GetNepaliDate()

	},
}

func init() {
	rootCmd.AddCommand(nepalidateCmd)

}

func GetNepaliDate() {
	scrapper.Scrape()
}

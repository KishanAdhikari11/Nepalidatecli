package cmd

import (
	"fmt"
	scrapper "nepalidate/Nepalidatescrapper"
	"os"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "nepalidate ",
	DisableFlagsInUseLine: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},

	Short: "Nepali date CLI",
	Long:  `A simple cli app written in golang that displays nepalidate, time, festivals, panchang, thithi and english date.`,
	Run: func(cmd *cobra.Command, args []string) {
		//wrap these things inside a table
		aaja := scrapper.Scrape()

		blue := color.New(color.FgBlue).SprintFunc()
		// green := color.New(color.FgCyan).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		fmt.Println(blue("Nepali Date: "), green(aaja.Date))
		fmt.Println(blue("English Date: "), green(aaja.EnglishDate))
		fmt.Println(blue("Time: "), green(aaja.Time))
		fmt.Println(blue("Thithi: "), green(aaja.Thithi))
		fmt.Println(blue("Panchang: "), green(aaja.Panchang))
		fmt.Println(blue("Event: "), green(aaja.Event))

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}

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

		green := color.New(color.FgGreen).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		fmt.Printf("\n")
		fmt.Println(cyan("\t\tDate(B.S):  "), green(aaja.Date))
		fmt.Println(cyan("\t\tDate(A.D):  "), green(aaja.EnglishDate))
		fmt.Println(cyan("\t\tTime:      "), green(aaja.Time))
		fmt.Println(cyan("\t\tEvents:     "), green(aaja.Event))
		fmt.Println(cyan("\t\tPanchang:   "), green(aaja.Panchang))
		fmt.Println(cyan("\t\tThithi:     "), green(aaja.Thithi))
		fmt.Printf("\n")

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

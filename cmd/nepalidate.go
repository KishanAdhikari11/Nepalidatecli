package cmd

import (
	scrapper "nepalidate/Nepalidatescrapper"

	"github.com/spf13/cobra"
)

// nepalidateCmd represents the nepalidate command
var nepalidateCmd = &cobra.Command{
	Use:   "nepalidate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
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

// type Jokes struct {
// 	ID     string `json:"id"`
// 	Joke   string `json:"joke"`
// 	Status int    `json:"Status"`
// }

// func getRandomJoke() {
// 	url := "https://icanhazdadjoke.com/"
// 	responseBytes := getJokeData(url)
// 	random := Jokes{}
// 	json.Unmarshal(responseBytes, &random)
// 	fmt.Println(string(random.Joke))

// }
// func getJokeData(baseAPI string) []byte {
// 	req, err := http.NewRequest("GET", baseAPI, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("User-Agent", "NepaliDate/1.0")

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Printf("Could not make a request %v", err)
// 	}
// 	responseBytes, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return responseBytes

// }

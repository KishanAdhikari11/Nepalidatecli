package Nepalidatescrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type NepaliDate struct {
	Date        string `json:"date"`
	Thithi      string `json:"thithi"`
	Event       string `json:"event"`
	Panchang    string `json:"panchang"`
	EnglishDate string `json:"englishdate"`
	Time        string `json:"time"`
}

func Scrape() {
	c := colly.NewCollector()
	date := NepaliDate{}

	c.OnHTML(".logo", func(e *colly.HTMLElement) {
		date.Date = e.ChildText(".nep")
	})
	c.OnHTML(".time", func(e *colly.HTMLElement) {
		date.Time = e.DOM.Find("span:not(.eng)").Text()
	})
	c.OnHTML(".time", func(e *colly.HTMLElement) {
		date.EnglishDate = e.ChildText(".eng")
	})
	c.OnHTML("a[class='event']", func(e *colly.HTMLElement) {

		date.Event = e.Text
	})
	c.OnHTML("[style='line-height: 1.9']", func(e *colly.HTMLElement) {
		e.DOM.Find("a").Remove()
		date.Panchang = e.DOM.Text()
	})
	c.OnHTML("[style='margin: 10px 0; color: white; font-size: 1.3rem']", func(e *colly.HTMLElement) {
		date.Thithi = e.Text
	})

	err := c.Visit("https://www.hamropatro.com/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Date : %s \n", date.Date)
	fmt.Printf("Time : %s \n", date.Time)
	fmt.Printf("English Date : %s \n", date.EnglishDate)
	fmt.Printf("Event : %s \n", date.Event)
	fmt.Printf("Panchang : %s \n", date.Panchang)
	fmt.Printf("Thithi : %s \n", date.Thithi)

}

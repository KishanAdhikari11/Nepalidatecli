package Nepalidatescrapper

import (
	"fmt"
	"strings"

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

func Scrape() NepaliDate {
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
	CleanStruct(&date)
	return date

}

func CleanStruct(s *NepaliDate) {
	s.Thithi = strings.TrimSpace(s.Thithi)
	s.Panchang = strings.Replace(s.Panchang, "पञ्चाङ्ग:", "", -1)
	s.Panchang = strings.TrimSpace(s.Panchang)
	s.Event = strings.TrimSpace(s.Event)

}

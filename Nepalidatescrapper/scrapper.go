package Nepalidatescrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type NepaliDate struct {
	Day              string `json:"day"`
	Date             string `json:"date"`
	Year             string `json:"year"`
	DayOfWeek        string `json:"day_of_week"`
	EnglishDate      string `json:"english_date"`
	EnglishDayOfWeek string `json:"english_day_of_week"`
	EnglishMonth     string `json:"english_month"`
	EnglishYear      string `json:"english_year"`
	Festival         string `json:"festival"`
	Events           string `json:"events"`
}

func Scrape() {
	c := colly.NewCollector()
	err := c.Visit("https://www.hamropatro.com/")
	if err != nil {
		fmt.Println(err)
	}

}

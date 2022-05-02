package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	fileName := "data.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("www.monsterindia.com"),
	)

	c.OnHTML(".card-apply-content", func(h *colly.HTMLElement) {
		writer.Write([]string{
			h.ChildText("a"),
			h.ChildAttr("a", "href"),
		})
	})
	c.Visit("https://www.monsterindia.com/search/golang-jobs?searchId=8fb3386f-2b19-4ce5-8a92-ee2dce121d9a")
	log.Printf("Scraping Finished\n")
	log.Println(c)
}

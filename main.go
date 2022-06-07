package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	crawl()
}

func crawl() {

	type data struct {
		City               string
		Type_of_employment string
		Experience         string
		Company            string
	}

	c := colly.NewCollector()
	infoCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Profile URL: ", r.URL.String())
	})

	c.Visit("https://justjoin.it/all/go")

	c.OnHTML(".css-1smbjja", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnHTML(".mode-detail", func(e *colly.HTMLElement) {
		profileUrl := e.ChildAttr(".css-ic7v2w div", "href")
		profileUrl = e.Request.AbsoluteURL(profileUrl)
		infoCollector.Visit(profileUrl)
	})

	infoCollector.OnHTML(".css-vuh3mm", func(e *colly.HTMLElement) {
		tmpProfile := data{}
		tmpProfile.City = e.ChildText("span.css-9wmrp4")
		tmpProfile.Type_of_employment = e.ChildText("span.css-ioglek")
		tmpProfile.Experience = e.ChildText(".css-1ji7bvd")
		tmpProfile.Company = e.ChildText(".css-1npsplc")

		js, err := json.MarshalIndent(tmpProfile, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))

	})

}

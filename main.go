package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type StockPrices struct {
	Name  string
	Price string
}

func main() {

	// saving the data as a .csv file
	file, err := os.Create("CryptoData.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write CSV header

	headers := []string{"Name", "Price"}
	writer.Write(headers)

	// actually scraping the data
	c := colly.NewCollector()

	c.OnHTML(".lcw-table-container.main-table", func(e *colly.HTMLElement) {

		data := StockPrices{}
		data.Name = e.ChildText(".filter-item-name.mb0.text-left")
		data.Price = e.ChildText("td.filter-item.table-item.main-price")

		row := []string{data.Name, data.Price}
		writer.Write(row)

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.livecoinwatch.com/")

}

// package main

// import (
// 	"encoding/csv"
// 	"log"
// 	"os"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	fName := "cryptocoinmarketcap.csv"
// 	file, err := os.Create(fName)
// 	if err != nil {
// 		log.Fatalf("Cannot create file %q: %s\n", fName, err)
// 		return
// 	}
// 	defer file.Close()
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// Write CSV header
// 	writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

// 	// Instantiate default collector
// 	c := colly.NewCollector()

// 	c.OnHTML("#currencies-all tbody tr", func(e *colly.HTMLElement) {
// 		writer.Write([]string{
// 			e.ChildText(".currency-name-container"),
// 			e.ChildText(".col-symbol"),
// 			e.ChildAttr("a.price", "data-usd"),
// 			e.ChildAttr("a.volume", "data-usd"),
// 			e.ChildAttr(".market-cap", "data-usd"),
// 			e.ChildText(".percent-1h"),
// 			e.ChildText(".percent-24h"),
// 			e.ChildText(".percent-7d"),
// 		})
// 	})

// 	c.Visit("https://coinmarketcap.com/all/views/all/")

// 	log.Printf("Scraping finished, check file %q for results\n", fName)
// }

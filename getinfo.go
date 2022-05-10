package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

//type House struct {

//}

func main() {
	go func(msg string) {
		c := colly.NewCollector()

		// On every a element which has href attribute call callback
		//var inStock bool

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			// Print link
			fmt.Println(e.Text)
		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		// Start scraping on https://hackerspaces.org
		c.Visit("https://www.linkedin.com/jobs/entry-level-software-engineer-jobs/")

		c2 := colly.NewCollector()

		// On every a element which has href attribute call callback
		//var inStock bool

		c2.OnHTML("a[href]", func(e *colly.HTMLElement) {
			// Print link
			fmt.Println(e.Text)
		})

		// Before making a request print "Visiting ..."
		c2.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		// Start scraping on https://hackerspaces.org
		c2.Visit("https://www.indeed.com/q-Entry-Level-Software-Engineer-jobs.html?vjk=74a4204805ec6555")
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}

package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

//type House struct {

//}

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("https://www.linkedin.com/jobs/entry-level-software-engineer-jobs/")

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

}

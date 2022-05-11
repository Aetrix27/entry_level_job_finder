package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func scrape(link string) {

	c := colly.NewCollector()

	// On every a element which has href attribute call callback

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Print link
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(link)

}
func main() {
	var websites []string

	for i := 0; i < 3; i++ {
		// Taking input from user
		//reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a job board to scrape: ")
		var jobBoard string

		// Taking input from user
		fmt.Scanln(&jobBoard)

		websites = append(websites, jobBoard)

	}

	//https://www.linkedin.com/jobs/entry-level-software-engineer-jobs/
	//https://www.indeed.com/q-Entry-Level-Software-Engineer-jobs.html?vjk=74a4204805ec6555

	go scrape(websites[0])
	go scrape(websites[1])
	go scrape(websites[2])
	time.Sleep(time.Second)

	fmt.Println("done")

}

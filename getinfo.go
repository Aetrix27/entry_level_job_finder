package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func scrape(link string, c *colly.Collector) {

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Print link
		output := strings.Fields(e.Text)
		for i := 0; i < len(output); i++ {
			if output[i] == "Entry" || output[i] == "entry" {
				if i+8 < len(output) {
					fmt.Println(output[i : i+7])
				}
			}

		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(link)

}

func main() {
	var websites []string
	c := colly.NewCollector()
	option := 0
	count := 0
	add_another := ""

	fmt.Println("Welcome to the Entry Level Job Board Scraper! Please enter some job boards: ")
	for option < 1 {
		reader := bufio.NewReader(os.Stdin)

		// Taking input from user
		//reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a job board to scrape: ")
		var jobBoard string

		// Taking input from user
		fmt.Scanln(&jobBoard)

		websites = append(websites, jobBoard)

		fmt.Print("Add another page? ")
		add_another, _ = reader.ReadString('\n')
		input2 := strings.TrimSpace(add_another)
		fmt.Print(add_another)
		count += 1
		if input2 == "yes" {
			continue
		} else if input2 == "no" {
			option = 1
			break
		}

	}

	//https://www.simplyhired.com/search?q=entry+level+software+engineer&l=los+angeles%2C+ca&job=3C1Bn0XyBJz10GUVBrsWKVgDDJ5wWIARRt-44vek8cIlDM_RRh7zZg
	//https://www.indeed.com/q-Entry-Level-Software-Engineer-jobs.html?vjk=74a4204805ec6555

	for i := 0; i < count; i++ {

		scrape(websites[i], c)
		time.Sleep(time.Second * 3)
	}

}

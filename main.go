package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"sync"
	"time"
)

// Ebay.
func GetEbayTags(ask string) {
	// The goroutine is done.
	// Defer statements are executed the last.
	defer wg.Done()

	// Creating url using given request and checking its correctness.
	var url string = fmt.Sprintf("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=%s&_sacat=0", ask)
	fmt.Println(url)

	// Creating request and checking if it can be opened.
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(1)
		return
	}

	// Closing request before leaving the function.
	// Defer statements are executed the last.
	defer res.Body.Close()

	// If cannot work with res.
	if res.StatusCode != 200 {
		fmt.Println(2)
		return
	}

	// Creating page source .html document and checking if it can be created.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(3)
		return
	}

	// Searching all needed div tags.
	doc.Find("div[class~=srp-related-searches]").Each(func(i int, s *goquery.Selection) {
		// Searching all needed span tags inside the div tags.
		s.Find("span[class=BOLD]").Each(func(i int, ss *goquery.Selection) {
			// Logging.
			fmt.Printf("%s ", ss.Text())
		})
	})

	// Just a new line.
	fmt.Println("")
}

// Wildberries.
func GetWildberriesTags(ask string) {
	// The goroutine is done.
	// Defer statements are executed the last.
	defer wg.Done()

	// Creating url using given request and checking its correctness.
	var url string = fmt.Sprintf("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=%s&_sacat=0", ask)
	fmt.Println(url)

	// TODO...

	// Just a new line.
	fmt.Println("")
}

func GetTags(ask string) {
	// Giving number of goroutines which we are going to wait.
	wg.Add(2)

	// Running both goroutines.
	go GetEbayTags(ask)
	go GetWildberriesTags(ask)

	// Waiter waits at this part of code.
	wg.Wait()
}

// Goroutines waiter.
var wg sync.WaitGroup

func main() {
	// Setting timer.
	start := time.Now()

	// If string is given, using it, otherwise set "anime", because anime rules.
	var ask string
	if len(os.Args) > 1 {
		ask = os.Args[1]
	} else {
		ask = "anime"
	}

	// Getting all tags.
	GetTags(ask)

	// Printing the total time elapsed.
	fmt.Printf("%v\n", time.Since(start))
}

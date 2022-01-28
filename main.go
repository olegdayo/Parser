package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

// Ebay.
func GetEbayTags(ask string, wg *sync.WaitGroup) ([]string, error) {
	// The goroutine is done.
	// Defer statements are executed the last.
	defer wg.Done()
	// Answer.
	var tags []string

	// Creating url using given request and checking its correctness.
	var url string = fmt.Sprintf("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=%s&_sacat=0", ask)
	fmt.Println(url)

	// Creating request and checking if it can be opened.
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(1)
		return tags, err
	}

	// Closing request before leaving the function.
	// Defer statements are executed the last.
	defer res.Body.Close()

	// If cannot work with res.
	if res.StatusCode != 200 {
		fmt.Println(2)
		return tags, err
	}

	// Creating page source .html document and checking if it can be created.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(3)
		return tags, err
	}

	// Searching all needed div tags.
	doc.Find("div[class~=srp-related-searches]").Each(func(i int, s *goquery.Selection) {
		// Searching all needed span tags inside the div tags.
		s.Find("span[class=BOLD]").Each(func(i int, ss *goquery.Selection) {
			// Logging.
			tags = append(tags, ss.Text())
		})
	})

	// Everything is ok.
	return tags, nil
}

// Wildberries.
func GetWildberriesTags(ask string, wg *sync.WaitGroup) ([]string, error) {
	// The goroutine is done.
	// Defer statements are executed the last.
	defer wg.Done()
	// Answer.
	var tags []string

	// Creating url using given request and checking its correctness.
	var url string = fmt.Sprintf("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=%s&_sacat=0", ask)
	fmt.Println(url)

	// TODO...

	// Everything is ok.
	return tags, nil
}

func GetTags(ask string, wg *sync.WaitGroup) ([]string, []string) {
	// Giving number of goroutines which we are going to wait.
	wg.Add(2)
	var ebayTags []string
	var ebayError error
	var wildberriesTags []string
	var wildberriesError error

	// Running both goroutines.
	go func() {
		ebayTags, ebayError = GetEbayTags(ask, wg)
	}()
	go func() {
		wildberriesTags, wildberriesError = GetWildberriesTags(ask, wg)
	}()

	// Waiter waits at this part of code.
	wg.Wait()

	if ebayError != nil {
		fmt.Println("Something went wrong in Ebay parsing!")
	}
	if wildberriesError != nil {
		fmt.Println("Something went wrong in Wildberries parsing!")
	}

	return ebayTags, wildberriesTags
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Setting maximum number of threads we need to use.
	runtime.GOMAXPROCS(min(runtime.NumCPU(), 2))
	// Setting timer.
	start := time.Now()
	// Goroutines waiter.
	var wg sync.WaitGroup

	// If string is given, using it, otherwise set "anime", because anime rules.
	var ask string
	if len(os.Args) > 1 {
		ask = os.Args[1]
	} else {
		ask = "anime"
	}

	// Getting all tags.
	ebay, wildberries := GetTags(ask, &wg)
	fmt.Println(ebay)
	fmt.Println(wildberries)

	// Printing the total time elapsed.
	fmt.Printf("%v\n", time.Since(start))
}

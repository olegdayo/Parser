package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"sync"
	"time"
)

func GetEbayTags(ask string) {
	defer wg.Done()

	var url string = fmt.Sprintf("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=%s&_sacat=0", ask)
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(1)
		return
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println(2)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(3)
		return
	}

	doc.Find("div[class~=srp-related-searches]").Each(func(i int, s *goquery.Selection) {
		s.Find("span[class=BOLD]").Each(func(i int, ss *goquery.Selection) {
			fmt.Printf("%s ", ss.Text())
		})
	})
	fmt.Println("")
}

func GetWildberriesTags(ask string) {
	defer wg.Done()

	var url string = fmt.Sprintf("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=%s&_sacat=0", ask)
	fmt.Println(url)

	// TODO...

	fmt.Println("")
}

func GetTags(ask string) {
	wg.Add(2)
	go GetEbayTags(ask)
	go GetWildberriesTags(ask)
	wg.Wait()
}

var wg sync.WaitGroup

func main() {
	start := time.Now()
	var ask string
	if len(os.Args) == 2 {
		ask = os.Args[1]
	} else {
		ask = "anime"
	}
	GetTags(ask)
	fmt.Printf("%v\n", time.Since(start))
}

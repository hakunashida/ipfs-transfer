package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	// OPTION A:
	// "github.com/djimenez/iconv-go"
	"github.com/PuerkitoBio/goquery"
	"goji.io"
	"goji.io/pat"
)

func ExampleScrape() {

	// OPTION A: check page encoding and convert to utf-8 if necessary
	/*url := "http://metalsucks.net"

	// Load the URL
	res, err := http.Get(url)
	if err != nil {
	    log.Fatal(err)
	}
	defer res.Body.Close()

	// Convert the designated charset HTML to utf-8 encoded HTML.
	// `charset` being one of the charsets known by the iconv package.
	// TODO: detect encoding on page
	utfBody, err := iconv.NewReader(res.Body, "utf-8", "utf-8")
	if err != nil {
	    log.Fatal(err)
	}

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
	    log.Fatal(err)
	}*/

	// OPTION B: just fetch the page if all scraped pages are utf-8
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

/*var (
	// Protect access to dup
	mu sync.Mutex
	// Duplicates table
	dup = map[string]bool{}

	// Command-line flags
	seed        = flag.String("seed", "https://ultimate-guitar.com", "seed URL")
	cancelAfter = flag.Duration("cancelafter", 0, "automatically cancel the fetchbot after a given time")
	cancelAtURL = flag.String("cancelat", "", "automatically cancel the fetchbot at a given URL")
	stopAfter   = flag.Duration("stopafter", 0, "automatically stop the fetchbot after a given time")
	stopAtURL   = flag.String("stopat", "", "automatically stop the fetchbot at a given URL")
	memStats    = flag.Duration("memstats", 0, "display memory statistics at a given interval")
)*/

func main() {

	beginFetching()

	fmt.Printf("ipfs-transfer started")

	// 1. bind to port
	gojiMux := goji.NewMux()
	gojiMux.HandleFunc(pat.Get("/hello/:name"), hello)
	http.ListenAndServe("localhost:8000", gojiMux)

	// 2. run scrape example
	// ExampleScrape()

	flag.Parse()
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

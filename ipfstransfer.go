package main

import (
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

func main() {

	connectDb()
	beginFetching()

	fmt.Println("ipfs-transfer started")

	// 1. bind to port
	gojiMux := goji.NewMux()
	gojiMux.HandleFunc(pat.Get("/hello/:name"), hello)
	gojiMux.HandleFunc(pat.Get("/search/:terms"), search)
	http.ListenAndServe("localhost:8000", gojiMux)

	defer disconnectDb()

	// 2. run scrape example
	// ExampleScrape()
}

func search(w http.ResponseWriter, r *http.Request) {
	terms := pat.Param(r, "terms")
	refs := searchDb(terms)
	fmt.Fprintf(w, "Searching for: %s\n", terms)
	fmt.Fprintf(w, "found:\n%s", refs)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

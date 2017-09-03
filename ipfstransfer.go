package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"goji.io"
	"goji.io/pat"
)

func main() {

	connectDb()
	beginFetching()

	fmt.Println("ipfs-transfer started")

	gojiMux := goji.NewMux()
	gojiMux.HandleFunc(pat.Get("/hello/:name"), hello)
	gojiMux.HandleFunc(pat.Get("/search/:terms"), search)
	http.ListenAndServe("localhost:8000", gojiMux)

	defer disconnectDb()
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

package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TabsShow(w http.ResponseWriter, r *http.Request) {
	tabs := printDb()
	w.Header().Set("Content-Type", "application.json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tabs); err != nil {
		panic(err)
	}
}

func TabsSearch(w http.ResponseWriter, r *http.Request) {

	// TODO: should use a querystring to perform RESTful searches
	// https://stackoverflow.com/questions/207477/restful-url-design-for-search

	vars := mux.Vars(r)
	searchTerm := vars["searchTerm"]

	results := searchDb(searchTerm)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if len(results) > 0 {
		if err := json.NewEncoder(w).Encode(results); err != nil {
			panic(err)
		}
		return
	}
}

/*func search(w http.ResponseWriter, r *http.Request) {
	terms := pat.Param(r, "terms")
	refs := searchDb(terms)
	fmt.Fprintf(w, "Searching for: %s\n", terms)
	fmt.Fprintf(w, "found:\n%s", refs)
}*/

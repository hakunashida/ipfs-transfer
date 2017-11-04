package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
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

/*func search(w http.ResponseWriter, r *http.Request) {
	terms := pat.Param(r, "terms")
	refs := searchDb(terms)
	fmt.Fprintf(w, "Searching for: %s\n", terms)
	fmt.Fprintf(w, "found:\n%s", refs)
}*/

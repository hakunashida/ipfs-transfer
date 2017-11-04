package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Body string
}

type TabsResponse struct {
	Data []TabReference `json:"data"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TabsShow(w http.ResponseWriter, r *http.Request) {
	tabs := getAllTabs()
	w.Header().Set("Content-Type", "application.json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	res := TabsResponse{tabs}
	if err := json.NewEncoder(w).Encode(res); err != nil {
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
	}
}

func TabContent(w http.ResponseWriter, r *http.Request) {

	// find the tab from the id
	vars := mux.Vars(r)
	id := vars["id"]
	tab := getTabById(id)

	// set the response headers
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// send back the tab content
	if tab.IpfsHash != "" {
		content := ipfsLoad(tab.IpfsHash)
		res := Message{content}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		return
	}

	// this should never happen, but in case no content is found...
	res := Message{"No tab content was found with the id " + id}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func Reset(w http.ResponseWriter, r *http.Request) {
	clearDb()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	res := Message{"Reset successful"}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

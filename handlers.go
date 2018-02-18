package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type MessageResponse struct {
	Body string `json:"message"`
}

type TabResponse struct {
	Data TabReference `json:"tab"`
}

type TabsResponse struct {
	Data []TabReference `json:"tabs"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func TabsShow(w http.ResponseWriter, r *http.Request) {
	tabs := getAllTabs()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
	searchResults := searchDb(searchTerm)

	// return an empty array if the search does not find any matches
	tabs := []TabReference{}
	if len(searchResults) > 0 {
		tabs = searchResults
	}
	res := TabsResponse{Data: tabs}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func TabGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	artist := vars["artist"]
	name := vars["name"]
	tab, foundOne := getTabByArtistAndName(artist, name)

	res := TabResponse{Data: tab}
	if !foundOne {
		// TODO: send empty json back
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
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
		res := MessageResponse{content}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		return
	}

	// this should never happen, but in case no content is found...
	res := MessageResponse{"No tab content was found with the id " + id}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func Reset(w http.ResponseWriter, r *http.Request) {
	clearDb()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	res := MessageResponse{"Reset successful"}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

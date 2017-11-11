package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TabsShow",
		"GET",
		"/tabs",
		TabsShow,
	},
	Route{
		"TabsSearch",
		"GET",
		"/tabs/search/{searchTerm}",
		TabsSearch,
	},
	Route{
		"TabContent",
		"GET",
		"/tabs/{id}/content",
		TabContent,
	},

	// TODO: this route should of course be removed before pushing to production!
	Route{
		"Reset",
		"POST",
		"/reset",
		Reset,
	},
}

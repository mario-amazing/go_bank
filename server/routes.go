package server

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
		"ClientIndex",
		"GET",
		"/clients",
		ClientIndex,
	},
	Route{
		"ClientShow",
		"GET",
		"/clients/{id}",
		ClientShow,
	},
	Route{
		"ClientCreate",
		"POST",
		"/clients",
		ClientCreate,
	},
}

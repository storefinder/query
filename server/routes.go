package server

import (
	"net/http"

	"github.com/storefinder/query/handlers"
)

//Route defines an http route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a collection of http routes supported by this web application
type Routes []Route

//BuildRoutes configures routes supported by the indexer
func BuildRoutes() []Route {
	var routes = Routes{
		Route{
			"StoreQuery",
			"POST",
			"/1.0/index/{indexName}/search",
			handlers.Search(),
		},
	}

	return routes
}

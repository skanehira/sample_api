package main

import (
	"net/http"
	"strings"
)

type QueryParams map[string]string

// Route http request info
type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	QueryParams QueryParams
}

// Router http request routing
type Router struct {
	Route map[string]*Route
}

// NewRoute generate new Route
func NewRoute(method, path string, handler http.HandlerFunc) *Route {
	r := &Route{
		Method:      method,
		Path:        getPath(path),
		Handler:     handler,
		QueryParams: getQueryParams(path),
	}
	return r
}

// NewRouter generate new router
func NewRouter() *Router {
	return &Router{
		Route: make(map[string]*Route),
	}
}

// Add add new route
func (r *Router) Add(method, path string, handler http.HandlerFunc) {
	// if url is not exsist
	if len(path) < 1 {
		panic("invalid path")
	}

	if path[0] != '/' {
		path = "/" + path
	}

	// Add new route to Router
	r.Route[path+method] = NewRoute(method, path, handler)
}

// add query parameter to route
func getQueryParams(path string) QueryParams {
	queryStartIndex := strings.Index(path, "?")
	queryParams := QueryParams{}

	// when query parameters is not exsist
	if queryStartIndex == -1 {
		return QueryParams{}
	}

	// Get query parameters
	params := strings.Split(path[queryStartIndex+1:], "?")

	// Add query parameters to map
	for _, q := range params {
		param := strings.Split(q, "=")
		if len(param) == 2 && param[0] != "" {
			queryParams[param[0]] = param[1]
		}
	}

	return queryParams
}

// getPath get URI
func getPath(path string) string {
	separators := []string{"/", "?"}

	for _, sep := range separators {
		pl := strings.Index(path[1:], sep)
		if pl != -1 {
			return path[:pl+1]
		}
	}

	return path
}

package main

import (
	"net/http"
	"strings"
)

type QueryParams map[string]string
type PathParams map[string]string

// Route http request info
type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	QueryParams QueryParams
	PathParams  PathParams
}

// Router http request routing
type Router struct {
	Route map[string]*Route
}

// NewRoute generate new Route
func NewRoute(method, path string, handler http.HandlerFunc) *Route {
	r := &Route{
		Method:      method,
		Path:        path,
		Handler:     handler,
		QueryParams: getQueryParams(path),
		PathParams:  PathParams{},
	}
	return r
}

// NewRouter generate new router
func NewRouter() *Router {
	return new(Router)
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
		queryParams[param[0]] = param[1]
	}

	return queryParams
}

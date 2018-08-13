package main

import (
	"net/http"
)

// Route http request info
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Router http request routing
type Router struct {
	Route map[string]*Route
}

// NewRoute generate new Route
func NewRoute(method, path string, handler http.HandlerFunc) *Route {
	r := &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
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

// Find find tareget route
func (r *Router) Find(request *http.Request) *Route {
	handlerKey := request.URL.Path + request.Method
	return r.Route[handlerKey]
}

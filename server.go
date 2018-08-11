package main

import (
	"fmt"
	"log"
	"net/http"
)

// Server http server
type Server struct {
	Address string
	Router  *Router
}

// NewServer generate new server
func NewServer(address string) *Server {
	return &Server{
		Address: address,
		Router:  NewRouter(),
	}
}

// POST add route POST
func (s *Server) POST(path string, handler http.HandlerFunc) {
	s.Router.Add(http.MethodPost, path, handler)
}

// GET add route GET
func (s *Server) GET(path string, handler http.HandlerFunc) {
	s.Router.Add(http.MethodGet, path, handler)
}

// PUT add route PUT
func (s *Server) PUT(path string, handler http.HandlerFunc) {
	s.Router.Add(http.MethodPut, path, handler)
}

// DELETE add route DELETE
func (s *Server) DELETE(path string, handler http.HandlerFunc) {
	s.Router.Add(http.MethodDelete, path, handler)
}

// ServeHTTP implements http ServeHTTP
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerKey := getPath(r.URL.Path) + r.Method

	if route, ok := s.Router.Route[handlerKey]; ok {
		route.Handler(w, r)
	}
}

// Start http server start
func (s *Server) Start() {
	// start server
	server := &http.Server{Addr: s.Address, Handler: s}
	fmt.Printf("http server start port in %s\n", s.Address)
	log.Fatal(server.ListenAndServe())
}

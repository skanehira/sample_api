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
	route := s.Router.Find(r)
	route.Handler(w, r)
}

// Start http server start
func (s *Server) Start() {
	fmt.Printf("http server start port in %s\n", s.Address)

	// start server
	log.Fatal(http.ListenAndServe(s.Address, s))
}

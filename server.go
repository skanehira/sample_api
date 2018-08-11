package main

import (
	"log"
	"net/http"
)

// Server http server
type Server struct {
	Address string
	Port    string
	Router  *Router
}

// NewServer generate new server
func NewServer(address, port string) *Server {
	return &Server{
		Address: address,
		Port:    port,
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

}

// ServerStart http server start
func (s *Server) ServerStart() {
	// start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

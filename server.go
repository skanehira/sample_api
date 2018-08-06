package main

import (
	"log"
	"net/http"
)

func main() {
	// http server start
	ServerStart()
}

// ServerStart http server start
func ServerStart() {

	// regist handler
	http.HandleFunc("/users", UserHandler)

	// start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

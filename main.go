package main

func main() {
	// new server
	s := NewServer(":8080")

	// add route
	s.POST("/users", UserRegister)
	s.GET("/users", UserReader)
	s.PUT("/users", UserUpdater)
	s.DELETE("/users", UserDeleter)

	// http server start
	s.Start()
}

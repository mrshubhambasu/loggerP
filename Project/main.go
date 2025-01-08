package main

import (
	"authservice/routes"
	"fmt"
	"net/http"
)

func main() {
	r := routes.InitRoutes()

	// Start server
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", r)
}

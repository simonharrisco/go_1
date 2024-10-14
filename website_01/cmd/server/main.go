package main

import (
	"fmt"
	"log"
	"net/http"
	"website-01/internal/routes"
)

func main() {

	// intialise database
	// 		connect
	// 		run migrations
	// 		seed data
	//
	//
	// initialise routes
	//
	// start server

	routes.InitialiseRoutes()

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

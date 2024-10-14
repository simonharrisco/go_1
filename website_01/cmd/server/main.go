package main

import (
	"fmt"
	"log"
	"net/http"

	"website-01/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandleRoot)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

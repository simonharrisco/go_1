package routes

import (
	"net/http"

	"website-01/internal/handlers"
)

func InitialiseRoutes() {
	http.HandleFunc("/", handlers.HandleRoot)
	http.HandleFunc("/about", handlers.HandleAbout)
}

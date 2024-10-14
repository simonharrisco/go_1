package handlers

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to your Go webserver!")
}

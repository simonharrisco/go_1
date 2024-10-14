package handlers

import (
	"fmt"
	"net/http"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handle About")
}

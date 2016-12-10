package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.ListenAndServe(":8080", &server{})
}

type server struct{}

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	fmt.Printf("Handling: %s\n", uri)
	if strings.Contains(uri, "PerformNextMove") {
		performNextMove(w, r)
	} else {
		infoHandler(w, r)
	}
}

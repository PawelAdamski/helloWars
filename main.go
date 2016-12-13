package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", infoHandler)
	http.HandleFunc("/PerformNextMove", performNextMoveHandler)
	http.HandleFunc("/Info", infoHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Errorf("Error: ", err)
	}
}

type server struct{}

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	fmt.Printf("Handling: %s\n", uri)
	if strings.Contains(uri, "PerformNextMove") {
		performNextMoveHandler(w, r)
	} else {
		infoHandler(w, r)
	}
}

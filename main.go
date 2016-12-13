package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func main() {
	http.HandleFunc("/", infoHandler)
	http.HandleFunc("/PerformNextMove", performNextMoveHandler)
	http.HandleFunc("/Info", infoHandler)

	http.ListenAndServe(":8080", nil)
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

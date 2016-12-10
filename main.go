package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/Info/", infoHandler)
	http.HandleFunc("/PerformNextMove/", performNextMove)
	http.ListenAndServe(":8080", nil)
}

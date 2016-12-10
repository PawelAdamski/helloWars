package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/Info/", infoHandler)
	http.ListenAndServe(":8080", nil)
}

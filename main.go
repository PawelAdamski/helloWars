package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"

	"github.com/PawelAdamski/helloWars/bf"
	"github.com/PawelAdamski/helloWars/game"
)

var manualMod = flag.Bool("manual", false, "enable manual mode")

var port = "8080"
var BotAlgorithm game.Algorithm

func main() {
	flag.Parse()
	fmt.Println(*manualMod)
	if *manualMod {
		fmt.Println("Manual mode")
		BotName = "manual"
		BotAlgorithm = bf.Strategy{}
		port = "8081"
	} else {
		BotAlgorithm = bf.Strategy{}
	}

	http.HandleFunc("/", infoHandler)
	http.HandleFunc("/PerformNextMove", performNextMoveHandler)
	http.HandleFunc("/Info", infoHandler)

	err := http.ListenAndServe(":"+port, nil)
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

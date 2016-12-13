package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PawelAdamski/helloWars/bf"
	"github.com/PawelAdamski/helloWars/game"
)

func performNextMoveHandler(w http.ResponseWriter, r *http.Request) {
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Can't read input", err)
	}
	gs := game.State{}
	err = json.Unmarshal(bb, &gs)
	if err != nil {
		fmt.Println("Can't unmarshal", err, string(bb))
	}
	nm := bf.NextMove(&gs)
	bb, err = json.Marshal(nm)
	if err != nil {
		fmt.Println("Can't marshall BotMove")
	}
	fmt.Println(gs.BotID, gs.BotLocation, nm.String())
	fmt.Fprint(w, string(bb))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func performNextMove(w http.ResponseWriter, r *http.Request) {
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Can't read input", err)
	}
	gs := GameState{}
	err = json.Unmarshal(bb, &gs)
	if err != nil {
		fmt.Println("Can't unmarshal", err)
	}
	nm := BotMove{
		Direction: Up,
		Action:    None,
	}
	bb, err = json.Marshal(nm)
	if err != nil {
		fmt.Println("Can't marshall BotMove")
	}
	fmt.Fprint(w, string(bb))
}

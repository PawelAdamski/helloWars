package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type info struct {
	Name        string `json:"Name"`
	Avatar      string `json:"AvatarUrl"`
	Description string `json:"Description"`
	GameType    string `json:"GameType"`
}

var BotName = "PAKD"

func infoHandler(w http.ResponseWriter, r *http.Request) {
	i := info{
		Name:        BotName,
		Avatar:      "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Ladle_steel_18j07.JPG/1920px-Ladle_steel_18j07.JPG",
		Description: "All your base are belong to us",
		GameType:    "TankBlaster",
	}
	s, err := json.Marshal(i)
	if err != nil {
		fmt.Println("Can't marshall info", err)
	}
	fmt.Fprint(w, string(s))
}

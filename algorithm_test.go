package main

import (
	"encoding/json"
	"testing"
)

func TestParsing(t *testing.T) {
	input := `{
    "RoundNumber": 3,
    "BotId": "a88454b0‐80ba‐4c10‐b162‐f1ca766f1e3f",
    "Board": [
        [
            2,
            2,
            3
        ],
        [
            0,
            0,
            1
        ],
        [
            0,
            0,
            0
        ]
    ],
    "BotLocation": "0, 0",
    "IsMissileAvailable": true,
    "OpponentLocations": [
        "1, 1"
    ],
    "Bombs": [
        {
            "RoundsUntilExplodes": 3,
            "Location": "0, 1",
            "ExplosionRadius": 2
        }
    ],
    "Missiles": [
        {
            "MoveDirection": 3,
            "Location": "1, 0",
            "ExplosionRadius": 2
        }
    ],
    "GameConfig": {
        "MapWidth": 20,
        "MapHeight": 20,
        "BombBlastRadius": 2,
        "MissileBlastRadius": 2,
        "RoundsBetweenMissiles": 5,
        "RoundsBeforeIncreasingBlastRadius": 70,
        "IsFastMissileModeEnabled": true
    }
}`
	state := GameState{}
	err := json.Unmarshal([]byte(input), &state)
	if err != nil {
		t.Fatal(err)
	}

}

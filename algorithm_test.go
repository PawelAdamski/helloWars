package main

import (
	"encoding/json"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ParsingSuite struct{}

var _ = Suite(&ParsingSuite{})

func (s *ParsingSuite) TestParsing(c *C) {
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
	c.Assert(err, IsNil)
	c.Assert(state.BotLocation, Equals, Location{x: 0, y: 0})
	expectedBomb := Bomb{
		Location:            Location{x: 0, y: 1},
		ExplosionRadius:     2,
		RoundsUntilExplodes: 3,
	}
	if state.GameConfig.MapHeight != 20 {
		t.Fatalf("GameConfig.MapHeight is %d", state.GameConfig.MapHeight)
	}
	c.Assert(state.Bombs, DeepEquals, []Bomb{expectedBomb})
	col1 := []int{2, 2, 3}
	col2 := []int{0, 0, 1}
	col3 := []int{0, 0, 0}
	expectedBoard := [][]int{col1, col2, col3}
	c.Assert(state.Board, DeepEquals, expectedBoard)
}

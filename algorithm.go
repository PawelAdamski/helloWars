package main

import (
	"strconv"
	"strings"
)

type Algorithm interface {
	nextAction(GameState) BotMove
}

type Board [][]int

type GameState struct {
	Board              [][]int
	BotLocation        Location
	IsMissileAvailable bool
	OpponentLocations  []Location
	Bombs              []Bomb
	Missiles           []Missile
}

type Location struct {
	x int
	y int
}

func (l *Location) UnmarshalJSON(data []byte) error {

	s := string(data)
	s = strings.Replace(s, "\"", "", -1)
	ss := strings.Split(s, ",")
	var err error
	if l.x, err = strconv.Atoi(strings.TrimSpace(ss[0])); err != nil {
		return err
	}
	if l.y, err = strconv.Atoi(strings.TrimSpace(ss[1])); err != nil {
		return err
	}
	return nil
}

type Bomb struct {
	RoundsUntilExplodes int
	Location            Location
	ExplosionRadius     int
}

type Missile struct {
	MoveDirection   int
	Location        Location
	ExplosionRadius int
}

type BotMove struct {
	Direction     int
	Action        int
	FireDirection int
}

const Up = 0
const Down = 1
const Right = 2
const Left = 3

const None = 0
const DropBomb = 1
const FireMissile = 2

const Empty = 0
const Regular = 1
const Fortified = 2
const Indestructible = 3

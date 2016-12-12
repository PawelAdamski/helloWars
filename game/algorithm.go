package game

import (
	"strconv"
	"strings"
)

type Algorithm interface {
	nextAction(State) BotMove
}

// Board of the game
type Board [][]int

// State is received each turn
type State struct {
	Board              Board
	BotLocation        Location
	IsMissileAvailable bool
	OpponentLocations  []Location
	Bombs              []Bomb
	Missiles           []Missile
	GameConfig         Config
}

func (s *State) IsInside(l Location) bool {
	return l.X >= 0 && l.X < len(s.Board) && l.Y >= 0 && l.Y < len(s.Board[0])
}

func (s *State) IsEmpty(l Location) bool {
	return s.Board[l.X][l.Y] == Empty
}

func (s *State) CanMoveTo(l Location) bool {
	return s.IsInside(l) && s.IsEmpty(l)
}

// Config of the game
type Config struct {
	MapWidth                          int
	MapHeight                         int
	BombBlastRadius                   int
	MissileBlastRadius                int
	RoundsBetweenMissiles             int
	RoundsBeforeIncreasingBlastRadius int
	IsFastMissileModeEnabled          bool
}

// Location description
type Location struct {
	X int
	Y int
}

// Distance between points
func (l *Location) Distance(other *Location) int {
	distX := l.X - other.X
	if distX < 0 {
		distX = -distX
	}
	distY := l.Y - other.Y
	if distY {
		distY = -distY
	}
	if distX < distY {
		return distX
	}
	return distY
}

// Distance between points
func (l *Location) Neighbours(gs State) []Location {
	locs := []Location{}
	add := func(l Location) {
		if gs.IsInside(l) {
			locs = append(locs, l)
		}
	}
	add(Location{X: l.X, Y: l.Y + 1})
	add(Location{X: l.X, Y: l.Y - 1})
	add(Location{X: l.X + 1, Y: l.Y})
	add(Location{X: l.X - 1, Y: l.Y})
	return locs
}

// UnmarshalJSON implements json interface
func (l *Location) UnmarshalJSON(data []byte) error {

	s := string(data)
	s = strings.Replace(s, "\"", "", -1)
	ss := strings.Split(s, ",")
	var err error
	if l.X, err = strconv.Atoi(strings.TrimSpace(ss[0])); err != nil {
		return err
	}
	if l.Y, err = strconv.Atoi(strings.TrimSpace(ss[1])); err != nil {
		return err
	}
	return nil
}

// Bomb info
type Bomb struct {
	RoundsUntilExplodes int
	Location            Location
	ExplosionRadius     int
}

func (b *Bomb) IsInRadius(l Location) bool {
	return b.Location.Distance(l) <= b.ExplosionRadius
}

// Missile info
type Missile struct {
	MoveDirection   int
	Location        Location
	ExplosionRadius int
}

func (m *Missile) IsInRadius(l Location) bool {
	return m.Location.Distance(l) <= m.ExplosionRadius
}

// BotMove returned to handler
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

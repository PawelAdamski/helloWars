package game

import (
	"strconv"
	"strings"
)

var directions = []Direction{
	Direction{X: 0, Y: 1},
	Direction{X: 0, Y: -1},
	Direction{X: 1, Y: 0},
	Direction{X: -1, Y: 0},
}

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
	Bombs              Bombs
	Missiles           []Missile
	GameConfig         Config
}

func (s *State) IsInside(l *Location) bool {
	return l.X >= 0 && l.X < len(s.Board) && l.Y >= 0 && l.Y < len(s.Board[0])
}

func (s *State) IsEmpty(l *Location) bool {
	if !s.IsInside(l) {
		return false
	}
	return s.Board[l.X][l.Y] == Empty
}

func (s *State) CanMoveTo(l *Location) bool {
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

type Locations []Location

func (ll Locations) Contains(lo Location) bool {
	for _, l := range ll {
		if l == lo {
			return true
		}
	}
	return false
}

type Direction Location

// Distance between points
func (l *Location) Distance(other *Location) int {
	distX := l.X - other.X
	if distX < 0 {
		distX = -distX
	}
	distY := l.Y - other.Y
	if distY < 0 {
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
		if gs.IsInside(&l) {
			locs = append(locs, l)
		}
	}
	add(Location{X: l.X, Y: l.Y + 1})
	add(Location{X: l.X, Y: l.Y - 1})
	add(Location{X: l.X + 1, Y: l.Y})
	add(Location{X: l.X - 1, Y: l.Y})
	return locs
}

// Distance between points
func (l *Location) Moves(gs *State) []Location {
	locs := []Location{}
	add := func(l Location) {
		if gs.CanMoveTo(&l) {
			locs = append(locs, l)
		}
	}
	add(*l)
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
	return b.Location.Distance(&l) <= b.ExplosionRadius
}

type Bombs []Bomb

func (bs Bombs) decreaseCounters() {
	for i, _ := range bs {
		bs[i].RoundsUntilExplodes--
	}
}

func (bs Bombs) findExploding() (int, Bomb) {
	for i, b := range bs {
		if b.RoundsUntilExplodes == 1 {
			return i, b
		}
	}
	return -1, Bomb{}
}

func (bs Bombs) findChainedExplosions(l Location) {
	for i := range bs {
		if bs[i].Location == l {
			bs[i].RoundsUntilExplodes = 1
		}
	}
}

// Missile info
type Missile struct {
	MoveDirection   int
	Location        Location
	ExplosionRadius int
}

func (m *Missile) IsInRadius(l Location) bool {
	return m.Location.Distance(&l) <= m.ExplosionRadius
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

type ByLocation []Location

func (a ByLocation) Len() int      { return len(a) }
func (a ByLocation) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLocation) Less(i, j int) bool {
	if a[i].X != a[j].X {
		return a[i].X < a[j].X
	} else {
		return a[i].Y < a[j].Y
	}
}

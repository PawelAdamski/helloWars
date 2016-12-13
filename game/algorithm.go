package game

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var Directions = map[int]Direction{
	Down:  Direction{X: 0, Y: 1},
	Up:    Direction{X: 0, Y: -1},
	Right: Direction{X: 1, Y: 0},
	Left:  Direction{X: -1, Y: 0},
}

var inverseDirections = map[Direction]int{}

type Algorithm interface {
	nextAction(State) BotMove
}

// Board of the game
type Board [][]int

func (b Board) AfterExplosions(damagedWalls []Location) Board {
	c := make([][]int, len(b))
	for i, row := range b {
		c[i] = make([]int, len(row))
		copy(c[i], row)
	}
	for _, w := range damagedWalls {
		b.OnExplosion(&w)
	}
	return c
}

func (b Board) OnExplosion(l *Location) {
	t := b[l.X][l.Y]
	switch t {
	case Regular:
		b[l.X][l.Y] = Empty
	case Fortified:
		b[l.X][l.Y] = Regular
	}
}

// State is received each turn
type State struct {
	Board              Board
	BotID              string `json:"BotId"`
	BotLocation        Location
	IsMissileAvailable bool
	OpponentLocations  []Location
	Bombs              Bombs
	Missiles           Missiles
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

func (s *State) moveMissiles() {
	for i := range s.Missiles {
		d := Directions[s.Missiles[i].MoveDirection]
		nextLocation := s.Missiles[i].Location
		nextLocation.move(d)
		if s.collision(nextLocation) {
			s.Missiles[i].hasExploded = true
		} else {
			s.Missiles[i].Location = nextLocation
		}
		if s.GameConfig.IsFastMissileModeEnabled {
			nextLocation.move(d)
			if s.collision(nextLocation) {
				s.Missiles[i].hasExploded = true
			}
		}
	}
}

func (s *State) isBotLocation(l Location) bool {
	if s.BotLocation == l {
		return true
	}
	return Locations(s.OpponentLocations).Contains(l)
}

func (s *State) collision(l Location) bool {
	return !s.IsEmpty(&l) ||
		!s.IsInside(&l) ||
		s.Bombs.contains(l) ||
		s.isBotLocation(l) ||
		s.isBotLocation(l)
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

func joinLocationSet(dest, src map[Location]bool) {
	for l := range src {
		dest[l] = true
	}
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

func (ll Locations) MinDistance(lo Location) int {
	m := math.MaxInt32
	for _, l := range ll {
		d := l.Distance(&lo)
		if m > d {
			m = d
		}
	}
	return m
}

func (l *Location) move(d Direction) {
	l.X += d.X
	l.Y += d.Y
}

type Direction Location

func (d *Direction) AsResponse() *int {
	i, ok := inverseDirections[*d]
	if !ok {
		return nil
	}
	return &i
}

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
	return distY + distX
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

func (l Location) Translate(d Direction) Location {
	return Location{
		X: l.X + d.X,
		Y: l.Y + d.Y,
	}
}

// Distance between points
func (l *Location) Moves(gs *State) map[Direction]Location {
	dirLoc := map[Direction]Location{}
	add := func(d Direction) {
		t := l.Translate(d)
		if gs.CanMoveTo(&t) {
			dirLoc[d] = t
		}
	}
	add(Direction{X: 0, Y: 0})
	add(Direction{X: 1, Y: 0})
	add(Direction{X: -1, Y: 0})
	add(Direction{X: 0, Y: 1})
	add(Direction{X: 0, Y: -1})
	return dirLoc
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
		if b.RoundsUntilExplodes == 0 {
			return i, b
		}
	}
	return -1, Bomb{}
}

func (bs Bombs) findChainedExplosions(l Location) {
	for i := range bs {
		if bs[i].Location == l {
			bs[i].RoundsUntilExplodes = 0
		}
	}
}

func (bs Bombs) contains(l Location) bool {
	for _, b := range bs {
		if b.Location == l {
			return true
		}
	}
	return false
}

// Missile info
type Missile struct {
	MoveDirection   int
	Location        Location
	ExplosionRadius int
	hasExploded     bool
}

func (m *Missile) IsInRadius(l Location) bool {
	return m.Location.Distance(&l) <= m.ExplosionRadius
}

type Missiles []Missile

func (ms Missiles) findExploding() (int, Missile) {
	for i, m := range ms {
		if m.hasExploded {
			return i, m
		}
	}
	return -1, Missile{}
}

func (ms Missiles) findChainedExplosions(l Location) {
	for i := range ms {
		if ms[i].Location == l {
			ms[i].hasExploded = true
		}
	}
}

// BotMove returned to handler
type BotMove struct {
	Direction     *int
	Action        int
	FireDirection int
}

func (bm *BotMove) String() string {
	direction := "stay"
	d := Direction{}
	if bm.Direction != nil {
		direction = fmt.Sprint(*bm.Direction)
		d = Directions[*bm.Direction]
	}
	return fmt.Sprintf("Direction: %s (%v), action: %v, fire: %d",
		direction, d, bm.Action, bm.FireDirection)

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

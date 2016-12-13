package bf

import (
	"fmt"
	"math/rand"

	"github.com/PawelAdamski/helloWars/game"
)

const maxDepth = 6

type direction struct {
	direction   game.Direction
	canDropBomb bool
	missiles    []game.Direction
	actions     []action
}

// OutcomeType of move
type OutcomeType float64

const (
	OutcomeDead OutcomeType = -1
	OutcomeSafe OutcomeType = 1
)

type botMove struct {
	Direction        game.Direction
	Action           int
	Safety           int
	MissileDirection int
}

var longSearch = depth{
	me:       7,
	opponent: 3,
}

var shortSearch = depth{
	me:       6,
	opponent: 2,
}

func NextMove(s *game.State) game.BotMove {
	dirs := directions(s, s.BotLocation)
	if len(dirs) == 0 {
		return game.BotMove{}
	}
	r := 0.0
	argRand := game.BotMove{}
	for _, dir := range dirs {
		for _, a := range dir.actions {
			for _, o := range s.OpponentLocations {
				fmt.Println(dir.direction, a.action)
				if a.action != game.None && !isSafe(a.state, o, &s.BotLocation, shortSearch) {
					return a.toMove()
				}
				if f := rand.Float64(); r < f {
					r = f
					argRand = a.toMove()
				}
			}
		}
	}
	return argRand
}

func directions(gs *game.State, me game.Location) []direction {
	dirMap := map[game.Direction]*direction{}
	for _, a := range actions(me, gs, true) {
		if isSafeAgainstAll(a.state, a.nextLocation, gs.OpponentLocations) {
			if _, ok := dirMap[a.direction]; !ok {
				dirMap[a.direction] = &direction{
					direction:   a.direction,
					canDropBomb: false,
				}
			}
			dirMap[a.direction].actions = append(dirMap[a.direction].actions, a)
			if a.action == game.DropBomb {
				dirMap[a.direction].canDropBomb = true
			} else if a.action == game.FireMissile {
				dirMap[a.direction].missiles = append(
					dirMap[a.direction].missiles,
					a.missile)
			}
		}
	}
	dirs := []direction{}
	for _, d := range dirMap {
		dirs = append(dirs, *d)
	}
	return dirs
}

func isSafeAgainstAll(gs *game.State, me game.Location, os []game.Location) bool {
	if !isSafe(gs, me, nil, longSearch) {
		return false
	}
	for _, o := range os {
		if !isSafe(gs, me, &o, shortSearch) {
			return false
		}
	}
	return true
}

func isSafe(gs *game.State, me game.Location, o *game.Location, d depth) bool {
	nextGS, locs := gs.Next()
	if locs.Contains(me) {
		return false
	}
	if d.opponent == 0 {
		o = nil
	}
	if o != nil && locs.Contains(*o) {
		o = nil
	}
	if d.me > 0 {
		nd := d.next()
		for _, loc := range me.Moves(gs) {
			if safetyByOpponent(nextGS, loc, o, nd) {
				return true
			}
		}
		return false
	}
	return true
}

func safetyByOpponent(gs *game.State, me game.Location, o *game.Location, d depth) bool {
	if o == nil {
		return isSafe(gs, me, o, d)
	}
	for _, a := range actions(*o, gs, true) {
		if !isSafe(a.state, me, &a.nextLocation, d) {
			return false
		}
	}
	return true
}

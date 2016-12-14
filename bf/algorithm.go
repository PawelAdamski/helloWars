package bf

import (
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
	me:            8,
	opponent:      3,
	opponentFires: 1,
}

var shortSearch = depth{
	me:            5,
	opponent:      3,
	opponentFires: 1,
}

type Strategy struct{}

func (ss Strategy) NextAction(s *game.State) game.BotMove {
	dirs := relaxedDirections(*s, s.BotLocation)
	if len(dirs) == 0 {
		return game.BotMove{}
	}
	if mv, ok := findKillingMove(dirs, s); ok {
		return mv
	}
	if mv, ok := findSameDirectionMove(dirs, s); ok {
		return mv
	}
	return findAnyMove(dirs)
}

func findKillingMove(dirs []direction, s *game.State) (game.BotMove, bool) {
	for _, dir := range dirs {
		for _, a := range dir.actions {
			for _, o := range s.OpponentLocations {
				if a.action != game.None && !isSafe(a.state, o, &s.BotLocation, shortSearch) {
					return a.toMove(), true
				}

			}
		}
	}
	return game.BotMove{}, false
}

func findSameDirectionMove(dirs []direction, s *game.State) (game.BotMove, bool) {
	for _, dir := range dirs {
		for _, o := range s.OpponentLocations {
			if sameDirection(dir.direction, &s.BotLocation, &o) {
				return dir.actions[rand.Intn(len(dir.actions))].toMove(), true
			}
		}
	}
	return game.BotMove{}, false
}

func findAnyMove(dirs []direction) game.BotMove {
	as := []action{}
	for _, dir := range dirs {
		for _, a := range dir.actions {
			as = append(as, a)
		}
	}
	return as[rand.Intn(len(as))].toMove()
}

func sameDirection(d game.Direction, a, b *game.Location) bool {
	dx := b.X - a.X
	dy := b.Y - a.Y

	sdx := 0
	if dx < 0 {
		sdx = -1
	} else if dx > 0 {
		sdx = 1
	}

	sdy := 0
	if dy < 0 {
		sdy = -1
	} else if dx > 0 {
		sdy = 1
	}

	return d.X == sdx || d.Y == sdy
}

func relaxedDirections(gs game.State, me game.Location) []direction {
	dirs := directions(&gs, me)
	if len(dirs) != 0 {
		return dirs
	}
	gs.OpponentLocations = nil
	return directions(&gs, me)
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
		if !isSafe(gs, me, &o, longSearch) {
			return false
		}
	}
	return true
}

func isSafe(gs *game.State, me game.Location, o *game.Location, d depth) bool {
	if d.opponent == 0 {
		o = nil
	}
	gs.BotLocation = me
	if o != nil {
		gs.OpponentLocations = []game.Location{*o}
	}
	if o == nil {
		nextGS, locs := gs.Next()
		if locs.Contains(me) {
			return false
		}
		if d.me > 0 {
			nd := d.next()
			for _, loc := range me.Moves(nextGS) {
				if isSafe(nextGS, loc, nil, nd) {
					return true
				}
			}
			return false
		}
		return true
	}
	for _, a := range actions(*o, gs, true) {
		if !isActionSafe(a, me, d) {
			return false
		}
	}
	return true
}

func isActionSafe(a action, me game.Location, d depth) bool {
	nextGS, locs := a.state.Next()
	if locs.Contains(me) {
		return false
	}
	if d.me > 0 {
		nd := d.next()
		for _, loc := range me.Moves(nextGS) {
			if isSafe(nextGS, loc, nil, nd) {
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

	return true
}

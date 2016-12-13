package bf

import (
	"math"

	"github.com/PawelAdamski/helloWars/game"
)

type direction struct {
	direction game.Direction
	safety    int
}

// OutcomeType of move
type OutcomeType float64

const (
	OutcomeDead OutcomeType = -1
	OutcomeSafe OutcomeType = 1
)

type botMove struct {
	Direction game.Direction
	Action    int
	Safety    int
}

func maxSafety(bms []botMove) int {
	maxSafety := 0
	for _, bm := range bms {
		if maxSafety < bm.Safety {
			maxSafety = bm.Safety
		}
	}
	return maxSafety
}

func NextMove(s *game.State) game.BotMove {
	const bombsExplodeIn = 5
	moves := safeMoves(s, s.BotLocation, bombsExplodeIn)
	if len(moves) == 0 {
		return game.BotMove{}
	}
	safetyMin := math.MaxInt32
	argMin := botMove{}
	for _, move := range moves {
		for _, opponent := range s.OpponentLocations {
			bombInNewPlace := stateWithBomb(s, s.BotLocation.Translate(move.Direction), bombsExplodeIn+1)
			opponentMoves := safeMoves(bombInNewPlace, opponent, 1)
			opponentSafety := maxSafety(opponentMoves)
			if opponentSafety < safetyMin {
				safetyMin = opponentSafety
				argMin = move
			}
		}
	}

	return game.BotMove{
		Action:    argMin.Action,
		Direction: argMin.Direction.AsResponse(),
	}
}

func stateWithBomb(gs *game.State, loc game.Location, bombsExplodeIn int) *game.State {
	gsWithBombs := *gs
	gsWithBombs.Bombs = append([]game.Bomb{}, gs.Bombs...)
	gsWithBombs.Bombs = append(gsWithBombs.Bombs, game.Bomb{
		Location:            loc,
		ExplosionRadius:     gs.GameConfig.BombBlastRadius,
		RoundsUntilExplodes: bombsExplodeIn,
	})
	return &gsWithBombs
}

func safeMoves(gs *game.State, loc game.Location, bombsExplodeIn int) []botMove {
	gsWithBombs := stateWithBomb(gs, loc, bombsExplodeIn)
	dirs := directions(gsWithBombs, loc)
	if len(dirs) > 0 {
		return directionsToMoves(dirs, game.DropBomb)
	}
	return directionsToMoves(directions(gs, loc), game.None)
}

func directionsToMoves(dirs []direction, action int) []botMove {
	mvs := []botMove{}
	for _, dir := range dirs {
		mvs = append(mvs, botMove{
			Direction: dir.direction,
			Action:    action,
			Safety:    dir.safety,
		})
	}
	return mvs
}

func directions(gs *game.State, loc game.Location) []direction {
	const depth = 6
	directions := []direction{}
	for dir, move := range loc.Moves(gs) {
		if s := safety(gs, move, depth); s > 0 {
			directions = append(directions, direction{
				direction: dir,
				safety:    s,
			})
		}
	}
	return directions
}

func safety(gs *game.State, loc game.Location, depth int) int {
	nextGS, locs := gs.Next()
	explosionsSafety := locs.MinDistance(loc)
	if explosionsSafety == 0 {
		return explosionsSafety
	}
	if depth > 0 {
		recursiveSafety := 0
		for _, move := range loc.Moves(gs) {
			if s := safety(nextGS, move, depth-1); s > recursiveSafety {
				recursiveSafety = s
			}
		}
		if explosionsSafety < recursiveSafety {
			return explosionsSafety
		}
		return recursiveSafety
	}
	return explosionsSafety
}

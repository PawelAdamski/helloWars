package bf

import (
	"math/rand"

	"github.com/PawelAdamski/helloWars/game"
)

// OutcomeType of move
type OutcomeType float64

const (
	PlayerDead   OutcomeType = -1
	PlayerUnsafe OutcomeType = -0.5
	OpponentDead OutcomeType = 1
)

type botMove struct {
	Direction game.Direction
	Action    int
}

func NextMove(s *game.State) game.BotMove {
	moves := safeMoves(s, s.BotLocation, 5)
	if len(moves) == 0 {
		return game.BotMove{}
	}
	m := moves[rand.Intn(len(moves))]
	return game.BotMove{
		Action:    m.Action,
		Direction: m.Direction.AsResponse(),
	}
}

func safeMoves(gs *game.State, loc game.Location, bombsExplodeIn int) []botMove {
	gsWithBombs := *gs
	gsWithBombs.Bombs = append([]game.Bomb{}, gs.Bombs...)
	gsWithBombs.Bombs = append(gsWithBombs.Bombs, game.Bomb{
		Location:            loc,
		ExplosionRadius:     gs.GameConfig.BombBlastRadius,
		RoundsUntilExplodes: bombsExplodeIn,
	})

	dirs := directions(&gsWithBombs, loc)
	if len(dirs) > 0 {
		return directionsToMoves(dirs, game.DropBomb)
	}
	return directionsToMoves(directions(gs, loc), game.None)
}

func directionsToMoves(dirs []game.Direction, action int) []botMove {
	mvs := []botMove{}
	for _, dir := range dirs {
		mvs = append(mvs, botMove{
			Direction: dir,
			Action:    action,
		})
	}
	return mvs
}

func directions(gs *game.State, loc game.Location) []game.Direction {
	const depth = 6
	directions := []game.Direction{}
	for dir, move := range loc.Moves(gs) {
		if IsSafeFromBombs(gs, move, depth) {
			directions = append(directions, dir)
		}
	}
	return directions
}

func IsSafeFromBombs(gs *game.State, loc game.Location, depth int) bool {
	nextGS, locs := gs.Next()
	if locs.Contains(loc) {
		return false
	}
	if depth > 0 {
		for _, move := range loc.Moves(gs) {
			if IsSafeFromBombs(nextGS, move, depth-1) {
				return true
			}
		}
		return false
	}
	return true
}

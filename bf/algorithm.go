package bf

import (
	"fmt"

	"github.com/PawelAdamski/helloWars/game"
)

// OutcomeType of move
type OutcomeType float64

const (
	PlayerDead   OutcomeType = -1
	PlayerUnsafe OutcomeType = -0.5
	OpponentDead OutcomeType = 1
)

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

func isBombThreat(gs game.State, loc game.Location, time int) bool {
	for _, bomb := range gs.Bombs {
		if bomb.RoundsUntilExplodes == time && bomb.IsInRadius(loc) {
			fmt.Println("B", bomb.Location.X, bomb.Location.Y, time)
			return true
		}
	}
	return false
}

//
//func isMissileThreat(gs game.State, loc game.Location) bool {
//	for _, m := range gs.Missiles {
//		if !gs.IsEmpty(m.Location) && m.IsInRadius(loc) {
//			return true
//		}
//	}
//	return false
//}

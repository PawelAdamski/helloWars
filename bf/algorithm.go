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

func IsSafeFromBombs(gs game.State, loc game.Location, time int, depth int) bool {
	fmt.Println("Checking", loc.X, loc.Y, time, depth)
	if isBombThreat(gs, loc, time) {
		return false
	}
	if depth > 0 {
		fmt.Println("Moves", loc.Moves(gs))
		for _, move := range loc.Moves(gs) {
			if IsSafeFromBombs(gs, move, time+1, depth-1) {
				fmt.Println("Safe", move.X, move.Y, time, depth)
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

package bf

import (
	"github.com/PawelAdamski/helloWars/game"
	. "gopkg.in/check.v1"
)

func (s *BombSafeSuite) TestRegression1(c *C) {
	gs := game.State{
		BotID: "a6e5eaaa-6fe7-463e-9f16-314612246e4f",
		Board: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{2, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 2, 0, 3},
			{0, 0, 0, 0, 0, 1, 0, 3, 0, 0}},
		BotLocation:       game.Location{X: 4, Y: 4},
		OpponentLocations: []game.Location{{X: 3, Y: 6}},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: 1,
				Location:            game.Location{X: 3, Y: 4},
				ExplosionRadius:     3,
			},
			{
				RoundsUntilExplodes: 5,
				Location:            game.Location{X: 2, Y: 6},
				ExplosionRadius:     3,
			},
		},
		GameConfig: game.Config{
			MapWidth:                 5,
			MapHeight:                10,
			BombBlastRadius:          3,
			MissileBlastRadius:       2,
			RoundsBetweenMissiles:    5,
			IsFastMissileModeEnabled: true,
		}}
	moves := relaxedDirections(gs, game.Location{X: 4, Y: 4})
	c.Assert(moves, HasLen, 1)
}

package bf

import (
	"github.com/PawelAdamski/helloWars/game"
	. "gopkg.in/check.v1"
)

func (s *BombSafeSuite) TestSimpleMissiles(c *C) {
	//s.testSimpleMissiles(false, c)
	s.testSimpleMissiles(true, c)
}

func (s *BombSafeSuite) testSimpleMissiles(fastMissilesEnabled bool, c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Empty, game.Empty, game.Empty},
		},
		Missiles: []game.Missile{
			{
				MoveDirection:   game.Down,
				Location:        game.Location{X: 0, Y: 0},
				ExplosionRadius: 2,
			},
		},
		GameConfig: game.Config{
			IsFastMissileModeEnabled: fastMissilesEnabled,
		},
	}
	c.Assert(isSafe(&gs, game.Location{X: 0, Y: 2}, nil, longSearch), Equals, !fastMissilesEnabled)
}

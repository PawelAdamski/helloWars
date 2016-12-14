package bf

import (
	"github.com/PawelAdamski/helloWars/game"
	. "gopkg.in/check.v1"
)

func (s *BombSafeSuite) TestSimpleMissiles(c *C) {
	s.testSimpleMissiles(false, c)
	s.testSimpleMissiles(true, c)
}

func (s *BombSafeSuite) testSimpleMissiles(fastMissilesEnabled bool, c *C) {
	gs := game.State{
		BotLocation: game.Location{X: 0, Y: 2},
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

func (s *BombSafeSuite) TestMissileDirections(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Regular, game.Empty, game.Regular},
		},
		Missiles: []game.Missile{
			{
				MoveDirection:   game.Down,
				Location:        game.Location{X: 0, Y: 0},
				ExplosionRadius: 2,
			},
		},
		GameConfig: game.Config{
			IsFastMissileModeEnabled: true,
		},
	}
	dirs := directions(&gs, game.Location{X: 1, Y: 2})
	c.Assert(dirs, HasLen, 1)
	c.Assert(dirs[0].direction, Equals, game.Direction{X: 0, Y: 0})
}

func (s *BombSafeSuite) TestEscapingFromOpponentMissile(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Regular, game.Empty, game.Regular, game.Regular, game.Regular, game.Regular},
			{game.Empty, game.Empty, game.Empty, game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Empty, game.Empty, game.Regular, game.Empty, game.Regular},
		},
		GameConfig: game.Config{
			IsFastMissileModeEnabled: true,
			MissileBlastRadius:       3,
		},
		OpponentLocations: []game.Location{{X: 0, Y: 1}},
	}
	dirs := directions(&gs, game.Location{X: 0, Y: 4})
	c.Assert(dirs, HasLen, 1)
}

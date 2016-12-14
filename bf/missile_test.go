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

func (s *BombSafeSuite) TestMissileDeflection(c *C) {
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
			BombBlastRadius:          2,
		},
	}
	dirs := directions(&gs, game.Location{X: 0, Y: 2})
	c.Assert(dirs, HasLen, 1)
	c.Assert(dirs[0].direction, Equals, game.Direction{X: 1, Y: 0})
	c.Assert(dirs[0].canDropBomb, Equals, false)
}

func (s *BombSafeSuite) TestEscapingFromOpponentMissile(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Regular, game.Regular, game.Regular, game.Empty, game.Regular},
			{game.Empty, game.Empty, game.Empty, game.Empty, game.Empty},
		},
		GameConfig: game.Config{
			IsFastMissileModeEnabled: true,
			MissileBlastRadius:       3,
			BombBlastRadius:          2,
		},
		OpponentLocations: []game.Location{{X: 1, Y: 0}},
	}
	c.Assert(isSafe(&gs, game.Location{1, 3}, &game.Location{X: 1, Y: 0}, depth{me: 6, opponent: 2, opponentFires: 1}), Equals, false)
	//dirs := directions(&gs, game.Location{X: 1, Y: 3})
	//c.Assert(dirs, HasLen, 1)
	//c.Assert(dirs[0].direction, Equals, game.Direction{X: -1, Y: 0})
}

func (s *BombSafeSuite) TestDoesNotShootMissileAtSelf(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Empty, game.Empty},
		},
		GameConfig: game.Config{
			IsFastMissileModeEnabled: true,
			MissileBlastRadius:       5,
		},
	}
	dirs := directions(&gs, game.Location{X: 1, Y: 1})
	c.Assert(dirs, HasLen, 5)
	c.Assert(dirs[0].missiles, HasLen, 0)
	c.Assert(dirs[1].missiles, HasLen, 0)
	c.Assert(dirs[2].missiles, HasLen, 0)
	c.Assert(dirs[3].missiles, HasLen, 0)
	c.Assert(dirs[4].missiles, HasLen, 0)
}

func (s *BombSafeSuite) TestCanShootMissile(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty, game.Empty, game.Empty, game.Empty},
		},
		GameConfig: game.Config{
			IsFastMissileModeEnabled: true,
			MissileBlastRadius:       3,
		},
	}
	dirs := directions(&gs, game.Location{X: 0, Y: 0})
	c.Assert(dirs, HasLen, 2)
	c.Assert(dirs[0].missiles, HasLen, 1)
	c.Assert(dirs[1].missiles, HasLen, 1)
}

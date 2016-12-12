package bf

import (
	"testing"

	"github.com/PawelAdamski/helloWars/game"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BombSafeSuite struct{}

var _ = Suite(&BombSafeSuite{})

func (s *BombSafeSuite) TestSimpleCase(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty},
			{game.Empty, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: 1,
				Location: game.Location{
					X: 0,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
		},
	}
	c.Assert(IsSafeFromBombs(&gs, game.Location{X: 0, Y: 1}, 0), Equals, false)
}

func (s *BombSafeSuite) TestIncomingExplosion(c *C) {
	gs := &game.State{
		Board: [][]int{
			{game.Empty, game.Empty},
			{game.Empty, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: 3,
				Location: game.Location{
					X: 0,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
			{
				RoundsUntilExplodes: 3,
				Location: game.Location{
					X: 1,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
		},
	}
	c.Assert(IsSafeFromBombs(gs, game.Location{X: 0, Y: 0}, 0), Equals, true)
	c.Assert(IsSafeFromBombs(gs, game.Location{X: 0, Y: 0}, 1), Equals, true)
	c.Assert(IsSafeFromBombs(gs, game.Location{X: 0, Y: 0}, 2), Equals, false)
}

func (s *BombSafeSuite) TestAvoidableExplosion(c *C) {
	s.testAvoidableExplosion(c, 1, false)
	s.testAvoidableExplosion(c, 2, false)
	s.testAvoidableExplosion(c, 3, true)
}

func (s *BombSafeSuite) testAvoidableExplosion(c *C, roundsUntilExplodes int, avoidable bool) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty},
			{game.Empty, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: roundsUntilExplodes,
				Location: game.Location{
					X: 0,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
		},
	}
	c.Assert(IsSafeFromBombs(&gs, game.Location{X: 0, Y: 0}, 7), Equals, avoidable)
}

func (s *BombSafeSuite) TestCanHideFromExplosion(c *C) {
	roundsUntilExplodes := 4
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Indestructible, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: roundsUntilExplodes,
				Location: game.Location{
					X: 0,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
			{
				RoundsUntilExplodes: roundsUntilExplodes,
				Location: game.Location{
					X: 1,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
		},
	}
	c.Assert(IsSafeFromBombs(&gs, game.Location{X: 0, Y: 0}, 7), Equals, true)
}

func (s *BombSafeSuite) TestRunsForCover(c *C) {
	roundsUntilExplodes := 5
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Indestructible, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: roundsUntilExplodes,
				Location: game.Location{
					X: 0,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
			{
				RoundsUntilExplodes: roundsUntilExplodes,
				Location: game.Location{
					X: 1,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
		},
	}
	dirs := directions(&gs, game.Location{X: 1, Y: 0})
	c.Assert(dirs, HasLen, 1)
	c.Assert(dirs[0], Equals, game.Direction{X: -1, Y: 0})
}

func (s *BombSafeSuite) TestTriesToSetBombs(c *C) {
	roundsUntilExplodes := 5
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Indestructible, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: roundsUntilExplodes,
				Location: game.Location{
					X: 0,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
		},
		GameConfig: game.Config{
			BombBlastRadius: 10,
		},
	}
	moves := Moves(&gs, game.Location{X: 1, Y: 0}, 5)
	c.Assert(moves, HasLen, 1)
	c.Assert(moves[0].Action, Equals, game.DropBomb)
}

func (s *BombSafeSuite) TestSettingBombUnsafe(c *C) {
	roundsUntilExplodes := 2
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty},
			{game.Empty, game.Indestructible, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: roundsUntilExplodes,
				Location: game.Location{
					X: 0,
					Y: 2,
				},
				ExplosionRadius: 2,
			},
		},
		GameConfig: game.Config{
			BombBlastRadius: 10,
		},
	}
	moves := Moves(&gs, game.Location{X: 0, Y: 0}, 2)
	c.Assert(moves, HasLen, 1)
	c.Assert(moves[0].Direction, Equals, game.Direction{X: 1, Y: 0})
	c.Assert(moves[0].Action, Equals, game.None)
}

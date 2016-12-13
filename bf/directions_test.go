package bf

import (
	"github.com/PawelAdamski/helloWars/game"
	. "gopkg.in/check.v1"
)

func (s *BombSafeSuite) TestRunsForCover1(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty},
			{game.Empty, game.Empty},
		},
		Bombs: []game.Bomb{
			{
				RoundsUntilExplodes: 2,
				Location: game.Location{
					X: 0,
					Y: 0,
				},
				ExplosionRadius: 2,
			},
		},
	}
	dirs := directions(&gs, game.Location{X: 0, Y: 0})
	c.Assert(dirs, HasLen, 2)
}

func (s *BombSafeSuite) TestRunsForCover2(c *C) {
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
	dirs := directions(&gs, game.Location{X: 1, Y: 0})
	c.Assert(dirs, HasLen, 1)
}

func (s *BombSafeSuite) TestRunsForCover3(c *C) {
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
	dirs := directions(&gs, game.Location{X: 1, Y: 0})
	c.Assert(dirs, HasLen, 1)
	c.Assert(dirs[0].direction, Equals, game.Direction{X: -1, Y: 0})
}

func (s *BombSafeSuite) TestTriesToSetBombs(c *C) {
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
		},
		GameConfig: game.Config{
			BombBlastRadius: 10,
		},
	}
	moves := directions(&gs, game.Location{X: 1, Y: 0})
	c.Assert(moves, HasLen, 1)
	c.Assert(moves[0].canDropBomb, Equals, true)
}

func (s *BombSafeSuite) TestSettingBombUnsafe(c *C) {
	roundsUntilExplodes := 1
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
	moves := directions(&gs, game.Location{X: 0, Y: 0})
	c.Assert(moves, HasLen, 1)
	c.Assert(moves[0].direction, Equals, game.Direction{X: 1, Y: 0})
	c.Assert(moves[0].canDropBomb, Equals, false)
}

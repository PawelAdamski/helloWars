package bf

import (
	"testing"

	"math"

	"github.com/PawelAdamski/helloWars/game"
	. "gopkg.in/check.v1"
)

func assertAvoidable(s int, avoidable bool, c *C) {
	if avoidable {
		if s == 0 {
			c.Fatalf("Expected avoidable, got %d", s)
		}
	} else {
		if s != 0 {
			c.Fatalf("Expected not avoidable, got %d", s)
		}
	}
}

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
	c.Assert(safety(&gs, game.Location{X: 0, Y: 1}, 0), Equals, 0)
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
	c.Assert(safety(gs, game.Location{X: 0, Y: 0}, 0), Equals, math.MaxInt32)
	c.Assert(safety(gs, game.Location{X: 0, Y: 0}, 1), Equals, math.MaxInt32)
	c.Assert(safety(gs, game.Location{X: 0, Y: 0}, 2), Equals, 0)
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
	assertAvoidable(safety(&gs, game.Location{X: 0, Y: 0}, 7), avoidable, c)
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
	assertAvoidable(safety(&gs, game.Location{X: 0, Y: 0}, 7), true, c)
}

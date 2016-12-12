package game

import (
	"testing"

	"sort"

	. "gopkg.in/check.v1"
)

func TestNextState(t *testing.T) { TestingT(t) }

type NextStateSuite struct{}

var _ = Suite(&NextStateSuite{})

func (s *NextStateSuite) TestNextMovesMissiles(c *C) {
}

func (s *NextStateSuite) TestMissileHitsWall(c *C) {
}

func (s *NextStateSuite) TestBombExplosion(c *C) {
	state := State{
		Board: emptyBoard(7, 7),
		Bombs: []Bomb{
			Bomb{
				RoundsUntilExplodes: 1,
				ExplosionRadius:     2,
				Location:            Location{X: 3, Y: 3}},
		}}
	nextState, explosions := state.Next()

	expectedExplosions := []Location{
		Location{X: 3, Y: 3},
		Location{X: 1, Y: 3},
		Location{X: 2, Y: 3},
		Location{X: 4, Y: 3},
		Location{X: 5, Y: 3},
		Location{X: 3, Y: 1},
		Location{X: 3, Y: 2},
		Location{X: 3, Y: 4},
		Location{X: 3, Y: 5},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
	c.Assert(len(nextState.Bombs), Equals, 0)
}

func (s *NextStateSuite) TestMissisleExplosion(c *C) {
}

func (s *NextStateSuite) TestChainedExplosion(c *C) {
	state := State{
		Board: emptyBoard(7, 7),
		Bombs: []Bomb{
			Bomb{
				RoundsUntilExplodes: 1,
				ExplosionRadius:     1,
				Location:            Location{X: 3, Y: 3}},
			Bomb{
				RoundsUntilExplodes: 10,
				ExplosionRadius:     1,
				Location:            Location{X: 2, Y: 3}},
			Bomb{
				RoundsUntilExplodes: 10,
				ExplosionRadius:     1,
				Location:            Location{X: 1, Y: 3}},
		}}
	nextState, explosions := state.Next()

	expectedExplosions := []Location{
		Location{X: 4, Y: 3},
		Location{X: 3, Y: 3},
		Location{X: 2, Y: 3},
		Location{X: 1, Y: 3},

		Location{X: 3, Y: 4},
		Location{X: 2, Y: 4},
		Location{X: 1, Y: 4},

		Location{X: 3, Y: 5},
		Location{X: 2, Y: 5},
		Location{X: 1, Y: 5},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
	c.Assert(len(nextState.Bombs), Equals, 0)
}

func emptyBoard(width, height int) Board {
	b := Board{}
	for _ := range width {
		b = append(b, make([]int, height))
	}
	return b
}

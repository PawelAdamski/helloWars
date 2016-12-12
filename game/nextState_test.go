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
		Board: board(7, 7),
		Bombs: []Bomb{
			Bomb{
				RoundsUntilExplodes: 1,
				ExplosionRadius:     2,
				Location:            Location{X: 3, Y: 3}},
		}}
	nextState, explosions := state.Next()

	expectedExplosions := Locations{
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
	state := State{
		Board: board(7, 7, Location{X: 3, Y: 3}),
		Missiles: []Missile{Missile{
			Location:        Location{X: 2, Y: 3},
			MoveDirection:   Right,
			ExplosionRadius: 2,
		}},
	}
	nextState, explosions := state.Next()

	expectedExplosions := Locations{
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

func (s *NextStateSuite) TestChainedExplosion(c *C) {
	state := State{
		Board: board(7, 7),
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

	expectedExplosions := Locations{
		Location{X: 4, Y: 3},
		Location{X: 3, Y: 3},
		Location{X: 2, Y: 3},
		Location{X: 1, Y: 3},
		Location{X: 0, Y: 3},

		Location{X: 3, Y: 4},
		Location{X: 2, Y: 4},
		Location{X: 1, Y: 4},

		Location{X: 3, Y: 2},
		Location{X: 2, Y: 2},
		Location{X: 1, Y: 2},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
	c.Assert(len(nextState.Bombs), Equals, 0)
}

func board(width, height int, walls ...Location) Board {
	b := Board{}
	for i := 0; i < width; i++ {
		b = append(b, make([]int, height))
	}
	for _, w := range walls {
		b[w.Y][w.X] = 1
	}
	return b
}
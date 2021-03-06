package game

import (
	"sort"

	. "gopkg.in/check.v1"
)

type NextStateSuite struct{}

var _ = Suite(&NextStateSuite{})

func (s *NextStateSuite) TestMissileHitsBomb(c *C) {
	state := State{
		Board: board(7, 7),
		Bombs: []Bomb{
			Bomb{
				RoundsUntilExplodes: 10,
				ExplosionRadius:     1,
				Location:            Location{X: 3, Y: 3}},
		},
		Missiles: Missiles{
			Missile{
				ExplosionRadius: 1,
				Location:        Location{X: 3, Y: 6},
				MoveDirection:   Up,
			},
		}}
	nextState, _ := state.Next()
	nextState, _ = nextState.Next()
	nextState, explosions := nextState.Next()

	expectedExplosions := Locations{
		Location{X: 3, Y: 3},
		Location{X: 2, Y: 3},
		Location{X: 3, Y: 2},
		Location{X: 4, Y: 3},
		Location{X: 3, Y: 4},
		Location{X: 3, Y: 5},
		Location{X: 2, Y: 4},
		Location{X: 4, Y: 4},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
	c.Assert(len(nextState.Bombs), Equals, 0)
	c.Assert(len(nextState.Missiles), Equals, 0)
}

func (s *NextStateSuite) TestMissileGoesOutOfBoard(c *C) {
	state := State{
		Board: board(7, 7),
		Missiles: []Missile{Missile{
			Location:        Location{X: 3, Y: 3},
			MoveDirection:   Right,
			ExplosionRadius: 2,
		}},
	}
	nextState, explosions := state.Next()
	c.Assert(nextState.Missiles[0].Location, DeepEquals, Location{X: 4, Y: 3})
	c.Assert(len(explosions), Equals, 0)

	nextState, explosions = nextState.Next()
	c.Assert(nextState.Missiles[0].Location, DeepEquals, Location{X: 5, Y: 3})
	c.Assert(len(explosions), Equals, 0)

	nextState, explosions = nextState.Next()
	c.Assert(nextState.Missiles[0].Location, DeepEquals, Location{X: 6, Y: 3})
	c.Assert(len(explosions), Equals, 0)

	nextState, explosions = nextState.Next()
	c.Assert(len(nextState.Missiles), Equals, 0)

	expectedExplosions := Locations{
		Location{X: 6, Y: 3},
		Location{X: 5, Y: 3},
		Location{X: 4, Y: 3},
		Location{X: 6, Y: 4},
		Location{X: 6, Y: 5},
		Location{X: 6, Y: 1},
		Location{X: 6, Y: 2},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
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

func (s *NextStateSuite) TestMissileExplosion(c *C) {
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
		Location{X: 0, Y: 3},
		Location{X: 1, Y: 3},
		Location{X: 2, Y: 3},
		Location{X: 2, Y: 4},
		Location{X: 2, Y: 5},
		Location{X: 2, Y: 2},
		Location{X: 2, Y: 1},
		Location{X: 3, Y: 3},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
	c.Assert(len(nextState.Bombs), Equals, 0)
}

func (s *NextStateSuite) TestChainedBombExplosion(c *C) {
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

func (s *NextStateSuite) TestChainedMissileExplosion(c *C) {
	state := State{
		Board: board(7, 7),
		Bombs: []Bomb{
			Bomb{
				RoundsUntilExplodes: 1,
				ExplosionRadius:     1,
				Location:            Location{X: 3, Y: 3}},
		},
		Missiles: Missiles{
			Missile{
				ExplosionRadius: 1,
				Location:        Location{X: 4, Y: 4},
				MoveDirection:   Left,
			},
		}}
	nextState, explosions := state.Next()

	expectedExplosions := Locations{
		Location{X: 3, Y: 3},
		Location{X: 2, Y: 3},
		Location{X: 3, Y: 2},
		Location{X: 4, Y: 3},
		Location{X: 3, Y: 4},
		Location{X: 3, Y: 5},
		Location{X: 2, Y: 4},
		Location{X: 4, Y: 4},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
	c.Assert(len(nextState.Bombs), Equals, 0)
	c.Assert(len(nextState.Missiles), Equals, 0)
}

func (s *NextStateSuite) TestMissilesMove(c *C) {
	state := State{
		Board: board(7, 7),
		Missiles: Missiles{
			Missile{
				ExplosionRadius: 1,
				Location:        Location{X: 0, Y: 0},
				MoveDirection:   Down,
			},
			Missile{
				ExplosionRadius: 1,
				Location:        Location{X: 1, Y: 1},
				MoveDirection:   Down,
			},
		}}
	nextState, explosions := state.Next()
	c.Assert(explosions, HasLen, 0)
	c.Assert(nextState.Missiles, HasLen, 2)

	c.Assert(nextState.Missiles[0].Location, Equals, Location{X: 0, Y: 1})
	c.Assert(nextState.Missiles[1].Location, Equals, Location{X: 1, Y: 2})

	c.Assert(state.Missiles[0].Location, Equals, Location{X: 0, Y: 0})
	c.Assert(state.Missiles[1].Location, Equals, Location{X: 1, Y: 1})
}

func (s *NextStateSuite) TestChainedMisslesAndBombsWithFastModeExplosion(c *C) {
	state := State{
		Board: board(7, 7, Location{X: 6, Y: 2}),
		Bombs: []Bomb{
			Bomb{
				RoundsUntilExplodes: 10,
				ExplosionRadius:     2,
				Location:            Location{X: 5, Y: 4}},
		},
		Missiles: Missiles{
			Missile{
				ExplosionRadius: 2,
				Location:        Location{X: 3, Y: 2},
				MoveDirection:   Right,
			},
			Missile{
				ExplosionRadius: 2,
				Location:        Location{X: 1, Y: 4},
				MoveDirection:   Right,
			},
		},
		GameConfig: Config{IsFastMissileModeEnabled: true},
	}
	nextState, _ := state.Next()
	nextState, explosions := nextState.Next()

	expectedExplosions := Locations{
		Location{X: 1, Y: 4},

		Location{X: 2, Y: 4},

		Location{X: 3, Y: 2},
		Location{X: 3, Y: 3},
		Location{X: 3, Y: 4},
		Location{X: 3, Y: 5},
		Location{X: 3, Y: 6},

		Location{X: 4, Y: 2},
		Location{X: 4, Y: 4},

		Location{X: 5, Y: 0},
		Location{X: 5, Y: 1},
		Location{X: 5, Y: 2},
		Location{X: 5, Y: 3},
		Location{X: 5, Y: 4},
		Location{X: 5, Y: 5},
		Location{X: 5, Y: 6},

		Location{X: 6, Y: 2},
		Location{X: 6, Y: 4},
	}
	sort.Sort(ByLocation(expectedExplosions))
	sort.Sort(ByLocation(explosions))
	c.Assert(explosions, DeepEquals, expectedExplosions)
	c.Assert(len(nextState.Bombs), Equals, 0)
	c.Assert(len(nextState.Missiles), Equals, 0)
}

func board(width, height int, walls ...Location) Board {
	b := Board{}
	for i := 0; i < width; i++ {
		b = append(b, make([]int, height))
	}
	for _, w := range walls {
		b[w.X][w.Y] = 1
	}
	return b
}

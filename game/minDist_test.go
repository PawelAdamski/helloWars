package game

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestMinDist(t *testing.T) { TestingT(t) }

type MinDistSuite struct{}

var _ = Suite(&MinDistSuite{})

func (s *MinDistSuite) TestMinDist(c *C) {
	state := State{
		Board:      board(4, 4, Location{1, 0}),
		GameConfig: Config{MapWidth: 4, MapHeight: 4},
	}
	dist := state.minDistanceMap(Location{0, 0})
	col1 := []int{0, 1, 2, 3}
	col2 := []int{100000, 2, 3, 4}
	col3 := []int{4, 3, 4, 5}
	col4 := []int{5, 4, 5, 6}

	expected := [][]int{col1, col2, col3, col4}
	c.Assert(dist, DeepEquals, expected)
}

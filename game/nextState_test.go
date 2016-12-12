package game

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestNextState(t *testing.T) { TestingT(t) }

type NextStateSuite struct{}

var _ = Suite(&NextStateSuite{})


func (s *NextStateSuite) TestNextDescreaseBombCounters(c *C) {
}

func (s *NextStateSuite) TestNextMovesMissiles(c *C) {
}

func (s *NextStateSuite) TestBombExplosion(c *C) {
}

func (s *NextStateSuite) TestMissisleExplosion(c *C) {
}

func (s *NextStateSuite) TestChainedExplosion(c *C) {
}

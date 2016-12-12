package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type DeathMapSuite struct{}

var _ = Suite(&DeathMapSuite{})

func (s *DeathMapSuite) TestBombExplosion(c *C) {
	//board := [][]int {
	//	[]int {0,1,0},
	//	[]int {2,0,2},
	//	[]int {0,1,0},
	//}
	//bomb = game.Bomb{
	//	Location:game.Location{X:1,Y:1},
	//	ExplosionRadius:
	//
	//}
}

func (s *DeathMapSuite) TestMissileExplosion(c *C) {

}

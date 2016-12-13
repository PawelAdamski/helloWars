package bf

import (
	"github.com/PawelAdamski/helloWars/game"
	. "gopkg.in/check.v1"
)

func (s *BombSafeSuite) TestSmartOpponentDropsBomb(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty, game.Empty, game.Empty},
			{game.Indestructible, game.Indestructible, game.Indestructible, game.Indestructible, game.Empty},
		},
		Bombs: []game.Bomb{},
		GameConfig: game.Config{
			BombBlastRadius: 10,
		},
	}
	c.Assert(isSafe(&gs, game.Location{X: 0, Y: 0}, &game.Location{X: 0, Y: 3}, shortSearch), Equals, false)
}

func (s *BombSafeSuite) TestSmartFiresMissile(c *C) {
	gs := game.State{
		Board: [][]int{
			{game.Empty, game.Empty, game.Empty, game.Empty, game.Empty, game.Empty, game.Empty},
		},
		Bombs: []game.Bomb{},
		GameConfig: game.Config{
			MissileBlastRadius: 5,
		},
	}
	c.Assert(isSafe(&gs, game.Location{X: 0, Y: 0}, &game.Location{X: 0, Y: 3}, shortSearch), Equals, false)
}

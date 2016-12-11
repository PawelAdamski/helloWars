package main

import "github.com/PawelAdamski/helloWars/game"

// DeathMap hold information if according to current game state there will explosion on each field. If yes how many round from now.
type DeathMap [][]int

func mkDeathMap(board game.Board, bombs []game.Bomb, missiles []game.Missile) DeathMap {

	return DeathMap{}
}

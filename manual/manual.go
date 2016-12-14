package manual

import (
	"bufio"
	"os"

	"fmt"

	"github.com/PawelAdamski/helloWars/game"
)

type Strategy struct{}

var wsad = map[string]int{
	"w": game.Up,
	"s": game.Down,
	"a": game.Left,
	"d": game.Right,
}

func (ss Strategy) NextAction(s *game.State) game.BotMove {
	move := game.BotMove{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Direction: ")
	direction, err := reader.ReadString('\n')
	direction = direction[:1]
	if err != nil {
		return move
	}
	d := wsad[direction]
	move.Direction = &d

	fmt.Print("Action: ")
	action, err := reader.ReadString('\n')
	action = action[:1]
	if err != nil {
		return move
	}
	if action == "b" {
		move.Action = game.DropBomb
	} else if action == "n" {
		move.Action = game.None
	} else if action == "n" {
		move.Action = game.None
	} else if action == "i" {
		move.Action = game.FireMissile
		move.FireDirection = game.Up
	} else if action == "j" {
		move.Action = game.FireMissile
		move.FireDirection = game.Left
	} else if action == "k" {
		move.Action = game.FireMissile
		move.FireDirection = game.Down
	} else if action == "l" {
		move.Action = game.FireMissile
		move.FireDirection = game.Right
	}

	return move
}

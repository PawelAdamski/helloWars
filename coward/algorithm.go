package coward

import "github.com/PawelAdamski/helloWars/game"

var moves = map[string]int{}

func NextMove(s *game.State) game.BotMove {

	_, explosions := s.Next()
	var botDirection *int
	for di, d := range game.Directions {
		nextPosition := s.BotLocation
		nextPosition.Translate(d)
		if s.IsEmpty(&nextPosition) && !explosions.Contains(nextPosition) {
			dic := di
			botDirection = &dic
			break
		}
	}

	ms := moves[s.BotID]
	moves[s.BotID] = ms + 1
	action := game.None
	if ms%2 == 0 {
		action = game.DropBomb
	}

	return game.BotMove{
		Action:    action,
		Direction: botDirection,
	}
}

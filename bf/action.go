package bf

import "github.com/PawelAdamski/helloWars/game"

const bombsExplodeIn = 5

type action struct {
	direction    game.Direction
	nextLocation game.Location
	state        *game.State
	action       int
	missile      game.Direction
}

func (a *action) toMove() game.BotMove {
	fd := a.missile.AsResponse()
	bm := game.BotMove{
		Direction: a.direction.AsResponse(),
		Action:    a.action,
	}
	if fd != nil {
		bm.FireDirection = *fd
	}
	return bm
}

func stateWithBomb(gs *game.State, loc game.Location, bombsExplodeIn int) *game.State {
	gsWithBombs := *gs
	gsWithBombs.Bombs = append([]game.Bomb{}, gs.Bombs...)
	gsWithBombs.Bombs = append(gsWithBombs.Bombs, game.Bomb{
		Location:            loc,
		ExplosionRadius:     gs.GameConfig.BombBlastRadius,
		RoundsUntilExplodes: bombsExplodeIn,
	})
	return &gsWithBombs
}

func stateWithMissile(gs *game.State, loc game.Location, d game.Direction) *game.State {
	gsWithMissiles := *gs
	gsWithMissiles.Missiles = append([]game.Missile{}, gs.Missiles...)
	gsWithMissiles.Missiles = append(gsWithMissiles.Missiles, game.Missile{
		Location:        loc.Translate(d),
		ExplosionRadius: gs.GameConfig.MissileBlastRadius,
		MoveDirection:   *d.AsResponse(),
	})
	return &gsWithMissiles
}

func actions(l game.Location, s *game.State, checkMissiles bool) []action {
	actions := []action{}
	moves := l.Moves(s)
	for direction, loc := range moves {
		actions = append(actions,
			action{
				direction:    direction,
				nextLocation: loc,
				action:       game.None,
				state:        s,
			},
			action{
				direction:    direction,
				nextLocation: loc,
				action:       game.DropBomb,
				state:        stateWithBomb(s, l, bombsExplodeIn),
			},
		)
		if checkMissiles {
			for _, missleDir := range game.Directions {
				if missleDir.X != 0 || missleDir.Y != 0 {
					actions = append(actions,
						action{
							direction:    direction,
							nextLocation: loc,
							action:       game.FireMissile,
							missile:      missleDir,
							state:        stateWithMissile(s, loc, missleDir),
						})
				}

			}
		}

	}
	return actions
}

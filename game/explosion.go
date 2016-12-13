package game

type explosionT struct {
	location Location
	radius   int
}

func (s *State) Next() (*State, Locations) {

	nextState := *s
	nextState.Bombs = append([]Bomb{}, nextState.Bombs...)
	nextState.Missiles = append([]Missile{}, nextState.Missiles...)

	damagedWalls := []Location{}
	explosions := map[Location]bool{}
	nextState.moveMissiles()
	nextState.Bombs.decreaseCounters()
	for {
		var explosion *explosionT
		if bi, b := nextState.Bombs.findExploding(); bi >= 0 {
			nextState.Bombs = append(nextState.Bombs[:bi], nextState.Bombs[bi+1:]...)
			explosion = &explosionT{location: b.Location, radius: b.ExplosionRadius}
		}
		if mi, m := nextState.Missiles.findExploding(); explosion == nil && mi >= 0 {
			nextState.Missiles = append(nextState.Missiles[:mi], nextState.Missiles[mi+1:]...)
			explosion = &explosionT{location: m.Location, radius: m.ExplosionRadius}
		}
		if explosion == nil {
			return &nextState, toLocationSlice(explosions)
		}
		exp, dw := nextState.explode(explosion)
		joinLocationSet(explosions, exp)
		damagedWalls = append(damagedWalls, dw...)

	}
	if len(damagedWalls) > 0 {
		nextState.Board = s.Board.AfterExplosions(damagedWalls)
	}
	return nil, []Location{}
}

func (s *State) explode(e *explosionT) ( map[Location]bool, []Location) {
	explosions := map[Location]bool{}
	damagedWalls := []Location{}
	for _, d := range Directions {
		l := e.location
		explosions[l] = true
		for i := 0; i < e.radius; i++ {
			l.move(d)
			if s.IsInside(&l) {
				if s.IsEmpty(&l) {
					explosions[l] = true
					s.Bombs.findChainedExplosions(l)
					s.Missiles.findChainedExplosions(l)
				} else {
					damagedWalls = append(damagedWalls, l)
					break
				}
			}
		}
	}
	return explosions, damagedWalls
}

func toLocationSlice(set map[Location]bool) Locations {
	ls := []Location{}
	for l := range set {
		ls = append(ls, l)
	}
	return Locations(ls)
}

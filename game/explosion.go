package game

type explosionT struct {
	location Location
	radius   int
}

func (s *State) Next() (*State, Locations) {

	nextState := *s
	nextState.Bombs = append([]Bomb{}, nextState.Bombs...)
	nextState.Missiles = append([]Missile{}, nextState.Missiles...)

	explosions := map[Location]bool{}
	explodingMissiles := nextState.moveMissiles()
	nextState.Bombs.decreaseCounters()
	damagedWalls := []Location{}
	for {
		var explosion *explosionT
		if bi, b := nextState.Bombs.findExploding(); bi >= 0 {
			nextState.Bombs = append(nextState.Bombs[:bi], nextState.Bombs[bi+1:]...)
			explosion = &explosionT{location: b.Location, radius: b.ExplosionRadius}
		}
		if explosion == nil && len(explodingMissiles) > 0 {
			m := explodingMissiles[0]
			explosion = &explosionT{location: m.Location, radius: m.ExplosionRadius}
			explodingMissiles = explodingMissiles[1:]
		}
		if explosion == nil {
			return &nextState, toLocationSlice(explosions)
		}
		for _, d := range Directions {
			l := explosion.location
			explosions[l] = true
			for i := 0; i < explosion.radius; i++ {
				l.move(d)
				if nextState.IsInside(&l) {
					if nextState.IsEmpty(&l) {
						explosions[l] = true
						nextState.Bombs.findChainedExplosions(l)
					} else {
						damagedWalls = append(damagedWalls, l)
						break
					}
				}
			}
		}
	}
	if len(damagedWalls) > 0 {
		nextState.Board = s.Board.AfterExplosions(damagedWalls)
	}
	return nil, []Location{}
}

func toLocationSlice(set map[Location]bool) Locations {
	ls := []Location{}
	for l := range set {
		ls = append(ls, l)
	}
	return Locations(ls)
}

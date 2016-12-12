package game

func (s *State) Next() (*State, Locations) {

	nextState := *s
	nextState.Bombs = append([]Bomb{}, nextState.Bombs...)

	explosions := map[Location]bool{}
	explodingMissiles := nextState.moveMissiles()
	damagedWalls := []Location{}
	for {
		anyExplosion := false
		explosionLocation := Location{}
		explosionRadius := 0
		bi, b := nextState.Bombs.findExploding()
		if bi >= 0 {
			nextState.Bombs = append(nextState.Bombs[:bi], nextState.Bombs[bi+1:]...)
			anyExplosion = true
			explosionLocation = b.Location
			explosionRadius = b.ExplosionRadius
		}
		if !anyExplosion && len(explodingMissiles) > 0 {
			anyExplosion = true
			explosionLocation = explodingMissiles[0].Location
			explosionRadius = explodingMissiles[0].ExplosionRadius
			explodingMissiles = explodingMissiles[1:]
		}
		if !anyExplosion {
			nextState.Bombs.decreaseCounters()
			return &nextState, toLocationSlice(explosions)
		}
		for _, d := range Directions {
			l := explosionLocation
			explosions[l] = true
			for i := 0; i < explosionRadius; i++ {
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
		nextState.Board = s.Board.Clone()
		for _, w := range damagedWalls {
			s.Board.OnExplosion(&w)
		}
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

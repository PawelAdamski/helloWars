package game

func (s *State) Next() (*State, Locations) {

	nextState := *s
	nextState.Bombs = append([]Bomb{}, nextState.Bombs...)

	// avoid cloning if not needed
	boardCloned := false
	nextState.Board = s.Board.Clone()

	explosions := map[Location]bool{}
	for bi, b := nextState.Bombs.findExploding(); bi >= 0; bi, b = nextState.Bombs.findExploding() {
		nextState.Bombs = append(
			append([]Bomb{}, nextState.Bombs[:bi]...),
			nextState.Bombs[bi+1:]...)
		for _, d := range Directions {
			l := b.Location
			explosions[l] = true
			for i := 0; i < b.ExplosionRadius; i++ {
				l.X += d.X
				l.Y += d.Y
				if nextState.IsInside(&l) {
					if nextState.IsEmpty(&l) {
						explosions[l] = true
						nextState.Bombs.findChainedExplosions(l)
					} else {
						if !boardCloned {
							nextState.Board = s.Board.Clone()
							boardCloned = true
						}
						nextState.Board.OnExplosion(&l)
						break
					}
				}
			}
		}
		bi, b = nextState.Bombs.findExploding()
	}
	nextState.Bombs.decreaseCounters()
	return &nextState, toLocationSlice(explosions)
}

func toLocationSlice(set map[Location]bool) Locations {
	ls := []Location{}
	for l := range set {
		ls = append(ls, l)
	}
	return Locations(ls)
}

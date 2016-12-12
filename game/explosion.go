package game

func (s *State) Next() (*State, Locations) {

	nextState := *s
	nextState.Bombs = append([]Bomb{}, nextState.Bombs...)
	explosions := map[Location]bool{}
	for bi, b := nextState.Bombs.findExploding(); bi >= 0; bi, b = nextState.Bombs.findExploding() {
		nextState.Bombs = append(
			append([]Bomb{}, nextState.Bombs[:bi]...),
			nextState.Bombs[bi+1:]...)
		for _, d := range directions {
			l := b.Location
			explosions[l] = true
			for i := 0; i < b.ExplosionRadius; i++ {
				l.X += d.X
				l.Y += d.Y
				if nextState.IsEmpty(&l) {
					explosions[l] = true
					nextState.Bombs.findChainedExplosions(l)
				} else {
					break
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

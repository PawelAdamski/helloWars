package game

func (s *State) Next() (*State, []Location) {

	nextState := *s
	explosions := map[Location]bool{}
	nextState.Bombs.decreaseCounters()

	for bi, b := nextState.Bombs.findExploding(); bi >= 0; bi, b = nextState.Bombs.findExploding() {
		nextState.Bombs = append(nextState.Bombs[:bi], nextState.Bombs[bi+1:]...)
		for _, d := range directions {
			l := b.Location
			for i := 0; i < b.ExplosionRadius; i++ {
				l.X += d.X
				l.Y += d.Y
				if nextState.IsEmpty(l) {
					explosions[l] = true
					nextState.Bombs.findChainedExplosions(l)
				}
			}
		}

		bi, b = nextState.Bombs.findExploding()
	}
	return &nextState, explosions
}

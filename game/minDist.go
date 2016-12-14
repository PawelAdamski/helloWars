package game

func (s *State) minDistanceMap(start Location) [][]int {
	dist := slice2d(s.GameConfig.MapWidth, s.GameConfig.MapHeight, 100000)
	queue := []Location{start}
	dist[start.X][start.Y] = 0
	for len(queue) > 0 {
		l := queue[0]
		queue = queue[1:]
		for _, d := range Directions {
			nl := l.Translate(d)
			if s.IsEmpty(&nl) && dist[nl.X][nl.Y] > dist[l.X][l.Y]+1 {
				queue = append(queue, nl)
				dist[nl.X][nl.Y] = dist[l.X][l.Y] + 1
			}
		}
	}
	return dist
}

func slice2d(width, height, def int) [][]int {
	s := make([][]int, width)
	for i := range s {
		s[i] = make([]int, height)
		for j := range s[i] {
			s[i][j] = def
		}
	}
	return s
}

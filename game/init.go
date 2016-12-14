package game

func init() {
	for i, d := range Directions {
		inverseDirections[d] = i
	}
}

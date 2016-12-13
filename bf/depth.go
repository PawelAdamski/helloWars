package bf

type depth struct {
	me, opponent int
}

func (d depth) next() depth {
	return depth{
		me:       decrease(d.me),
		opponent: decrease(d.opponent),
	}
}

func decrease(i int) int {
	if i == 0 {
		return 0
	}
	return i - 1
}

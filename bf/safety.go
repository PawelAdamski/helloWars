package bf

type safety struct {
	me, other int
}

func (s *safety) or(o safety) {
	if s.me < o.me {
		s.me = o.me
	}
	if s.other > o.other {
		s.other = o.other
	}
}

func (s *safety) nor(o safety) {
	if s.me > o.me {
		s.me = o.me
	}
	if s.other < o.other {
		s.other = o.other
	}
}

func (s *safety) and(o safety) {
	if s.me > o.me {
		s.me = o.me
	}
	if s.other > o.other {
		s.other = o.other
	}
}

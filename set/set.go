package set

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{make(map[T]struct{})}
}

func (s Set[T]) Add(val T) {
	s.data[val] = struct{}{}
}

func (s Set[T]) Contains(val T) bool {
	_, ok := s.data[val]
	return ok
}

func (s Set[T]) Values() []T {
	var vv []T
	for v := range s.data {
		vv = append(vv, v)
	}
	return vv
}

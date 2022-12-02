package utils

type Set[T comparable] map[T]bool

func (s Set[T]) Add(i T) {
	s[i] = true
}

func (s Set[T]) Contains(i T) bool {
	return s[i]
}

func NewSet[T comparable](val ...T) (s Set[T]) {
	s = make(Set[T])
	for _, v := range val {
		s.Add(v)
	}
	return
}

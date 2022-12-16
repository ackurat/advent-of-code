package utils

type Set[T comparable] map[T]interface{}

func (s Set[T]) Add(i T) {
	s[i] = i
}

func (s Set[T]) Remove(i T) {
	s[i] = nil
}

func (s Set[T]) AddListOfItems(i []T) {
	for _, k := range i {
		s.Add(k)
	}
}

func (s Set[T]) Len() (len int) {
	for range s {
		len += 1
	}
	return
}

func (s Set[T]) Contains(i T) bool {
	return s[i] != nil
}

func NewSet[T comparable](val ...T) (s Set[T]) {
	s = make(Set[T])
	for _, v := range val {
		s.Add(v)
	}
	return
}

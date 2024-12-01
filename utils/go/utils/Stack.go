package utils

type Stack[T any] struct {
	stack []T
}

func NewStack[T any](item ...T) *Stack[T] {
	s := &(Stack[T]{stack: []T{}})
	for _, it := range item {
		s.Push(it)
	}
	return s
}

func (s *Stack[T]) Push(item T) {
	(*s).stack = append((*s).stack, item)
}

func (s *Stack[T]) Pop() (item T) {
	if len((*s).stack) == 0 {
		return
	}
	(*s).stack, item = (*s).stack[:len((*s).stack)-1], (*s).stack[len((*s).stack)-1]
	return
}

func (s *Stack[T]) Peek() (item T) {
	if len((*s).stack) == 0 {
		return
	}
	item = (*s).stack[len((*s).stack)-1]
	return
}

func (s *Stack[T]) IsEmpty() bool {
	return len((*s).stack) == 0
}

func (s *Stack[T]) Length() int {
	return len((*s).stack)
}

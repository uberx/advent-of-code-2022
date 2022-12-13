package util

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (element T, popped bool) {
	if !s.IsEmpty() {
		index := len(s.elements) - 1
		element = s.elements[index]
		s.elements = s.elements[:index]
		popped = true
	}
	return element, popped
}

package util

type Queue[T any] struct {
	elements []T
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Queue[T]) Queue(element T) {
	s.elements = append(s.elements, element)
}

func (s *Queue[T]) Dequeue() (element T, dequeued bool) {
	if !s.IsEmpty() {
		element = s.elements[0]
		s.elements = s.elements[1:]
		dequeued = true
	}
	return element, dequeued
}

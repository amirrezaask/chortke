package main

type stack[T any] struct {
	data []T
}

func (s *stack[T]) push(elem T) {
	s.data = append(s.data, elem)
}
func (s *stack[T]) top() *T {
	if len(s.data) == 0 { return nil }
	return &s.data[len(s.data)-1]
}

func (s *stack[T]) pop() *T {
	if len(s.data) == 0 { return new(T) }
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return &top
}

type queue[T any] struct {
	data []T
}

func (s *queue[T]) push(elem T) {
	s.data = append(s.data, elem)
}

func (s *queue[T]) pop() *T {
	if len(s.data) == 0 { return new(T) }
	top := s.data[0]
	s.data = s.data[1:]
	return &top
}

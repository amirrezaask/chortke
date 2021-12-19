package main
import (
    "fmt"
)

type stack[T any] struct {
    data []T
}

func (s *stack[T]) push(elem T) {
	s.data = append(s.data, elem)
}

func (s *stack[T]) pop() T {
	fmt.Println("stack len ", len(s.data))

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

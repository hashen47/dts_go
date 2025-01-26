package stack

import "fmt"

type Stack[T comparable] struct {
	I   int
	Buf []T
}

func StackInit[T comparable]() *Stack[T] {
	buf := make([]T, 0)

	return &Stack[T]{
		I:   -1,
		Buf: buf,
	}
}

func (s *Stack[T]) Push(el T) {
	s.I++
	s.Buf = append(s.Buf, el)
}

func (s *Stack[T]) IsEmpty() bool {
	if s.I < 0 {
		return true
	}
	return false
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var result T
		return result, fmt.Errorf("stack is empty")
	}

	el := s.I
	s.I--

	return s.Buf[el], nil
}

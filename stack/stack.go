package stack

import "fmt"

type Stack[T comparable] struct {
	i   int
	buf *[]T
}

var StackEmptyErr error = fmt.Errorf("Stack is Empty")

func StackInit[T comparable]() *Stack[T] {
	buf := make([]T, 0)

	return &Stack[T]{
		i:   -1,
		buf: &buf,
	}
}

func (s Stack[T]) String() string {
    return fmt.Sprintf("Stack<i: %d, buf: %v>\n", s.i, *s.buf)
}

func (s *Stack[T]) Push(el T) {
	s.i++
	*s.buf = append(*s.buf, el)
}

func (s *Stack[T]) IsEmpty() bool {
	if s.i < 0 {
		return true
	}
	return false
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var result T
		return result, StackEmptyErr 
	}

	el := s.i
	s.i--

	return (*s.buf)[el], nil
}

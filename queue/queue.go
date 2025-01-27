package queue

import (
	"fmt"
)

type Queue[T comparable] struct {
	i   int
	buf *[]T
}

var QEmptyErr error = fmt.Errorf("Queue is empty")

func Queueinit[T comparable]() *Queue[T] {
	buf := make([]T, 0)

	return &Queue[T]{
		i:   -1,
		buf: &buf,
	}
}

func (q Queue[T]) String() string {
	return fmt.Sprintf("Queue<i: %d, buf: %v>\n", q.i, *q.buf)
}

func (q *Queue[T]) Enqueue(el T) {
	if q.i == -1 {
		q.i = 0
	}
	*q.buf = append(*q.buf, el)
}

func (q *Queue[T]) Peek() (el T, err error) {
	var result T

	if q.i < 0 || q.i >= len(*q.buf) {
		return result, QEmptyErr
	}

	return (*q.buf)[q.i], nil
}

func (q *Queue[T]) Dequeue() (el T, err error) {
	var result T

	if q.i < 0 || q.i >= len(*q.buf) {
		return result, QEmptyErr
	}

	index := q.i
	q.i++

	return (*q.buf)[index], nil
}

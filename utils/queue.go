package uitls

import "fmt"

type Stack[T any] []T

func (s Stack[T]) Empty() bool {
	return len(s) == 0
}

func (s Stack[T]) Len() int {
	return len(s)
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s Stack[T]) Top() T {
	n := len(s)
	return s[n-1]
}

func (s *Stack[T]) Pop() T {
	old := *s
	n := len(old)
	v := old[n-1]
	*s = old[:n-1]
	return v
}

func (s *Stack[T]) Clear() {
	old := *s
	*s = old[:0]
}

type Queue[T any] struct {
	input  Stack[T]
	output Stack[T]
}

func NewQueueSize[T any](size int) *Queue[T] {
	return &Queue[T]{
		input:  make(Stack[T], 0, size),
		output: make(Stack[T], 0, size),
	}
}

func (q Queue[T]) Empty() bool {
	return q.input.Empty() && q.output.Empty()
}

func (q Queue[T]) Len() int {
	return q.input.Len() + q.output.Len()
}

func (q *Queue[T]) Push(v T) {
	q.input.Push(v)
}

func (q Queue[T]) Front() T {
	if q.output.Empty() {
		return q.input[0]
	}
	return q.output.Top()
}

func (q Queue[T]) Back() T {
	if q.input.Empty() {
		return q.output[0]
	}
	return q.input.Top()
}

func (q *Queue[T]) Pop() T {
	q.Pour()
	return q.output.Pop()
}

func (q *Queue[T]) Pour() {
	if len(q.output) == 0 {
		for !q.input.Empty() {
			q.output.Push(q.input.Pop())
		}
	}
}

func (q Queue[T]) Slice() []T {
	a := make([]T, 0, q.Len())
	for i := len(q.output) - 1; i >= 0; i-- {
		a = append(a, q.output[i])
	}
	for _, v := range q.input {
		a = append(a, v)
	}
	return a
}

func (q Queue[T]) String() string {
	return fmt.Sprint(q.Slice())
}

func (q *Queue[T]) Clear() {
	q.input.Clear()
	q.output.Clear()
}

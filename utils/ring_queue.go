package uitls

// RingQueue (!) хз на сколько рабочее...
type RingQueue[T any] struct {
	items []T
	len   int
	front int
}

func NewDequeSize[T any](size int) *RingQueue[T] {
	return &RingQueue[T]{
		items: make([]T, size),
	}
}

func (q *RingQueue[T]) Len() int {
	return q.len
}

func (q *RingQueue[T]) Empty() bool {
	return q.len == 0
}

func (q *RingQueue[T]) Clear() {
	q.front = 0
	q.len = 0
}

func (q *RingQueue[T]) grow(n int) {
	size := len(q.items)
	newSize := max(size*2, size+n, 8)

	buf := make([]T, newSize)
	if back := q.front + q.len; back <= size {
		copy(buf, q.items[q.front:back])
	} else {
		copy(buf, q.items[q.front:])
		copy(buf[size-q.front:], q.items[:back-size])
	}

	q.items = buf
	q.front = 0
}

func (q *RingQueue[T]) Push(v T) {
	if q.len == len(q.items) {
		q.grow(1)
	}
	back := q.front + q.len
	if size := len(q.items); back >= size {
		back -= size
	}
	q.items[back] = v
	q.len++
}

func (q *RingQueue[T]) Pop() T {
	if q.len == 0 {
		panic("queue is empty")
	}
	v := q.items[q.front]
	q.front++
	if q.front == len(q.items) {
		q.front = 0
	}
	q.len--
	return v
}

package containers

// Queue is a simple FIFO (first-in, first-out) queue data structure.
// It is implemented using a slice and supports generic types.
type Queue[T any] struct {
	data []T
}

// NewQueue returns an initialized empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0)}
}

// Enqueue adds a value to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

// Dequeue removes and returns the front value from the queue.
// If the queue is empty, it returns the zero value of T and false.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

// Peek returns the front value of the queue without removing it.
// If the queue is empty, it returns the zero value of T and false.
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}
	return q.data[0], true
}

// Len returns the number of elements currently in the queue.
func (q *Queue[T]) Len() int {
	return len(q.data)
}

// Package containers provides basic data structures that are not part of
// Go's standard library. These structures are implemented with idiomatic Go
// patterns and are generic over element type.
package containers

// Stack is a simple LIFO (last-in, first-out) stack data structure.
// It is implemented using a slice and supports generic types.
type Stack[T any] struct {
	data []T
}

// NewStack returns an initialized empty stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

// Push adds a value onto the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top value from the stack.
// If the stack is empty, it returns the zero value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

// Peek returns the top value of the stack without removing it.
// If the stack is empty, it returns the zero value of T and false.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Len returns the number of elements currently in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

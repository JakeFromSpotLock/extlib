// Package containers provides common data structures like Stack, Queue, and Set.
package containers

// Set is a generic implementation of a mathematical set.
// It uses a Go map with empty structs for memory efficiency.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet creates and returns a new empty set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// Add inserts an element into the set.
func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

// Remove deletes an element from the set.
func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

// Contains checks whether the set contains the given value.
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Values returns a slice of all elements in the set.
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for v := range s.data {
		values = append(values, v)
	}
	return values
}

// Union returns a new set containing all elements from both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for v := range s.data {
		result.Add(v)
	}
	for v := range other.data {
		result.Add(v)
	}
	return result
}

// Intersection returns a new set containing only elements present in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for v := range s.data {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns a new set containing elements in s but not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for v := range s.data {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

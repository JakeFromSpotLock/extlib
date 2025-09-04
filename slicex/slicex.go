// Package slicex provides generic slice utilities for Go. These functions
// complement Go's standard library by offering common operations such as
// reversing, mapping, filtering, and checking membership.
package slicex

// Reverse reverses the order of elements in the given slice in-place.
//
// Example:
//   s := []int{1, 2, 3}
//   slicex.Reverse(s) // s = []int{3, 2, 1}
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Contains returns true if the slice contains the given element.
//
// Example:
//   found := slicex.Contains([]int{1, 2, 3}, 2) // true
func Contains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

// Map applies a function f to each element of the slice and returns a new
// slice containing the results.
//
// Example:
//   s := []int{1, 2, 3}
//   doubled := slicex.Map(s, func(x int) int { return x*2 }) // [2, 4, 6]
func Map[T any, R any](s []T, f func(T) R) []R {
	result := make([]R, len(s))
	for i, x := range s {
		result[i] = f(x)
	}
	return result
}

// Filter returns a new slice containing only elements for which the
// predicate function returns true.
//
// Example:
//   s := []int{1, 2, 3, 4}
//   even := slicex.Filter(s, func(x int) bool { return x%2 == 0 }) // [2, 4]
func Filter[T any](s []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, x := range s {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

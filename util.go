// Package extlib provides general-purpose utilities that extend Go's standard
// library. The top-level package exposes helpers that are broadly applicable
// and do not belong in a more specific subpackage.
package extlib

// Ternop acts like a ternary operator. If cond is true, it returns a;
// otherwise, it returns b.
//
// Example:
//
//   max := extlib.Ternop(x > y, x, y)
//
// This is generic over any comparable type.
func Ternop[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

// Ptr returns a pointer to the given value.
//
// Example:
//
//   s := extlib.Ptr("hello")
//   fmt.Println(*s) // "hello"
//
// This is useful for initializing optional values without writing boilerplate.
func Ptr[T any](v T) *T {
	return &v
}

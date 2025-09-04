// Package mathx provides extended mathematical utilities that complement
// Go's standard library. It includes generic helpers, slice operations,
// combinatorics, and bitwise operations for numbers.
//
// Compatible with Go 1.20+ using built-in constraints for ordered types.
package mathx

import "cmp"

// Max returns the larger of a and b.
//
// Example:
//   m := Max(3, 7) // 7
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of a and b.
//
// Example:
//   m := Min(3, 7) // 3
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Clamp restricts x to the range [min, max].
//
// Example:
//   c := Clamp(10, 0, 5) // 5
func Clamp[T cmp.Ordered](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Abs returns the absolute value of x.
//
// Example:
//   a := Abs(-5) // 5
func Abs[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sum returns the sum of a slice of numbers.
//
// Example:
//   s := Sum([]int{1,2,3}) // 6
func Sum[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// Product returns the product of a slice of numbers.
//
// Example:
//   p := Product([]int{2,3,4}) // 24
func Product[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](nums []T) T {
	prod := T(1)
	for _, n := range nums {
		prod *= n
	}
	return prod
}

// IsPowerOf returns true if n is an exact power of the given base.
// base must be >= 2, otherwise the function returns false.
//
// Examples:
//   IsPowerOf(8, 2)  // true, because 2^3 = 8
//   IsPowerOf(27, 3) // true, because 3^3 = 27
//   IsPowerOf(16, 4) // true, because 4^2 = 16
//   IsPowerOf(10, 2) // false
func IsPowerOf(n int, base int) bool {
	if n < 1 || base < 2 {
		return false
	}

	for n > 1 {
		if n%base != 0 {
			return false
		}
		n /= base
	}
	return true
}

// Factorial returns n! for small integers (overflow can happen quickly).
//
// Example:
//   f := Factorial(5) // 120
func Factorial(n int) int64 {
	if n < 0 {
		return 0
	}
	var result int64 = 1
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

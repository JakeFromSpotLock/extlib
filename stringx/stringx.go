// Package stringx provides extended string utilities.
// All functions are rune-safe and aim to handle UTF-8 correctly.
package stringx

import (
	"strings"
	"unicode/utf8"
)

// Reverse returns a new string with the runes of s reversed.
// Handles multi-byte UTF-8 characters safely.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// PadLeft returns s left-padded with pad until it reaches length n.
// If s is already >= n runes, it is returned unchanged.
func PadLeft(s, pad string, n int) string {
	length := utf8.RuneCountInString(s)
	if length >= n {
		return s
	}
	needed := n - length
	return strings.Repeat(pad, needed) + s
}

// PadRight returns s right-padded with pad until it reaches length n.
// If s is already >= n runes, it is returned unchanged.
func PadRight(s, pad string, n int) string {
	length := utf8.RuneCountInString(s)
	if length >= n {
		return s
	}
	needed := n - length
	return s + strings.Repeat(pad, needed)
}

// EqualsFold checks if two strings are equal under Unicode
// case-folding (i.e., case-insensitive comparison).
func EqualsFold(a, b string) bool {
	return strings.EqualFold(a, b)
}

// Package extlib (Go Extended Library) provides a set of utilities and data
// structures designed to complement Go's standard library. Its purpose is to
// offer practical, frequently-needed helpers while maintaining idiomatic Go
// design.
//
// # Overview
//
// The Go Extended Library (extlib) is organized as a core package (`extlib`)
// with additional functionality provided in subpackages. Together, they form
// a collection of tools that extend Go's standard capabilities without
// replacing them.
//
// Top-level utilities in `extlib` are small, general-purpose helpers that
// don't warrant a dedicated subpackage. For example:
//
//   - Ternop: A generic ternary operator function.
//   - Ptr: A helper for obtaining pointers to values.
//
// Subpackages are grouped by domain. For example:
//
//	extlib/containers   → stack and related data structures.
//	extlib/mathx   		→ extended mathematical utilities.
//	extlib/slicex       → extended slice utilities.
//
// # File Structure
//
// A typical extlib project layout:
//
//	extlib/
//	  doc.go         → package-level documentation
//	  util.go        → general-purpose helpers (Ternop, Ptr, etc.)
//	  version.go     → library version metadata
//	  containers/    → stack/queue/etc and related data structures
//	  mathx/         → extended math utilities
//	  slicex/        → extended slice utilities
//
// # Example Usage
//
// Importing the base package:
//
//	import "github.com/jakefromspotlock/extlib"
//
// Using a top-level utility:
//
//	max := extlib.Ternop(x > y, x, y)
//
// Importing and using a subpackage:
//
//	import "github.com/jakefromspotlock/extlib/containers"
//
//	s := containers.NewStack[int]()
//	s.Push(10)
//	s.Push(20)
//	v := s.Pop() // 20
//
// # Contributing
//
// This library is intended to grow gradually with practical, reusable
// utilities. Contributions are welcome, provided they follow idiomatic Go
// patterns and avoid unnecessary abstraction.
//
// # Author
//
// Jake Wagner
// jake@spotlockapp.com
//
// Feel free to reach out with suggestions, bug reports, or contributions.
package extlib

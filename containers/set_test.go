package containers

import "testing"

func TestSetBasicOperations(t *testing.T) {
	s := NewSet[int]()

	// Test Add and Contains
	s.Add(1)
	s.Add(2)
	if !s.Contains(1) || !s.Contains(2) {
		t.Errorf("Expected set to contain 1 and 2")
	}
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}

	// Test Remove
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Expected 1 to be removed from set")
	}
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
}

func TestSetUnionIntersectionDifference(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewSet[int]()
	b.Add(3)
	b.Add(4)

	// Union
	union := a.Union(b)
	if union.Size() != 4 {
		t.Errorf("Expected union size 4, got %d", union.Size())
	}

	// Intersection
	intersection := a.Intersection(b)
	if intersection.Size() != 1 || !intersection.Contains(3) {
		t.Errorf("Expected intersection to contain only 3")
	}

	// Difference
	diff := a.Difference(b)
	if diff.Size() != 2 || !diff.Contains(1) || !diff.Contains(2) {
		t.Errorf("Expected difference to contain 1 and 2")
	}
}

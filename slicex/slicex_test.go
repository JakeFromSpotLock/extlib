package slicex

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4}
	Reverse(s)
	want := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("Reverse failed: got %v, want %v", s, want)
	}
}

func TestContains(t *testing.T) {
	s := []string{"a", "b", "c"}
	if !Contains(s, "b") {
		t.Error("Contains failed: 'b' should be found")
	}
	if Contains(s, "x") {
		t.Error("Contains failed: 'x' should not be found")
	}
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3}
	doubled := Map(s, func(x int) int { return x * 2 })
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(doubled, want) {
		t.Errorf("Map failed: got %v, want %v", doubled, want)
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	even := Filter(s, func(x int) bool { return x%2 == 0 })
	want := []int{2, 4}
	if !reflect.DeepEqual(even, want) {
		t.Errorf("Filter failed: got %v, want %v", even, want)
	}
}

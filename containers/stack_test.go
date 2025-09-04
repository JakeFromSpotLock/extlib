package containers

import "testing"

func TestStack(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)

		val, ok := s.Pop()
		if !ok || val != 2 {
			t.Errorf("expected 2, got %v (ok=%v)", val, ok)
		}

		val, ok = s.Pop()
		if !ok || val != 1 {
			t.Errorf("expected 1, got %v (ok=%v)", val, ok)
		}

		_, ok = s.Pop()
		if ok {
			t.Errorf("expected empty pop to return ok=false")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		s := NewStack[string]()
		s.Push("hello")
		s.Push("world")

		val, ok := s.Peek()
		if !ok || val != "world" {
			t.Errorf("expected 'world', got %v (ok=%v)", val, ok)
		}

		if s.Len() != 2 {
			t.Errorf("expected length 2, got %d", s.Len())
		}
	})
}

package containers

import "testing"

func TestQueue(t *testing.T) {
	t.Run("Enqueue and Dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		q.Enqueue(2)

		val, ok := q.Dequeue()
		if !ok || val != 1 {
			t.Errorf("expected 1, got %v (ok=%v)", val, ok)
		}

		val, ok = q.Dequeue()
		if !ok || val != 2 {
			t.Errorf("expected 2, got %v (ok=%v)", val, ok)
		}

		_, ok = q.Dequeue()
		if ok {
			t.Errorf("expected empty dequeue to return ok=false")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		q := NewQueue[string]()
		q.Enqueue("foo")
		q.Enqueue("bar")

		val, ok := q.Peek()
		if !ok || val != "foo" {
			t.Errorf("expected 'foo', got %v (ok=%v)", val, ok)
		}

		if q.Len() != 2 {
			t.Errorf("expected length 2, got %d", q.Len())
		}
	})
}

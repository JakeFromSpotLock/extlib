package extlib

import "testing"

func TestTernop(t *testing.T) {
	tests := []struct {
		cond bool
		a, b int
		want int
	}{
		{true, 42, 7, 42},
		{false, 42, 7, 7},
	}

	for _, tt := range tests {
		got := Ternop(tt.cond, tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Ternop(%v, %v, %v) = %v; want %v",
				tt.cond, tt.a, tt.b, got, tt.want)
		}
	}
}

func TestPtr(t *testing.T) {
	val := 99
	p := Ptr(val)

	if p == nil {
		t.Fatalf("Ptr(%v) returned nil", val)
	}

	if *p != val {
		t.Errorf("Ptr(%v) = %v; want %v", val, *p, val)
	}
}

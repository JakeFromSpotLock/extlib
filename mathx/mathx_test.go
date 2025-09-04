package mathx

import "testing"

func TestMinMaxClamp(t *testing.T) {
	if Max(3, 7) != 7 {
		t.Error("Max failed")
	}
	if Min(3, 7) != 3 {
		t.Error("Min failed")
	}
	if Clamp(10, 0, 5) != 5 {
		t.Error("Clamp failed")
	}
	if Clamp(-2, 0, 5) != 0 {
		t.Error("Clamp failed")
	}
	if Clamp(3, 0, 5) != 3 {
		t.Error("Clamp failed")
	}
}

func TestAbs(t *testing.T) {
	if Abs(-5) != 5 {
		t.Error("Abs failed")
	}
	if Abs(5) != 5 {
		t.Error("Abs failed")
	}
}

func TestSumProduct(t *testing.T) {
	nums := []int{2, 3, 4}
	if Sum(nums) != 9 {
		t.Error("Sum failed")
	}
	if Product(nums) != 24 {
		t.Error("Product failed")
	}
}

func TestIsPowerOf(t *testing.T) {
	tests := []struct {
		n, base int
		want    bool
	}{
		{8, 2, true},   // 2^3 = 8
		{16, 2, true},  // 2^4 = 16
		{27, 3, true},  // 3^3 = 27
		{9, 3, true},   // 3^2 = 9
		{16, 4, true},  // 4^2 = 16
		{81, 3, true},  // 3^4 = 81
		{10, 2, false}, // not a power of 2
		{15, 3, false}, // not a power of 3
		{0, 2, false},  // invalid n
		{1, 1, false},  // invalid base
		{32, 5, false}, // not a power of 5
	}

	for _, tt := range tests {
		got := IsPowerOf(tt.n, tt.base)
		if got != tt.want {
			t.Errorf("IsPowerOf(%d, %d) = %v; want %v", tt.n, tt.base, got, tt.want)
		}
	}
}

func TestFactorial(t *testing.T) {
	if Factorial(5) != 120 {
		t.Error("Factorial failed")
	}
	if Factorial(0) != 1 {
		t.Error("Factorial failed")
	}
	if Factorial(-3) != 0 {
		t.Error("Factorial failed")
	}
}

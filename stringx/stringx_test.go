package stringx

import "testing"

func TestReverse(t *testing.T) {
	tests := map[string]string{
		"abc":    "cba",
		"世界":      "界世",
		"a😊b":    "b😊a", // emoji check
	}
	for input, expected := range tests {
		if got := Reverse(input); got != expected {
			t.Errorf("Reverse(%q) = %q, want %q", input, got, expected)
		}
	}
}

func TestPadLeft(t *testing.T) {
	got := PadLeft("42", "0", 5)
	want := "00042"
	if got != want {
		t.Errorf("PadLeft() = %q, want %q", got, want)
	}
}

func TestPadRight(t *testing.T) {
	got := PadRight("42", " ", 5)
	want := "42   "
	if got != want {
		t.Errorf("PadRight() = %q, want %q", got, want)
	}
}

func TestEqualsFold(t *testing.T) {
	if !EqualsFold("GoLang", "golang") {
		t.Errorf("Expected EqualsFold to consider 'GoLang' == 'golang'")
	}
	if EqualsFold("Go", "Java") {
		t.Errorf("Expected EqualsFold to consider 'Go' != 'Java'")
	}
}

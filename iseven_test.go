package iseven_test

import (
	"testing"
)

func TestSimple(t *testing.T) {
	tests := []struct {
		input interface{}
		want  bool
	}{
		{1, false},
		{2, true},
		{"2", true},
		{"Three", false},
		{"Four", true},
	}
	for _, test := range tests {
		got := IsEven(test.input)
		if got != test.want {
			t.Fatalf("want %v for %v, but %v:", tests.want, test.input, got)
		}
	}
}

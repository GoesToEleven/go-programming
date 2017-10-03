package mymath

import "testing"

// Fibonacci sequence adds up the previous two numbers:
// 0 1 1 2 3 5 8 13 21
func TestFib(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}

	for _, test := range tests {
		if got := Fib(test.input); got != test.want {
			t.Errorf(`got Fib(%d) = %d wanted %d`, test.input, got, test.want)
		}
	}
}

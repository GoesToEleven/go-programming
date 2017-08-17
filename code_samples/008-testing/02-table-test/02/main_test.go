package main

import "testing"

func TestMyAdd(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 3, 4}, 9},
		test{[]int{42, 43, 44}, 129},
		test{[]int{5, 7}, 12},
		test{[]int{7, 7, 7}, 21},
		test{[]int{14, 18}, 32},
		test{[]int{12, 19}, 31},
		test{[]int{1, 1}, 2},
	}

	for _, v := range tests {
		x := myAdd(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

package main

import "testing"

func TestMyAdd(t *testing.T) {
	if myAdd(2, 3) != 2+3 {
		t.Error("Expected", 5, "Got", myAdd(2, 3))
	}
}

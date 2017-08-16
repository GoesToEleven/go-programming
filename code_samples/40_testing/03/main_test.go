package main

import (
	"github.com/GoesToEleven/go-programming/code_samples/40_testing/03/example"
	"testing"
)

func TestSum(t *testing.T) {
	var n int
	n = example.Sum(1, 2)
	if n != 3 {
		t.Error("Expected 3, got ", n)
	}
}

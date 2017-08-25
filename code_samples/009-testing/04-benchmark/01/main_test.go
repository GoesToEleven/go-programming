package main

import (
	"testing"
	"fmt"
)

func BenchmarkHell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("Hello")
	}
}
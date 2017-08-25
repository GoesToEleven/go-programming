package word

import (
	"testing"
	"fmt"
)

func TestCount(t *testing.T) {
	s := "Shaken not stirred"
	n := Count(s)
	if n != 3 {
		t.Error("got", s, "want", 3)
	}
}

func TestUseCount(t *testing.T) {
	s := "stirred Shaken not stirred"
	m := UseCount(s)
	for k, v := range m {
		switch k {
		case "stirred":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		case "Shaken":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "not":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		}
	}
}

func ExampleCount() {
	s := "Shaken not stirred"
	fmt.Println(Count(s))
	// Output:
	// 3
}

func ExampleUseCount() {
	s := "stirred Shaken not stirred"
	m := UseCount(s)
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(UseCount(s))
	// Output:
	// Shaken 1
	// not 1
	// stirred 2
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(Q)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(Q)
	}
}

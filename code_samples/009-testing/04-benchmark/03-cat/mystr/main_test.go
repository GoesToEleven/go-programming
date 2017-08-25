package mystr

import (
	"testing"
	"strings"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
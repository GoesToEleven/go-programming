package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("mom") {
		t.Error(`IsPalindrome("mom") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("dog") {
		t.Error(`IsPalindrone("dog") = true`)
	}
}

func TestFrenchPalindrone(t *testing.T) {
	s := "été"
	if !IsPalindrome(s) {
		t.Errorf(`IsPalindrone(%q) = false`, s)
	}
}

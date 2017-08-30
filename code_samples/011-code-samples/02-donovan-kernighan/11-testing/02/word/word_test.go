package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"mom", true},
		{"kayak", true},
		{"dog", false},
		{"été", true},
		{"amanaplanacanalpanama", true},
		{"", true},
		{"", true},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf(`IsPalindrone(%q) = %v`, test.input, got)
		}
	}
}

// Package word provides utilities for word games.
package word

import "unicode"

// IsPalidrome reports whether s reads the same forward and backward.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[(len(letters)-1)-i] {
			return false
		}
	}
	return true
}

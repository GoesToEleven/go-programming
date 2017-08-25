package mystr

import "strings"

func Cat(xs []string) string {
	s := ""
	for _, v := range xs {
		s += v
		s += " "
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}


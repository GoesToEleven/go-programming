package main

import "fmt"

func main() {
	fmt.Println(Adder(2, 3))
}

func Adder(xs ...int) int {
	res := 0
	for _, v := range xs {
		res += v
	}
	return res
}

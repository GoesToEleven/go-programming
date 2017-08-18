package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", myAdd(2, 3))
	fmt.Println("5 + 7 =", myAdd(5, 7))
}

func myAdd(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

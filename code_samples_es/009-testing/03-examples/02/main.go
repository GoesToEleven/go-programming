package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/007-documentation/01/mymath"
)

func main() {
	fmt.Println("2 + 3 =", mymath.Sum(2, 3))
	fmt.Println("4 + 7 =", mymath.Sum(4, 7))
	fmt.Println("5 + 9 =", mymath.Sum(5, 9))
}

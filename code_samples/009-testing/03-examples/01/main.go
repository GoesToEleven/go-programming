package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/009-testing/03-examples/01/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}

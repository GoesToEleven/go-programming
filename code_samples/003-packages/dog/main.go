package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/003-packages/cat"
)

func main() {
	fmt.Println("Hello from dog.")
	cat.Hello()
	cat.Eats()
}

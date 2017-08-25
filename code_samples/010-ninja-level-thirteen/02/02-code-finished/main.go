package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/02-code-finished/word"
)


func main() {
	fmt.Println(word.Count(word.Q))

	for k, v := range word.UseCount(word.Q) {
		fmt.Println(v, k)
	}
}
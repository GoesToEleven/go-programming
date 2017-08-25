package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/word"
)


func main() {
	fmt.Println(word.Count(q))

	for k, v := range word.UseCount(q) {
		fmt.Println(v, k)
	}
}
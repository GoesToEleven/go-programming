package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/011-code-samples/02-donovan-kernighan/02-structure/02-celsius-farenheit/tempconv"
)

func main() {
	fmt.Printf("Fahrenheit freezing %g\n", tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("Fahrenheit boiling %g\n", tempconv.CToF(tempconv.BoilingC))
}

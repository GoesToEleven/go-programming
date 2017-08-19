package main

import (
	"github.com/GoesToEleven/go-programming/code_samples/008-ninja-level-twelve/01x/dog"
	"fmt"
)

type canine struct {
	name string
	age int
}

func main() {
	fido := canine{
		name: "Fido",
		age: dog.Years(10),
	}
	fmt.Println(fido)
}

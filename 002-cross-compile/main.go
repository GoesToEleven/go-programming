package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime: %s\narchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}

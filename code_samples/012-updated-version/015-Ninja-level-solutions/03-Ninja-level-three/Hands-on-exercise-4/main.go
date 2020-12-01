package main

import "fmt"

func main() {
	bd := 1985
	for {
		if bd > 2017 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

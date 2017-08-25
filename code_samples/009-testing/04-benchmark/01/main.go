package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	n := 4
	wg.Add(n)
	xi := gen(n)
	c := fact(xi)

	for v := range c {
		fmt.Println(v)
	}
}

func gen(n int) []int {
	xi := make([]int, n)
	for i := 1; i <= n; i++ {
		xi = append(xi, i)
	}
	return xi
}


func fact(xi []int) chan map[int]int {
	c := make(chan map[int]int)

	// calculate factorial
	for _, v := range xi {
		go func(w int){
			c <- map[int]int{w:factorial(w)}
			wg.Done()
		}(v)
	}

	// close channel when done
	go func(){
		wg.Wait()
		close(c)
	}()
	return c
}

func factorial(w int) int {
	y := 1
	for i := w; i > 0; i-- {
		y = y * i
	}
	return y
}
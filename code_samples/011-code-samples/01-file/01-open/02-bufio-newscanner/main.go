package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("dalia-shevin.txt")
	if err != nil {
		log.Fatalf("error opening  file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			fmt.Println(strings.ToUpper(line))
		}
	}
}

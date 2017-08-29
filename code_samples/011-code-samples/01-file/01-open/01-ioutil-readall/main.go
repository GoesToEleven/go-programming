package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func main() {
	f, err := os.Open("rumi.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	fmt.Println(string(bs))
}
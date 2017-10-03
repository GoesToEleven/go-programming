package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("dog-poem.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	nf, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer nf.Close()

	_, err = io.Copy(nf, f)
	if err != nil {
		log.Fatalf("error copying file: %v", err)
	}
}

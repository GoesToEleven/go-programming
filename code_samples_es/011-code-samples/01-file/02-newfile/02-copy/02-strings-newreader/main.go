package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	nf, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer nf.Close()

	s := `
	Don’t surrender your loneliness So quickly.
	Let it cut more deep. Let it ferment and season you
	as few human or even divine ingredients can.
	- Hafiz
	`

	rdr := strings.NewReader(s)

	_, err = io.Copy(nf, rdr)
	if err != nil {
		log.Fatalf("error copying to file: %v", err)
	}
}

package main

import (
	"strings"
	"io"
	"os"
	"log"
)

func main() {

	r := strings.NewReader(os.Args[1])

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
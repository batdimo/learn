package main

import (
	"fmt"
	"io"
	"strings"
	"os"
)

func main() {
	reader := strings.NewReader("Hello World!")

	written, err := io.CopyN(os.Stdout, reader, 8) // copy up to 8 bytes

	fmt.Printf("\n%d, %v", written, err)
}
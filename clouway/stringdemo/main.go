package main

import (
	"fmt"
	"github.com/batdimo/gostring"
	"strings"
)

func main() {
	leng := 0
	fmt.Print("Enter length of string: ")
	fmt.Scan(&leng)
	fmt.Println(strings.Join(gostring.Str(leng), ""))
}

package main

import (
	"fmt"
	"github.com/drafailov/dimitar_rafailov_go_intro/cmd/stringdemo/gostring"
)

func main() {
	leng := 0
	fmt.Print("Enter length of string: ")
	fmt.Scan(&leng)
    fmt.Print(gostring.Str(leng))
}

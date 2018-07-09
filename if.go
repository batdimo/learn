package main

import "fmt"

func main() {
	if 5 > 7 {
		fmt.Print(" if")
	} else {
		fmt.Print("else")
	}

	if fmt.Print("if you"); 5 > 2 {
		fmt.Print("belive")
	} else if 4 > 3 {
		fmt.Print("no way")
	} else {
		fmt.Print("trwe")
	}

	var c, a, b int
	if c = b; a > b {
		c = a
		fmt.Print(c)

	}
}

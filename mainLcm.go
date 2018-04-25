package main

import (
	"fmt"
	d "github.com/drafailov/learn/divisor"
)

func main() {
	var a, b int
	fmt.Print("Insert first number: ")
	fmt.Scanln(&a)
	fmt.Print("Insert second number: ")
	fmt.Scanln(&b)
	fmt.Printf("lcm(%d,%d)= ", a, b) //Least common multiple
	if a&b == 0 {
		fmt.Print(0)
	} else {
		fmt.Print(a * b / d.Div(a, b))
	}
}
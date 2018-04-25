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
	fmt.Printf("gcd(%d,%d)= ", a, b) // Greatest common divisor
	fmt.Println(d.Div(a, b))
}

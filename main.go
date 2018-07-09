package main

import (
	"fmt"
	d "github.com/batdimo/learn/divisor"
)

func main() {
	var a,b int
	fmt.Print("Insert first number: ")
	fmt.Scanln(&a)
	fmt.Print("Insert second number: ")
	fmt.Scanln( &b)
	c:= d.Div(a,b)
	fmt.Print("The greatest common divisor is: ")
	fmt.Println(c)
	fmt.Print("The least common multiple  is: ")
	fmt.Print(a*b/c)
}

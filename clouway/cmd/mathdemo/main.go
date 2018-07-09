package main

import (
	"fmt"
	"github.com/drafailov/dimitar_rafailov_go_intro/gomath"
)

func main() {
	var a, b int
	fmt.Print("Insert first number: ")
	fmt.Scanln(&a)
	fmt.Print("Insert second number: ")
	fmt.Scanln(&b)
	fmt.Printf("gcd(%d,%d)= ", a, b)
	fmt.Println(gomath.GCD(a, b))
	fmt.Printf("lcm(%d,%d)= ", a, b)
	fmt.Println(gomath.LCM(a, b))
}

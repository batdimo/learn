package main

import "fmt"

func divb(a, b int) int {
	var k int
	for b != 0 {
		k = a % b
		a = b
		b = k
	}
	return a
}
func main() {

	fmt.Print(divb(64, 344))
}

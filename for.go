package main

import (
	"fmt"
)

func Sqrt1(x float64) float64 {

	z:=1.0
	for i:=0;i<10;i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt1(169))
}
func Print(arr []int) {
	fmt.Print("[")
	for i, e := range arr {
		if i == len(arr)-1 {
			fmt.Printf("%v", e)
		} else {
			fmt.Printf("%v ", e)
		}
	}
	fmt.Println("]")
}
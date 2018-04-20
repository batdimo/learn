package main

import "fmt"

func MinElement(array []int) int {
	var m int
	for _, e := range array {
		if e < m {
			m = e
		}
	}
	return m

}
func main (){
	b :=[] int {3,4,44,2,11}
	fmt.Print(MinElement(b))
}
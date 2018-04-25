package main

import (
	"fmt"
	a "github.com/drafailov/learn/arrayOperations"
)

func main() {
	var array = [] int{3, 5, 6, 7, 0, 6, -3}
	//fmt.Println(a.MinElement(array))
	//fmt.Println(a.Sum(array))
	//fmt.Println(a.Print(array))
	fmt.Println(a.Quicksort(array))

}

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	fmt.Println(tmp)
	fmt.Println(arr)




	//The copy function supports copying between slices of different lengths (it will copy only up to the smaller number of elements)
	//
	//The usual example is:

	t := make([]byte, len(s), (cap(s)+1)*2)
	copy(t, s)
	s = t
}

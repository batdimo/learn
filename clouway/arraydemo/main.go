package main

import (
	"github.com/drafailov/dimitar_rafailov_go_intro/arrayutil"
	"math"
	"fmt"
)

func main() {
	var arr = [] int{23, -23, 0, 2, math.MinInt16, math.MaxInt16}
	fmt.Println(arrayutil.MinElement(arr))
	fmt.Println(arrayutil.Sum(arr))
	arrayutil.Print(arr)
	fmt.Println(arrayutil.Sort(arr))
	fmt.Println(arrayutil.Reverse(arr))
}

package arrayOperations

import "math"

func MinElement(arr []int) int {
	min := math.MaxInt32
	for _, e := range arr {
		if min > e {
			min = e
		}
	}
	return min
}

func Sum(arr []int) int {
	sum:=0
	for _, e:= range arr {
		sum+=e
	}
return sum
}

func Print(arr []int)int {
	for _, e:=range arr{
		return e
	}
}

func Quicksort(arr []int) [] int  {
	pivot := arr[0]
	i:=1
	j:=1
	for i != j {
		for arr [i] < pivot && i!=j{
			i++
		}
		for arr[j] >= pivot && i!=j{
			j--
		}
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[j] >= pivot {
		j--
	}


	return arr
}
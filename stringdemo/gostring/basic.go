package gostring

import (
	"math/rand"
	"time"
)
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Str(leng int) [] string {
	var arr [94] string
	for i := 0; i < 94; i++ {//33 do 127
		arr[i] = string(i+32)
		//fmt.Print(arr[i])
	}
	b := make([]string,leng)
	for i := range b {
		b[i] = arr[rand.Intn(len(arr))]
	}
	return b
}
package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("is it far Saturday")
	today:=time.Now().Weekday()
	switch time.Saturday {
	case today+0:
		fmt.Print("it's today")
	case today +1:
		fmt.Print("it's tomorow")
	default:
		fmt.Print("it's far")

	}
}
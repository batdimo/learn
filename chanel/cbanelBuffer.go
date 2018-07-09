package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main()  {
	massages := make(chan string,2)
	massages <- Reverse("1234567890")
	massages <- "chanell"
	fmt.Println(<-massages)
	fmt.Println(<-massages)
}
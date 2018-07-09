package main

import (
	"log"
	"fmt"
)

func main() {
	fmt.Println("this is me1")
	//os.Exit(1)
	log.Println("This is a regular message.")
	fmt.Println("this not me")
	log.Fatalln("This is a fatal error.") //Writes a message to os.Stderr and then exits with an error code calls os.Exit(1)


	fmt.Println("this is me2")
	log.Println("This is the end of the function.")
	fmt.Println("this is me3")
}

package main

import (
	"log"
	"os"
)

//Creates a logger
func main() {
	logfile, _ := os.Create("./log.txt")
	defer
		logfile.Close() //	Creates a log file	Makes sure it	gets closed
	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)
	//Sends it some	messages
	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
	logger.Println("This is the end of the function.")
	//As before, this will 	never get called.
}

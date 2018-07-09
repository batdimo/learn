package main

import (
	"fmt"
	"os"
)

var path = "/home/clouway/drafailov/go/src/github.com/batdimo/go_io/cmd/writefile/ad.txt"

func main() {
	createFile()
	writeFile()
	readFile()
	//deleteFile()
}

func createFile() {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err) //okay to call os.exit()
		defer file.Close()
	}
}

func writeFile() {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// write some text to file
	_, err = file.WriteString("halo\n")
	if err != nil {
		fmt.Println(err.Error())
		return //must return here for defer statements to be called
	}
	_, err = file.WriteString("mari belajar golang\n")
	if err != nil {
		fmt.Println(err.Error())
		return //same as above
	}

	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
		return //same as above
	}
}

func readFile() {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// read file
	var text = make([]byte, 1024)
	n, err := file.Read(text)
	if n > 0 {
		fmt.Println(string(text))
	}
	//if there is an error while reading
	//just print however much was read if any
	//at return file will be closed
}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
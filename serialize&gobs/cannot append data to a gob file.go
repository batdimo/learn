package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

type Struct1 struct {
	IntVal    int
	StringVal string
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func readData(filePath string) {
	if !PathExists(filePath) {
		return
	}
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("os.Open('%s') failed with '%s'\n", filePath, err.Error())
	}
	dec := gob.NewDecoder(file)
	var val Struct1
	for {
		err = dec.Decode(&val)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatalf("dec.Decode() failed with '%s'\n", err.Error())
		}
		fmt.Printf("Read: %#v\n", val)
	}
}

func addData(filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("os.OpenFile('%s') failed with '%s'\n", filePath, err.Error())
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	for i := 0; i < 1; i++ {
		n := int(rand.Intn(65536))
		v := Struct1{
			IntVal:    n,
			StringVal: fmt.Sprintf("%d", n),
		}
		err = enc.Encode(v)
		if err != nil {
			log.Fatalf("enc.Encode() failed with '%s'\n", err.Error())
		}
		fmt.Printf("Wrote: %#v\n", v)
	}
}

func main() {
	filePath := "test.gob"
	readData(filePath)
	addData(filePath)
	readData(filePath)
	addData(filePath)
	readData(filePath)
}



//Shows that one cannot append data to a gob file - the decoder will error out with 'extra data in buffer` when trying to decode data written by append.
package main
import (
	"log"
	"net"
)
//Connects to the log server
func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()
	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
	//Makes sure you clean	up by closing the	connection, even on	panic of the defer logger.Fatalln exit and miss desfer
	//Sends log messages to	the network connection 	Logs a message and then
	//panics–don’t use Fatalln here.
}
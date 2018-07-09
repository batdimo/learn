package main





import (
"fmt"
)

type Number struct {
	X int // X is an integer
	Y int // Y is an integer
}

func main() {
	m := Number{3, 4}

	// Prints 3
	fmt.Println(m.X)
}

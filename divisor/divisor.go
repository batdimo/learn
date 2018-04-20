package divisor

//import "fmt"

func Div(a, b int) int {
	var k int
	for b != 0 {
		k = a % b
		a = b
		b = k
	}
	return a
}


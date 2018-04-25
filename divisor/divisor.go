package divisor

//Div greatest common divisor from two integer numbers
func Div(a, b int) int {
	if a&b == 0  {
		return 0
	} else {
		var k int
		for b != 0 {
			k = a % b
			a = b
			b = k
		}
		return a
	}
}

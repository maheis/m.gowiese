package calculator

// Privat
//const pi float64 = 3.141592

// Öffentlich
const Pi float64 = 3.141592

func Add(a int, b int) int {
	var sum int = a + b
	// var sum = a + b
	// sum := a + b //Syntax für neue Variable mit direkter Zuweisung!

	return sum
}

func Sub(a int, b int) int {
	return a - b
}

// Zwei Rückgabetypen mit name
func Div(a int, b int) (quotient int, rest int) {
	quotient = a / b
	rest = a % b

	return
}

// Package calculator bietet grundlegende Rechenoperationen an.
package calculator

// pi mit kleinem p Privat
//const pi float64 = 3.141592

// Pi mit großem P Öffentlich
const Pi float64 = 3.141592

func Add(a int, b int) int {
	var sum = a + b
	// var sum = a + b
	// sum := a + b //Syntax für neue Variable mit direkter Zuweisung!

	return sum
}

func Sub(a int, b int) int {
	return a - b
}

// Div führt eine Division durch und gibt den Quotienten und den Rest zurück.
func Div(a int, b int) (quotient int, rest int) {
	quotient = a / b
	rest = a % b

	return
}

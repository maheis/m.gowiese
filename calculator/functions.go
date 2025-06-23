// Package calculator bietet grundlegende Rechenoperationen an.
package calculator

import (
	"fmt"
	"math"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// SquareRoot berechnet die Quadratwurzel eines Wertes.
func SquareRoot(value float64) (float64, bool) {
	if value < 0 {
		return 0, false
	}
	return math.Sqrt(value), true
}

// SquareRootWithError berechnet die Quadratwurzel eines Wertes und gibt einen Fehler zurück, wenn der Wert negativ ist.
func SquareRootWithError(value float64) (float64, error) {
	if value < 0 {
		return 0, fmt.Errorf("no square root of negative value (%f)", value)
		// return 0, errors.New("there is no square root of negative value")
	}
	return math.Sqrt(value), nil
}

func IsSquareNumber(x int) bool {
	sqrt := math.Sqrt(float64(x))

	if sqrt != float64(int(sqrt)) {
		fmt.Println("sqrt: ", false)
		return false
	}

	//Verkürzte Schreibweise, wenn Variable nur innerhalb des if-Blocks benötigt wird
	//if sqrt := math.Sqrt(float64(x)); sqrt != float64(int(sqrt)) {

	fmt.Println("sqrt: ", true)
	return true
}

func RunOperations(operation string, left int, right int) int {
	switch operation {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		fmt.Println("Operation not supported")
		return 0
	}
}

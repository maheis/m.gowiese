// Package calculator bietet grundlegende Rechenoperationen an.
package calculator

// Sum berechnet die Summe der Ganzzahlen von start bis end.
func Sum(start int, end int) int {
	sum := 0

	for i := start; i <= end; i++ {
		sum += i
	}

	return sum
}

// SumWhile berechnet die Summe der Ganzzahlen, solange die Summe kleiner als max ist.
func SumWhile(max int) int {
	sum := 1

	for sum < max {
		sum += sum
	}

	return sum
}

// SumDoWhile berechnet die Summe der Ganzzahlen, bis die Summe größer als max ist.
func SumDoWhile(max int) int {
	sum := 1

	for {
		sum += sum

		if sum > max {
			break
		}
	}

	return sum
}

// SumInfinite berechnet die Summe der Ganzzahlen in einer unendlichen Schleife.
func SumInfinite() {
	for {
		// ..

		break //bricht die Schleife ab!
		//break in einem If-Block erzeugt eine DoWhile/Until
	}
}

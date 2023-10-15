package calculator

// Normale for-Schleife
func Sum(start int, end int) int {
	sum := 0

	for i := start; i <= end; i++ {
		sum += i
	}

	return sum
}

// DoWhile-Schleife mit der for-Schleife
func SumWhile(max int) int {
	sum := 1

	for sum < max {
		sum += sum
	}

	return sum
}

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

// Endlos- bzw. DoWhile/Until-Schleife mit der for-Schleife....
func SumInfinite() {
	for {
		// ..

		break //bricht die Schleife ab!
		//break in einem If-Block erzeugt eine DoWhile/Until
	}
}

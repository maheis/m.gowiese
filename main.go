package main

//https://www.youtube.com/watch?v=eqSjKOPt7dg

import (
	"fmt"

	"github.com/maheis/m.gowiese/calculator"
	"github.com/maheis/m.gowiese/composite"
	"github.com/maheis/m.gowiese/files"
)

func main() {
	hello()
	fmt.Println("#######")
	playground()
	fmt.Println("#######")
	examples()
	fmt.Println("#######")
	random()
	fmt.Println("#######")
	calc() // Überarbeiten!
	fmt.Println("#######")
	compo() // Überarbeiten!
	fmt.Println("#######")
	asynchron()
	fmt.Println("#######")
	_heap()
	fmt.Println("#######")
	files.Jsonfile()
}

func calc() {
	fmt.Println("42 + 21 = ", calculator.Add(42, 21))
	fmt.Println("42 - 21 = ", calculator.Sub(42, 21))
	fmt.Print("17 : 3 = ")
	fmt.Println(calculator.Div(17, 3))
	fmt.Println("Summe aller Zahlen von 1-10: ", calculator.Sum(1, 10))
	fmt.Println("Summe aller Zahlen von 1-20: ", calculator.Sum(1, 20))
	fmt.Println("Summe aller Zahlen von 1-100: ", calculator.Sum(1, 100))
	fmt.Println("Summe von 1 bis max 1000 in While: ", calculator.SumWhile(1000))
	fmt.Println("Summe von 1 bis max 1000 in DoWhile: ", calculator.SumDoWhile(1000))

	fmt.Println(calculator.SquareRoot(4))
	result, ok := calculator.SquareRoot(4)
	if ok {
		fmt.Println(result)
	}
	if !ok {
		fmt.Println("Error")
	}

	fmt.Println(calculator.SquareRootWithError(-4))
	result, err := calculator.SquareRootWithError(-4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func compo() {
	fmt.Println("42 + 21 = ", composite.Add(42, 21))

	point := composite.Point{X: 3, Y: 7}
	// point := composite.Point{3, 4}

	pointPtr := &point
	fmt.Println("Point: ", point)
	fmt.Println("Point.X: ", point.X)
	fmt.Println("Point.Y: ", point.Y)
	fmt.Println("PointPtr: ", pointPtr)
	fmt.Println("PointPtr.X: ", pointPtr.X)
	fmt.Println("PointPtr.Y: ", pointPtr.Y)

	composite.DemoCollections()

	fmt.Println("DistanceFromZero: ", point.DistanceFromZero())
	fmt.Println("DistanceFromZero: ", pointPtr.DistanceFromZero())
}

func examples() {
	fmt.Println("Hier spielen auch anderen mit...")
}

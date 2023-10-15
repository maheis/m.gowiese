package main

//https://www.youtube.com/watch?v=eqSjKOPt7dg

import (
	"fmt"
	"time"

	"github.com/maheis/m.gowiese/calculator"
	"github.com/maheis/m.gowiese/composite"
	"github.com/maheis/m.gowiese/jsonfile"
)

func main() {
	//Channel um die Asynchrone Funktion bzw. das Ergebnis wieder Synchrion zu bringen
	queue := make(chan int)
	// Der obige Channel hat einen Puffer von 1 (default) d.h. es kann ein Wert zwischengespeichert werden.
	// Wird der Puffer voll, dann blockiert der Channel bis ein Wert gelesen wurde.
	// Bzw. der Reader blockiert bis ein Wert geschrieben wurde.
	// queue = make(chan int, 3)
	// Der obige Channel hat einen Puffer von 3 d.h. es können 3 Werte zwischengespeichert werden.

	//Asynchroner Aufruf, go func ist ein "anonymer Thread"
	go func(q chan int) {
		time.Sleep(time.Second * 1)
		fmt.Println("Hello from another thread")
		q <- 23
		time.Sleep(time.Second * 1)
		q <- 42
		time.Sleep(time.Second * 1)
		q <- 17
	}(queue)

	//Synchroner Aufruf

	hello()

	// Warten auf das Ergebnis der asynchronen Funktion bzw. des Channels
	valueFromQueue := <-queue
	fmt.Println("Value from queue: ", valueFromQueue)

	// time.Sleep(time.Second * 3)
	valueFromQueue = <-queue
	fmt.Println("Value from queue: ", valueFromQueue)
	valueFromQueue = <-queue
	fmt.Println("Value from queue: ", valueFromQueue)

	playground()
	examples()
	random()
	calc()
	compo()

	fmt.Println("#######")
	myHeap()
	jsonfile.Jsonfile()
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

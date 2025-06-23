// Package composite implementiert verschiedene Datensammlungen in Go, wie Arrays, Slices und Maps.
package composite

import (
	"fmt"
)

func DemoCollections() {
	// Array mit unveränderlicher Größe
	var primesArray = [5]int{2, 3, 5, 7, 11}
	fmt.Println("primesArray: ", primesArray)

	// Slice ist ein Array mit veränderlicher Größe
	// var primesSlice []int = []int{2, 3, 5, 7, 11}
	primesSlice := make([]int, 5, 7) // mit make() wird ein Slice erzeugt und kann eine Anfangsgröße und eine Kapazität steuern
	primesSlice = append(primesSlice, 13)
	primesSlice = append(primesSlice, 17)
	primesSlice = append(primesSlice, 19)
	fmt.Println("primesSlice: ", primesSlice)
	fmt.Println(len(primesSlice))
	fmt.Println(cap(primesSlice))

	for _, i := range primesArray {
		fmt.Println(i)
	}

	for _, i := range primesSlice {
		fmt.Println(i)
	}

	// Map ist ein assoziatives Array
	points := map[string]Point{
		"A": *NewPoint(3, 7),
		"B": *NewPoint(4, 8),
		"C": *NewPoint(5, 9),
		"D": *NewPoint(6, 10),
		"E": *NewPoint(7, 11),
	}
	fmt.Println("points: ", points)
	fmt.Println("points[A]: ", points["A"])
	fmt.Println("points[B]: ", points["B"])
	fmt.Println("points[C]: ", points["C"])
	fmt.Println("points[D]: ", points["D"])
	fmt.Println("points[E] X: ", points["E"].X)
	fmt.Println("points[E] Y: ", points["E"].Y)

	//NullPointer abfangen
	nullPointer, err := points["F"]
	fmt.Println("nullPointer: ", nullPointer)
	if !err {
		fmt.Println("Fehler: ", err)
	}

	///Löschen!
	delete(points, "D")
	fmt.Println("points: ", points)

}

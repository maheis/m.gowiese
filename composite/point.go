package composite

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

// Frage: was ist sinnvoller, einen Pointer ...
func NewPoint(X, Y int) *Point {
	return &Point{X, Y}
}

// ... oder eine Instanz zu übergeben?
// func New(X, Y int) Point {
// 	return Point{X, Y}
// }

// Implementiert das Stringer interface und Formatiert die Ausgabe von fmt.Print()
// Diese Funktion darf nur auf dem Struct direkt genutzt werden, nicht auf einem Pointer!
func (point Point) String() string {
	return "(" + fmt.Sprint(point.X) + " | " + fmt.Sprint(point.Y) + ")"
}

func (p *Point) DistanceFromZero() float64 {
	return math.Sqrt(math.Pow(float64(p.X), 2) + math.Pow(float64(p.Y), 2))
}

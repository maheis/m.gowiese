// Package composite implementiert verschiedene Datensammlungen in Go, wie Arrays, Slices und Maps.
package composite

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

// NewPoint ist ein Konstruktor für die Point-Struktur.
// Frage: was ist sinnvoller, einen Pointer ...
func NewPoint(X, Y int) *Point {
	return &Point{X, Y}
}

// ... oder eine Instanz zu übergeben?
// func New(X, Y int) Point {
// 	return Point{X, Y}
// }

// String gibt eine formatierte Zeichenkette für den Punkt zurück.
// Implementiert das Stringer interface und Formatiert die Ausgabe von fmt.Print()
// Diese Funktion darf nur auf dem Struct direkt genutzt werden, nicht auf einem Pointer!
func (p Point) String() string {
	return "(" + fmt.Sprint(p.X) + " | " + fmt.Sprint(p.Y) + ")"
}

// DistanceFromZero calculates and returns the Euclidean distance of the point
// from the origin (0, 0) in a 2D Cartesian coordinate system.
func (p *Point) DistanceFromZero() float64 {
	return math.Sqrt(math.Pow(float64(p.X), 2) + math.Pow(float64(p.Y), 2))
}

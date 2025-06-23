// Package composite implementiert verschiedene Datensammlungen in Go, wie Arrays, Slices und Maps.
package composite

// Add addiert zwei Zahlen und gibt das Ergebnis zurück.
// Pointer... sinnfreies Beispiel
func Add(a int, b int) int {
	//Pointer auf Variablen erstellen
	aPtr := &a
	bPtr := &b

	//Werte der Variablem auslesen auf die verpointert wurde
	var sum = *aPtr + *bPtr

	return sum
}

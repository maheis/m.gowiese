package composite

// Pointer... sinnfreies Beispiel
func Add(a int, b int) int {
	//Pointer auf Variablen erstellen
	aPtr := &a
	bPtr := &b

	//Werte der Variablem auslesen auf die verpointert wurde
	var sum int = *aPtr + *bPtr

	return sum
}

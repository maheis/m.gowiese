package main

import "fmt"

func array() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 13}
	b := a[2:5] // [3, 4, 5] (inklusive 2, exklusive 5)

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	ptrA := &a
	fmt.Println("ptrA:", ptrA)
	// fmt.Println("ptrA[0]:", ptrA[0]) <- das geht nicht! Der Pointer muss dereferenziert werden.
	fmt.Println("c:", (*ptrA)[0]) // das geht wenn es in () gesetzt wird ...
	c := *ptrA                    // ... oder führ Mehrere Zugriffe in eine Variable dereferenziert wird
	fmt.Println("c:", c[0])

	//Itterieren durch den Array, der Index i wird dabei durch den _ ignoriert!
	for _, num := range a {
		fmt.Print(num, ", ")
	}
	fmt.Println()

	for i := range a {
		fmt.Print(a[i], ", ")
	}
	fmt.Println()

	//Itterieren durch den Array,
	for i, num := range a {
		fmt.Print(i, ": ", num, ", ")
	}
	fmt.Println()
}

package main

import "fmt"

func array() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[2:5] // [3, 4, 5] (inklusive 2, exklusive 5)

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	ptrA := &a
	fmt.Println("ptrA:", ptrA)
	// fmt.Println("ptrA[0]:", ptrA[0]) <- das geht nicht! Der Pointer muss dereferenziert werden.
	c := *ptrA
	fmt.Println("c:", c[0])
}

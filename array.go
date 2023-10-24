package main

import "fmt"

func array() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[2:5] // [3, 4, 5] (inklusive 2, exklusive 5)

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	//dereferenzieren wenn das es nur einen Pointer auf das array gibt
}

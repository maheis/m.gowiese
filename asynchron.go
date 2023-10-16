package main

import (
	"fmt"
	"time"
)

func asynchron() {
	// Channel um die Asynchrone Funktion bzw. das Ergebnis wieder Synchrion zu bringen
	queue := make(chan int)
	// Der obige Channel hat einen Puffer von 1 (default) d.h. es kann ein Wert zwischengespeichert werden.
	// Wird der Puffer voll, dann blockiert der Channel bis ein Wert gelesen wurde.
	// Bzw. der Reader blockiert bis ein Wert geschrieben wurde.
	// queue = make(chan int, 3)
	// Der obige Channel hat einen Puffer von 3 d.h. es können 3 Werte zwischengespeichert werden.

	// Asynchroner Aufruf, go func ist ein "anonymer Thread"
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

	playground()

	time.Sleep(time.Second * 3)
	valueFromQueue = <-queue
	fmt.Println("Value from queue: ", valueFromQueue)
	valueFromQueue = <-queue
	fmt.Println("Value from queue: ", valueFromQueue)

}

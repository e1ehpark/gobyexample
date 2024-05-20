package main

import (
	"fmt"
	"sync"
)

/*
Problem Statement:
 Design and implement a Golang program that prints
 numbers from 1 to 5 in the correct sequence,
 with even numbers printed by one goroutine and
 odd numbers printed by a separate goroutine.
*/

var wg sync.WaitGroup

func main() {
	evenCh, oddCh := make(chan bool, 1), make(chan bool, 1)
	defer close(evenCh)
	defer close(oddCh)

	wg = sync.WaitGroup{}
	wg.Add(2)

	go even(evenCh, oddCh)
	go odd(oddCh, evenCh)

	// initiate the odd function to start first
	evenCh <- true

	wg.Wait()
}

func even(evenCh, oddCh chan bool) {
	for i := 2; i <= 5; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- true
	}

	wg.Done()
}

func odd(oddCh, evenCh chan bool) {
	for i := 1; i <= 5; i += 2 {
		<-evenCh
		fmt.Println(i)
		oddCh <- true
	}

	wg.Done()
}

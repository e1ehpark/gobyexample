package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	messages <- "buffered"
	messages <- "channel"
	messages <- "3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
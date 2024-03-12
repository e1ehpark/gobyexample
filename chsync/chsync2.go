package main

import "fmt"

func produce(numbers chan int) {
	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers)
}

func main() {
	numbers := make(chan int)

	go produce(numbers)

	for number := range numbers {
		fmt.Println(number)
	}
}
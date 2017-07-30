package main

import (
	"fmt"
	"os"
)

func counter(output chan<- int) {
	for i := 0; i < 300; i++ {
		output <- i
	}
	close(output)
}

func multiply( output chan<- int, input <-chan int) {
	for i := range input {
		output <- i*i
	}
	close(output)
}

func outputValues(input <- chan int) {
	for i := range input {
		fmt.Println("Squared value: ", i)
	}
}


func main() {
	regular   := make(chan int)
	mult := make(chan int)
	go counter(regular)
	go multiply(mult,regular)
	outputValues(mult)
	os.Exit(0)
}
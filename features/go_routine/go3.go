package main

import (
	"fmt"
	"os"
)

func DisplayValues(over chan<- int) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%x\n", i)
	}
	over <- 0
}


func main() {
	rt := make(chan int)

	for c := 0; c < 5; c++ {
		go DisplayValues(rt)
	}

	for e := 0; e < 5; e++ {
		fmt.Println(<-rt)
	}
	os.Exit(0)
}
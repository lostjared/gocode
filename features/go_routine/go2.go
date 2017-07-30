package main

import(
	"fmt"
	"os"
)

func printOut(ch chan<- int) {
	c := 0
	for i := 0; i < 600; i++ {
		fmt.Println(i)
		c = i
	}
	ch <- c
}


func main() {
	val := make(chan int)
	go printOut(val)
	value := <- val // wait till go routine finishes
	fmt.Println("Program returned ", value)
	os.Exit(0)
}
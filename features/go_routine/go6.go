package main

import(
	"fmt"
	"os"
	"sync"
)

var mutex sync.Mutex

func printData(x int, channel chan <- int) {
	for i := 0; i < 100; i++ {
		mutex.Lock()
		fmt.Println("In goroutine: ", x, " loop index: ", i)
		mutex.Unlock()
	}
	channel <- 1
}

func main() {
	channel := make(chan int)
	for i := 0; i < 100; i++ {
		go printData(i, channel)
	}
	for i := 0; i < 100; i++ {
		<- channel
	}
	os.Exit(0)
}
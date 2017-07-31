package main

import (
	"fmt"
	"os"
	"sync"
	"time"
	"math/rand"
)

func procData() {
	channel := make(chan string)
	var wt sync.WaitGroup
	for i := 0; i < 100; i++ {
		wt.Add(1)
		go func(x int) {
			defer wt.Done()
			time.Sleep((time.Duration(rand.Intn(100)) * time.Millisecond))
			fmt.Println("Go routine.. ")
			channel <- fmt.Sprintf("Value %d", x)
		}(i)
	}
	go func() {
		wt.Wait()
		close(channel)
	}()

	for svalue := range channel {
		fmt.Println("Value: ", svalue)
	}

	fmt.Println("Over")
}

func main() {
	fmt.Println("Go routine #4")
	procData()
	os.Exit(0)
}
/*

Concurrent nested for loop
It is much faster without the creation of a go routine every index
would be better to break the data up into pieces then have maybe 4-5 different go routines

*/


package main

import(
	"fmt"
	"os"
	"sync"
	"time"
	"math/rand"
)

func procData() {
	channel := make(chan string)
	var w sync.WaitGroup
	for i := 0; i < 320; i++ {
		for z := 0; z < 240; z++ {
			w.Add(1)
			go func(x,y int, pix int) {
				defer w.Done()
				time.Sleep( time.Duration(rand.Intn(10)) * time.Millisecond)
				channel <- fmt.Sprintf("%d,%d Set to: %d", x,y,pix*pix)
			}(i, z, i*z)
		}
	}
	go func() {
		w.Wait()
		close(channel)
	}()
	for string_value := range channel {
		fmt.Println("Value returned: ", string_value)
	}
}

func main() {
	fmt.Println("Concurrent nested proc proc")
	procData()
	os.Exit(0)
}
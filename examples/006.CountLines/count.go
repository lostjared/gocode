package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Error requires at least one argument.\n")
		os.Exit(1)
	}
	channel := make(chan int64)
	file_channel := make(chan string)
	for _, filen := range os.Args[1:] {
		go countLines(filen,channel,file_channel)
	}
	var total int64 = 0
	for  range os.Args[1:] {
		count := <- channel
		fmt.Printf("%s contains %d lines\n", <- file_channel, count)
		total += count
	}
	fmt.Printf("All files contain %d lines.\n", total)
}

func countLines(filename string, channel chan<- int64, file_name chan<- string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error could not open file: %v", err)
		return
	}
	input := bufio.NewScanner(f)
	var counter int64 = 0
	for input.Scan() {
		counter++
	}
	channel <- counter
	file_name <- filename
}
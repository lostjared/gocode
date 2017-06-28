package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	echo_file := os.Args[1:]
	if len(echo_file) == 0 {
		fmt.Printf("Please provide filename\n")
	} else {
		f, err := os.Open(echo_file[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error could not open file.\n")
			return
		}

		input := bufio.NewScanner(f)
		for input.Scan() {
			fmt.Printf("%s\n", input.Text())
		}
	}
}

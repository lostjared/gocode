package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s", "Requires one argumnet")
		os.Exit(1)
	}
	fptr, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open file: %s", os.Args[1])
		os.Exit(1)
	}
	input := bufio.NewScanner(fptr)
	inputValues := make(map[string]int)
	for input.Scan() {
		text := input.Text()
		_, ok := inputValues[text]
		if(!ok) {
			fmt.Printf("%s\n", text)
			inputValues[text] = 1
		}
	}
}

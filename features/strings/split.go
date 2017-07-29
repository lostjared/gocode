package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func readAndTokenize(source string) []string {
	fptr, err := os.Open(source)
	if(err != nil) {
		fmt.Fprintln(os.Stderr, "Error could not open: ", source, " Error: ", err)
		os.Exit(1)
	}

	s, err := ioutil.ReadAll(fptr)
	if(err != nil) {
		fmt.Fprintln(os.Stderr, "Error could not read file: ", source, " Error: ", err)
		os.Exit(1)
	}
	fptr.Close()
	return strings.Split(string(s), " ")
}


func main() {
	if (len(os.Args) != 2) {
		fmt.Fprintln(os.Stderr, "Error requires one argument of the file to open.\n")
		os.Exit(1)
	}
	s := readAndTokenize(os.Args[1])
	fmt.Println("Tokens sep by space")
	index := 1
	for _, i := range(s) {
		fmt.Println(index, ": ", i)
		index ++
	}
	os.Exit(0);
}

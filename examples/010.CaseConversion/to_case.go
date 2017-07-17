package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
)


func Echo(text string, up_l bool) {
	var temp string
	if(up_l == true) {
		temp = strings.ToUpper(text)
	} else {
		temp = strings.ToLower(text)
	}
	fmt.Printf("%s", temp)
}


func main() {
	if(len(os.Args) != 3) {
		fmt.Fprintln(os.Stderr, "Error requires text file as argument. as well as -u  or -l")
		os.Exit(1)
	}
	up_l := false
	if(os.Args[2] == "-u") {
		up_l = true

	} else if(os.Args[2] == "-l") {
		up_l = false

	} else {
		fmt.Fprintln(os.Stderr, "Error requires -u (upper) or -l (lower) case")
		os.Exit(1)
	}

	fn,err := os.Open(os.Args[1])
	if (err != nil) {
		fmt.Fprintln(os.Stderr, "Error could not open file: ", err)
		os.Exit(1)
	}
	text,err := ioutil.ReadAll(fn)
	if(err != nil) {
		fmt.Fprintln(os.Stderr, "Error reading from: ", os.Args[1])
		os.Exit(1)
	}
	Echo(string(text), up_l)
	fn.Close()
}
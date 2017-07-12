package main

import (
	"fmt"
	"os"
	"io/ioutil"
)


func ReadText(filename string) (string,bool) {
	f,err := os.Open(filename);
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", false
	}
	s,err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	defer f.Close()
	return string(s),true
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error requires 1 argument of file to read")
		os.Exit(1)
	}
	s, err := ReadText(os.Args[1])
	if err == true {
		fmt.Println(s);
	}
	defer func() {
		fmt.Println("Defered until after return\n")
	}()
}
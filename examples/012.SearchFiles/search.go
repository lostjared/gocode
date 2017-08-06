package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"sync"
)

var mut sync.Mutex


func SearchFile(filename string, search_string string, done chan<- int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not open: ", filename)
		return
	}
	defer f.Close()
	file_data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Error could not read file: ", err)
		return
	}
	if strings.Contains(string(file_data), search_string) {
		fmt.Println(filename, ": Search String Found")
	} else {
		fmt.Println(filename, ": Search String Not Found")
	}
	done <- 0
}

func SearchFiles(search []string, search_string string) {

	ch := make(chan int)

	for _, i := range search {
		go SearchFile(i, search_string, ch)
	}

	for i := 0; i < len(search); i++ {
		<-ch
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Requires at least two arguments")
		os.Exit(0)
	}
	search_string := os.Args[1]
	var search []string
	for i := 2; i < len(os.Args); i++ {
		search = append(search, os.Args[i])
	}
	SearchFiles(search[:], search_string)
}
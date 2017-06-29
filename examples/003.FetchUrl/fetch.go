package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error requires one argument containing path")
		os.Exit(1)
	}
	fetchUrl(os.Args[1])
}

func fetchUrl(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error could not access URL")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Could not read bytes from url")
		os.Exit(1)
	}
	fmt.Printf("%s", bytes)
}

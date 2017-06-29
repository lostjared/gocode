package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	channel := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchUrl(url, channel)
	}

	for range os.Args[1:] {
		fmt.Println(<- channel)
	}
}

func fetchUrl(url string, channel chan<- string) {
	response, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprint(err)
		return
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		channel <- fmt.Sprintf("Error reading %s : %v", url, err)
	}
	channel <- fmt.Sprintf("URL: %s\n%s\n", url, bytes)
}
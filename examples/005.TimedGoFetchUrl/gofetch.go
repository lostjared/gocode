package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start_time := time.Now()
	channel := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchUrl(url, channel)
	}

	for range os.Args[1:] {
		fmt.Println(<- channel)
	}
	if len(os.Args) > 1 {
		fmt.Printf("%.2f seconds elapsed \n", time.Since(start_time).Seconds())
	}
}

func fetchUrl(url string, channel chan<- string) {
	start_time := time.Now()
	response, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprint(err)
		return
	}
	bytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		channel <- fmt.Sprintf("Error reading %s : %v", url, err)
	}
	channel <- fmt.Sprintf("URL: %s Download time took: %.2f seconds \n%s\n",url,time.Since(start_time).Seconds(),bytes)
}

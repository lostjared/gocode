package main

import(
	"fmt"
	"os"
	"bufio"
)


func main() {

	fmt.Println("Panic example")

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {

		value := scan.Text()
		if(value == "exit") {
			os.Exit(0)
		} else {
			panic("You didn't input a valid command")
		}

	}

	os.Exit(0)
}
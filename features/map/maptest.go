// Using simple map (hash table)
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// will scan input if same text is entered will increase variable
	// if print is entered it will print out the keys
	fmt.Println("Map example")
	value := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "print" {
			for key, value := range value {
				fmt.Println("Key = ", key, " Value = ", value)
			}
			continue
		} else if text == "quit" {
			break
		}
		value[text]++
	}
	fmt.Println("Loop exited")
	string_values := map[string]string {
		"Jared": "Programmer",
		"Go": "Cool language",
	}

	for key,value := range string_values {
		fmt.Println("Key = ", key, " Value = ", value)
	}
}
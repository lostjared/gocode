package main

import(
	"fmt"
	"os"
)

func main() {
	fmt.Println("Array example")
	// use  := to init
	arr := [...]int{0,10, 20, 30}
	// range for loop over array
	for _, i := range(arr) {
		fmt.Println("Value is: ", i)
	}
	// pass array as function argument
	outputArray(arr)

	// string array
	var str_array [4]string

	// set values for string array
	str_array = [4]string{"Hey ", "One", "Two", "Three"}

	// range based for loop
	for _, i := range(str_array) {
		fmt.Println("String value is: ", i)
	}
	os.Exit(0)
}

// pass as argument to function
func outputArray(value [4]int) {
	// regular for loop
	for i := 0; i < len(value); i++ {
		fmt.Printf("loop value: %d\n", value[i])
	}
}
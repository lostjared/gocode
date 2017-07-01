package main

import(
"fmt"
)

func main() {
	values := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	for i := 0; i <= 9; i++ {
		printNumber(i, values)
	}

	x := 100
	y := 200
	fmt.Printf("%d\n", (x+y))
	a, b := 0, 10
	fmt.Printf("%d:%d\n", a, b)
}


func printNumber(index int, values []string) {
	if index >= 0 && index <= 9 {
		fmt.Printf("%s\n", values[index])
	}
}


package main

import (
	"fmt"
	"os"
)
// simple function two args
func addValues(x int, y int) int {
	return (x+y)
}

// multiple return values
func returnValues(x int, y int) (int, int) {
	return x+1,y+1
}

// takes pointer
func ptrManip(s *int) {
	*s = 100
}

func recursionCounter(x int) {
	if x == 10 {
		return
	}
	fmt.Println("Counting: ", x)
	recursionCounter(x+1)

}

func main() {
	fmt.Println("Function Example")
	value := addValues(100, 200)
	x,y := returnValues(100, 200)
	fmt.Println("value: ", value, " x: ", x, " y: ", y)
	s := 10
	fmt.Println("Value before: ", s)
	ptrManip(&s)
	fmt.Println("Value after: ", s)
	recursionCounter(0)
	os.Exit(0)
}




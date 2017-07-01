package main;

import (
"fmt"
)

type LargeInteger int64

func main() {

	var value LargeInteger = 5000
	value = add(value, value)
	fmt.Printf("Value is: %d\n", value)
}

func add(x LargeInteger, y LargeInteger) LargeInteger {
	return (x+y)
}

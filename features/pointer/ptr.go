package main

import (
"fmt"
)

func main() {
	var squared float64 = 100.5
	square(&squared)
	fmt.Printf(" 100.5  squared is: %f \n", squared);
	var ptr *int64;
	ptr = return_Ptr();
	*ptr += 100;
	fmt.Printf("100 + 100 = %d\n", *ptr);
	new_value := return_NewPtr();
	*new_value = 300;
	fmt.Printf("%d\n", *new_value);
}


// pass pointer as argument
func square(x *float64) {
	*x = *x**x;
}

//  return pointer
func return_Ptr() *int64 {
	var value int64 = 100;
	return &value;
}

// return pointer using new
func return_NewPtr() *int64 {
	return new(int64);
}


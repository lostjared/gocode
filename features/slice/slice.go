package main

import (
	"fmt"
	"os"
)


func main() {

	fmt.Println("Slice example")

	values := [...]int{0, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	partial_after := values[4:]
	partial_before := values[:4]

	for _,v := range(values) {
		fmt.Println("Values = ", v)
	}

	for _,i := range(partial_after) {
		fmt.Println("Value after index [4:] ", i)
	}

	for _,i := range(partial_before) {
		fmt.Println("Value before index [:4] ", i)
	}

	addTo(values[:])

	os.Exit(0)
}

func addTo(val []int) {
	for i := 2048; i < 100000; i = i*2 {
		val = append(val, i)
	}
	fmt.Println("Values appended are = {")
	for _,i := range(val) {
		fmt.Println(i)
	}
	fmt.Println("}");
}


// sort integers

package main

import (
	"fmt"
	"sort"
)

func main() {

	lst := []int {5, 0, 6, 1, 3, 9, 4}
	sort.Ints(lst)
	fmt.Println("Values sorted: {")
	for _, i := range(lst) {
		fmt.Printf("%d, ", i)
	}
	fmt.Println("\n}")
}
/* Program reads file, and removes empty lines then outputs to stdout */

package main

import(
	"os"
	"fmt"
	"bufio"
)

func checkPrint(value string) {
	check := false
	for _, i  := range(value)  {
		if i != ' ' && i != '\t' && i != '\n' {
			check = true
			break
		}
	}
	if check == true {
		fmt.Println(value)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stdin, "Requires one argument of text file")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stdin, "Error could not open file: ", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		checkPrint(input.Text())
	}
	f.Close()
}

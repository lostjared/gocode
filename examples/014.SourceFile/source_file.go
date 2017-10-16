/*
		Print out Source code with Line Numbers
*/



package main

import (
	"fmt"
	"os"
	"bufio"
)

type Source_File struct {
	code []string
	size uint
}


func(c *Source_File) loadSource(filename string) uint {
	f,err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error could not open file.")
		return 0
	}

	defer f.Close()
	input_ := bufio.NewScanner(f)
	var lines = []string{}
	len := uint(0)
	for input_.Scan() {
		lines = append(lines, input_.Text())
		len++
	}
	c.size = len
	c.code = lines
	return c.size
}

func(c Source_File) atLine(index uint) string {
	return c.code[index];
}

func main() {
	if(len(os.Args) != 2) {
		fmt.Fprintln(os.Stderr, "Error requires one argument of source file")
		os.Exit(1)
	}
	the_code := Source_File{}
	len := the_code.loadSource(os.Args[1])
	if(len == 0) {
		fmt.Fprintln(os.Stderr, "Error could not load file..\n")
		os.Exit(1)
	}
	for i := uint(0); i < the_code.size; i++ {
		fmt.Printf("%d:\t%s\n", i, the_code.atLine(i))
	}
}
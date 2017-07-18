package main

import (
	"fmt"
)


type Symbols string

type Output interface {
	Write()
}

type Value struct {
	value string
}

type Incorrect struct {
	x,y,z int
}

type IntValue struct {
	value int
}

func (v Value) Write() {
	fmt.Println(v.value)
}

func (v IntValue) Write() {
	fmt.Println(v.value)
}

func(s Symbols) Write() {
	fmt.Println(s)
}

func main() {
	fmt.Println("Interface ..\n")
	value := Value{"Hello World"}
	intvalue := IntValue{10}
	intvalue.value = 0xDEADBEEF
	var o Output
	o = value
	o.Write()
	o = intvalue
	o.Write()
	var stri Symbols = "Hey whats up?"
	o = stri
	o.Write()

// below will not work because it does not implement the Write method for interface Output
//	var err Incorrect = Incorrect{10,20,30}
//	o = err

}
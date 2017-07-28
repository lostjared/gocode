package main

import (
	"fmt"
	"os"
)

type Does interface {
	walk()
	stop()
}

type Person interface {
	Does
	name() string
	says(x string)
}

type Jared struct {
	my_name string
}

func (j Jared) name() string {
	return j.my_name
}

func (j Jared) says(x string) {
	fmt.Println("Jared: ", x)
}

func (j Jared) walk() {
	fmt.Println("Jared: ", "[walk]")
}

func (j Jared) stop() {
	fmt.Println("Jared: ", "[stop]")
}

func main() {
	fmt.Println("Type assertions")
	var inter Person
	jared := Jared{"Jared Bruni"}
	inter = jared
	// yes it contains Does
	if s, ok := inter.(Does); ok {
		fmt.Println(s, ":", ok, "!")
	} else {
		fmt.Println(s, ":", ok);
	}
	type stringSay interface {
		name() string
	}
	type stringNo interface {
		sayNo(x string)
	}
	// yes it has name
	if s, ok := inter.(stringSay); ok {
		fmt.Println(s, ":", ok, "!")
	} else {
		fmt.Println(s, ":", ok);
	}
	// should return false
	if s, ok := inter.(stringNo); ok {
		fmt.Println(s, ":", ok, "!")
	} else {
		fmt.Println(s, ":", ok);
	}
	os.Exit(0)
}
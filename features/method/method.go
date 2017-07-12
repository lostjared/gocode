package main


import (
	"fmt"
	"os"
)

type Value struct {
	x,y int
	name string
}

// print values
func(val Value) PrintValue() {
	fmt.Println("X: ", val.x, " Y: ", val.y, "\n", val.name)
}

// get value
func(val Value) GetStringValue() string {
	return val.name
}

// change values in structure
func(val *Value) setValue(xvalue int, yvalue int, sv string) {
	val.x = xvalue
	val.y = yvalue
	val.name = sv
}

func main() {
	fmt.Println("Method ex")
	value := Value{x: 100, y:100, name:"Text"}
	value.PrintValue()
	fmt.Println(value.GetStringValue())
	value.setValue(100, 200, "Jared")
	value.PrintValue()
	os.Exit(0)
}
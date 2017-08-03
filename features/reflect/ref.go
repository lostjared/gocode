package main

import (
	"fmt"
	"os"
	"reflect"
)

func IntOrString(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Invalid:
		fmt.Println("Invalid type")
		return false
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		fmt.Println("Its an integer value: ", value.Int())
		return true
	case reflect.String:
		fmt.Println("Its a string value: ", value.String())
		return true
	default:
		fmt.Println("Not a integer or a string")
		return false
	}
	return false
}

func main() {
	fmt.Println("Refelct example")
	var x int
	x = 100
	IntOrString(reflect.ValueOf(x))
	var s string
	s = "Hello Reflection"
	IntOrString(reflect.ValueOf(s))
	var f float64
	f = 100.0
	IntOrString(reflect.ValueOf(f))
	os.Exit(0)
}
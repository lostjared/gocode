package main

import (
	"fmt"
	"os"
)

type Point struct {
	x,y int
}


func (p* Point) setPoint (x,y int) {
	p.x = x
	p.y = y
}

func (p Point) PrintPoint() {
	fmt.Println("Point: (", p.x, ",", p.y, ")")
}

func (p* Point) AddPoint(px Point) {
	p.x +=  px.x
	p.y +=  px.y
}



func main() {
	fmt.Println("Point example")
	x := Point{100,100}
	x.PrintPoint()
	x.setPoint(200, 300)
	x.PrintPoint()
	y := Point{400, 500}
	fmt.Println("Added two: ")
	x.PrintPoint()
	fmt.Println("+")
	y.PrintPoint()
	fmt.Println("=")
	y.AddPoint(x)
	y.PrintPoint()
	os.Exit(0)
}
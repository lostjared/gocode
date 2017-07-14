package point

import (
	"fmt"
)

type Point struct {
	X,Y int
}


func (p* Point) setPoint (x,y int) {
	p.X = x
	p.Y = y
}

func (p Point) PrintPoint() {
	fmt.Println("Point: (", p.X, ",", p.Y, ")")
}

func (p* Point) AddPoint(px Point) {
	p.X +=  px.X
	p.Y +=  px.Y
}

func createPoint(x, y int) Point {
	return Point{x,y}
}




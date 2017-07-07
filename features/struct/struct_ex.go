package main

import (
	"fmt"
	"os"
)

type Point struct {
	x,y int
}

type Rect struct {
	Point
	w,h int
}

type Person struct {
	Rect
	name string
}

func main() {
	p := Point{100, 200}
	z := Point{x:100, y:200}
	if p == z {
		fmt.Println("Values are equal\n")
	}
	r := Rect {Point{0, 10}, 100, 200}
	if(r.x == 0 && r.y == 10) {
		fmt.Println("Values: ", r.x, " ", r.y, " ", r.w, " ", r.h)
	}
	jared := createPerson(Rect{Point{0, 0}, 100, 200}, "Jared ")
	PrintPerson(&jared)
	os.Exit(1)
}

func createPerson(r Rect, name string) Person {
	p := Person{r, name}
	return p
}

func PrintPerson(p *Person) {
	fmt.Printf("%s {%d,%d,%d,%d}\n", p.name, p.Rect.x, p.Rect.y, p.Rect.w, p.Rect.h)
}

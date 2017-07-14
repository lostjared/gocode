package main


import ( 
	"fmt"
	"github.com/lostjared/gocode/features/point_package"
)



func main() {
	fmt.Println("Hey world")
	p := point.Point{100, 200}
	p.PrintPoint()	
}

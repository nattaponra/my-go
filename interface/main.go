package main

import (
	"fmt"

	"github.com/nattaponra/my-go/interface/geometry"
)

//Interface
type Geometry interface {
	Area() float64
	Perim() float64
}

func Measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.Area())
	fmt.Println("Perim:", g.Perim())
}

func main() {
	Measure(geometry.Rect{Height: 10, Width: 10})
	Measure(geometry.Circle{Radius: 50})
}

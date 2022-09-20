package interfaces

import (
	"fmt"
	"math"
)

// shape=şekil
type shape interface {
	area() float64
	//
	//
}

// dikdörtgen
type rectangle struct {
	width, height float64
}

// daire
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println(s.area())
}

func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r)
}

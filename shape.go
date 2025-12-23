package main

import (
	"fmt"
	"math"
)

type shapee interface {
	area() float64
}
type circle struct {
	size float64
}
type rec struct {
	width  float64
	height float64
}

func (c circle) area() float64 {
	return math.Pi * c.size * c.size
}
func (r rec) area() float64 {
	return r.height * r.width
}
func shape() {
	c1 := circle{4.5}
	c2 := rec{2, 3}
	shapes := []shapee{c1, c2}
	for _, v := range shapes {
		fmt.Println(v.area())
	}

}

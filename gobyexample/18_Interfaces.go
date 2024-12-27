package main

import (
	"fmt"
	"math"
)

func main() {

	c1 := circle{12}
	c2 := circle{2}
	fmt.Println(c1, c1.radious)

	meassure(c1)
	meassure(c2)
}

type geometry interface {
	area() float64
}

type circle struct {
	radious float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radious * c.radious)
}

func meassure(g geometry) {
	fmt.Println(g.area())
}

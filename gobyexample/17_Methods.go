package main

import "fmt"

func main() {

	rect1 := rect{22, 24}

	fmt.Println(rect1.width)
	fmt.Println(rect1.height)

	fmt.Println(rect1.area1())
	fmt.Println(rect1.area2())

	rect2 := &rect1

	fmt.Println(rect2.area1())
	fmt.Println(rect2.area2())
}

type rect struct {
	height, width int
}

func (r rect) area1() int {
	return r.width * r.height
}

func (r *rect) area2() int {
	return r.width * r.height
}

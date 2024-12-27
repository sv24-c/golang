package main

import "fmt"

type person struct {
	name string
	id   int
}

func main() {

	fmt.Println(person{"p1", 1})

	fmt.Println(person{name: "p2", id: 2})

	fmt.Println(person{name: "p3"})

	p4 := person{name: "p4", id: 4}
	fmt.Println(p4.name, p4.id)

	p5 := p4
	fmt.Println(p5, p4, p5.name, p4.name)

	p6 := &p4
	fmt.Println(p6, p4, p6.name, p4.name)

	object := struct {
		name     string
		isObject bool
	}{"ob1", true}

	fmt.Println(object)
}

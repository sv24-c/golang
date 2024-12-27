package main

import "fmt"

func main() {

	j1 := juice(Apple)
	fmt.Println(j1)

	j2 := juice(Papaya)
	fmt.Println(j2)
}

type Fruits int

const (
	Apple Fruits = iota
	Banana
	Mango
	Papaya
)

var fruitList = map[Fruits]string{
	Apple: "a", Banana: "b",
	Mango: "m", Papaya: "p",
}

func (f Fruits) String() string {
	return fruitList[f]
}

func juice(f Fruits) Fruits {

	switch f {

	case Apple:

		return Apple

	case Mango:

		return Mango

	case Banana:

		return Banana

	case Papaya:

		return Papaya

	default:

		panic(fmt.Errorf("no type found"))
	}
}

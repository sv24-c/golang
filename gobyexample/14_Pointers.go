package main

import "fmt"

func main() {

	i := 11

	fmt.Println(i)

	pointer(i)

	fmt.Println(i)

	pointerRef(&i)

	fmt.Println(i)
	fmt.Println(&i)
}

func pointerRef(i *int) {
	*i = 12
}

func pointer(i int) {
	i = 12
}

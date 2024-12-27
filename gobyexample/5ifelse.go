package main

import "fmt"

func main() {
	i := 11

	if i > 1 && i == 11 {
		fmt.Println("tue")
	} else {
		fmt.Println("false")
	}

	if i > 1 {
		fmt.Println("greater")
	} else {
		fmt.Println("lower")
	}

	if i := 2; i > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	if j := 7; j%2 == 0 {
		fmt.Println("non prime")
	} else {
		fmt.Println("prime")
	}

}

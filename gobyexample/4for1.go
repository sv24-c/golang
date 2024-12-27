package main

import "fmt"

func main() {

	i := 0
	for i < 2 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 10; i > 0; i-- {
		fmt.Print(i, ' ')
	}

	//fmt.Println(" smit", 1111, ' ', "smit")

	fmt.Println()

	num := []int{12, 23, 3}

	for value := range num {

		if value%2 == 0 {
			fmt.Println("non prime")
		} else {
			fmt.Println("prime")
		}
	}

	for {
		fmt.Println("loop1")
		break
	}

}

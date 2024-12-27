package main

import "fmt"

func main() {

	fmt.Println(funcOne(1, 4, 2))

	_, min := funcOne(1, 4, 2)
	fmt.Println(min)

	sum, _ := funcOne(1, 4, 2)
	fmt.Println(sum)

	nums := []int{1, 2, 3, 4, 5}
	funcTwo(nums...)
}

func funcOne(a, b, c int) (int, int) {
	return a + b, b - c
}

func funcTwo(a ...int) {

	for _, num := range a {

		fmt.Println("num : ", num)
	}
	fmt.Println(a)
}

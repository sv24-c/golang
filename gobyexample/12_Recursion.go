package main

import "fmt"

func main() {

	fmt.Println(facto(6))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

func facto(i int) int {

	if i == 0 {
		return 1
	}
	return i * facto(i-1)
}

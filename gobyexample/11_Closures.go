package main

import "fmt"

func main() {

	funcOfFunc := closure()

	fmt.Println(funcOfFunc())
	fmt.Println(funcOfFunc())
	fmt.Println(funcOfFunc())

	funcOfFuncNew := closure()

	fmt.Println(funcOfFuncNew())
}

func closure() func() int {

	i := 0

	return func() int {
		i++
		return i
	}
}

package main

import "fmt"

func main() {

	var arr [5]int

	fmt.Println(arr)

	arr[4] = 1

	fmt.Println(arr)

	fmt.Println(len(arr))

	a2 := [4]int{1, 2, 3, 4}

	fmt.Println(a2)

	a3 := [...]int{1, 23, 5, 4}

	fmt.Println(a3)

	fmt.Println(len(a3))

	a4 := [...]int{1, 2, 3: 10, 8}

	fmt.Println(a4)

	var twoDArr [2][3]int

	fmt.Println(twoDArr)

	twoDArr = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(twoDArr)
}

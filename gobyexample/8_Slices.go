package main

import "fmt"

func main() {

	var s []string

	fmt.Println(s)

	s = make([]string, 3)

	fmt.Println(s)

	s[0] = "1"

	s[1] = "2"

	fmt.Println(s, " len :", len(s), "cap :", cap(s))

	s = append(s, "three")

	fmt.Println(s)
}

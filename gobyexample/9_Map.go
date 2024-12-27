package main

import "fmt"

func main() {

	map1 := make(map[string]int)

	fmt.Println(map1)

	map1["one"] = 1

	map1["two"] = 2

	fmt.Println(map1)

	fmt.Println(map1["one"], len(map1))

	delete(map1, "one")

	fmt.Println(map1)

	key, value := map1["one"]

	fmt.Println(key, value)

	map2 := map[string]int{"two": 2, "one": 1}

	fmt.Println(map2)
}

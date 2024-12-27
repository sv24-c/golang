package main

import (
	"fmt"
	"time"
)

func main() {

	f1("f1 called")

	go f1("f1 called with go")

	go func(string2 string) {
		fmt.Println(string2)
	}("going")

	time.Sleep(time.Second)

	fmt.Println("finish")
}

func f1(s string) {

	for i := 0; i < 3; i++ {

		fmt.Println()
		fmt.Println(s)
	}
}

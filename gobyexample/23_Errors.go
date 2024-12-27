package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {

	if arg == 32 {
		return -1, errors.New("should not be 32")
	}

	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("error ", e)
		} else {
			fmt.Println("result", r)
		}
	}

	fmt.Println(fmt.Errorf("test"))

}

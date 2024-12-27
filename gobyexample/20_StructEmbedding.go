package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func main() {

	con := container{base: base{
		num: 1,
	}, str: "str",
	}

	fmt.Printf("co={num: %v, str: %v}\n", con.num, con.str)

	fmt.Println("also num:", con.base.num)

	fmt.Println("describe:", con.describe())

	type describer interface {
		describe() string
	}

	var d describer = con
	fmt.Println("describer:", d.describe())
}

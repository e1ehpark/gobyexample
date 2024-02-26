package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %d", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v}\n", co.num, co.str)

	fmt.Println("also numm", co.base, co.num)

	fmt.Println("described", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("described", d.describe())
}

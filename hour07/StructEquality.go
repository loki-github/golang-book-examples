package main

import "fmt"

type Drink1 struct {
	name string
	ice  bool
}

func main() {
	drink1 := Drink1{
		name: "whiskey",
		ice:  true,
	}
	drink2 := &drink1

	drink2.ice = false

	fmt.Printf("%+v\n", *drink2)
	fmt.Printf("%+V\n", drink1)

	fmt.Printf("%p\n", drink2)
	fmt.Printf("%p\n", &drink1)
}

package main

import "fmt"

func main() {

	var cheeses = make([]string, 4)
	cheeses[0] = "cheese1"
	cheeses[1] = "cheese2"
	cheeses[2] = "cheese3"
	cheeses[3] = "cheese4"
	var cheesesCopy = make([]string, 2)
	copy(cheesesCopy, cheeses[2:4])
	fmt.Println("original cheeses=", cheeses)
	fmt.Println("copy of cheeses=", cheesesCopy)

}

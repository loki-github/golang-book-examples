package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "cheese1"
	cheeses[1] = "cheese2"
	var copyOfcheeses = make([]string, 1)
	copy(copyOfcheeses, cheeses)
	fmt.Println(copyOfcheeses)

}

package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "cheese1"
	cheeses[1] = "cheese2"
	cheeses = append(cheeses, "cheese3")
	cheeses = append(cheeses, "cheese4")
	cheeses = append(cheeses, "cheese5", "cheeses6", "cheeses7", "cheeses8", "cheeses9")

	cheeses = append(cheeses[:2], cheeses[3:]...)

	for i := 0; i < len(cheeses); i++ {
		fmt.Println(cheeses[i])
	}

}

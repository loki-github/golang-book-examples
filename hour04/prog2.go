package main

import "fmt"

func getPrizes() (int, string) {
	i := 2
	s := "Gold fish"
	return i, s
}

func main() {
	quantity, prize := getPrizes()
	fmt.Printf("You won %v %v\n", quantity, prize)
}

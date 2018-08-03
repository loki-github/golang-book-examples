package main

import "fmt"

func main() {
	numbers := []int{3, 2, 1, 5}
	for i, n := range numbers {
		fmt.Println(i)
		fmt.Println(n)
	}
}

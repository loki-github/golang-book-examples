package main

import "fmt"

func sumN_Numbers(numbers ...int) int {
	total := 0
	for n := range numbers {
		total = total + n
	}
	return total
}

func main() {
	result := sumN_Numbers(2, 2, 2, 2, 2)
	fmt.Printf("Sum of numbers is %v\n", result)
}

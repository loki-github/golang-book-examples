package main

import "fmt"

func main() {
	fmt.Println(isEven2(7))
}

func isEven2(x int) bool {
	return x%2 == 0
}

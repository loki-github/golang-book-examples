package main

import "fmt"

func main() {
	a, b, s := test(45, 67)
	fmt.Printf("returned values from function test are %v %v and %v\n", a, b, s)
}

func test(x, y int) (a, b int, s string) {
	fmt.Printf("passed arguments are %v n %v\n", x, y)
	a = 100
	b = 200
	s = "returning string"
	return a, b, s
}

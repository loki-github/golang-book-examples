package main

import "fmt"

func sayHi2() (x, y string) {
	x = "Hello"
	y = "World"
	return
}

func main() {
	fmt.Println(sayHi2())
}

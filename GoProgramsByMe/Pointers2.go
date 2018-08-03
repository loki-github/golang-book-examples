package main

import "fmt"

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(&i)

}

func showMemoryAddress(x *int) {
	fmt.Println(*x)
}

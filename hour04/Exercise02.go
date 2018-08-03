package main

import "fmt"

func main() {
	call10Times(1)
}

func call10Times(n int) {
	fmt.Printf("Calling %v st/th time\n", n)
	if n == 10 {
		return
	}
	call10Times(n + 1)
}

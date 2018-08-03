package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(feedMe2(1, 0))
}

func feedMe2(portion, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I am full, I have eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I am still hungry, I have eaten %v\n", eaten)
	return feedMe2(portion, eaten)
}

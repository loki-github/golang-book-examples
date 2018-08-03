package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strNumber = "1234"
	n, err := strconv.ParseInt(strNumber, 10, 64)
	fmt.Println(err)
	fmt.Println(n)
}

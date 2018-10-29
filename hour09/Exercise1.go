package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Oh I do like to be beside the seaside"
	fmt.Println(strings.ToUpper(s))

	fmt.Println(strings.Replace("seaside", "seaside", "bar", 2))

	fmt.Println(strings.Index("what the fuck", "the"))
}

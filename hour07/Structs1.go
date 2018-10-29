package main

import "fmt"

type Movie1 struct {
	name   string
	rating float32
}

func main() {
	m := new(Movie1)
	m.name = "some movie"
	fmt.Printf("Movie %+v", m)
}

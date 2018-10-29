package main

import "fmt"

type SuperHero struct {
	name    string
	age     string
	address Address2
}

type Address2 struct {
	number int
	street string
	city   string
}

func main() {
	superHero1 := SuperHero{
		name: "Loki",
		age:  "1000 years",
		address: Address2{
			number: 23,
			street: "hawthorn ct",
			city:   "brentwood",
		},
	}
	fmt.Printf("%+v", superHero1)
}

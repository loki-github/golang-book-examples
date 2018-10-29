package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	DOB       string
	ssn       string
	address   Address1
}

type Address1 struct {
	addressline1 string
	city         string
	state        string
	zip          int
}

func main() {
	person1 := Person{
		FirstName: "loki",
		LastName:  "kempanna",
		DOB:       "01/19/1986",
		ssn:       "12345645",
		address: Address1{
			addressline1: "2839 hawthorn ct",
			city:         "Brentwood",
			state:        "CA",
			zip:          94513,
		},
	}
	fmt.Printf("%+v", person1)
}

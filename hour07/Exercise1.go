package main

import "fmt"

type ConsoleGame struct {
	name       string
	releseDate string
	platform   string
}

func main() {
	mario := ConsoleGame{
		name:       "Mario",
		releseDate: "01/01/1996",
		platform:   "windows",
	}
	fmt.Printf("%+v", mario)
}

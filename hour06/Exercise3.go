package main

import "fmt"

func main() {
	var htmlMap = make(map[string]string)
	htmlMap["p"] = "Paragraph"
	htmlMap["img"] = "Image"
	htmlMap["h1"] = "Heading One"
	htmlMap["h2"] = "Heading Two"
	fmt.Println(htmlMap)
}

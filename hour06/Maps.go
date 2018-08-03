package main

import "fmt"

func main() {

	var myMap = make(map[string]int, 2)
	myMap["name1"] = 1
	myMap["name2"] = 2
	myMap["name3"] = 3
	myMap["name4"] = 4

	for entry := range myMap {
		fmt.Printf("key=%v and value=%v\n", entry, myMap[entry])
	}

	delete(myMap, "name1")
	fmt.Println(myMap)

}

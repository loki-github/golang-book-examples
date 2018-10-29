package main

import (
	"fmt"
	"time"
)

func slowFunc1() {
	time.Sleep(time.Second * 2)
	fmt.Println("Sleeper() finished")
}

func main() {
	go slowFunc1()
	fmt.Println("I will be called even when the slowFunc is called and I will wait till slowFunc finsihes")
	time.Sleep(time.Second * 3)
}

package main

import (
	"fmt"
	"time"
)

func slowFunc2(ch chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func main() {
	messages := make(chan string, 2)
	messages <- "hello loki"
	messages <- "hello sneha"
	close(messages)
	fmt.Println("Pushed two messages onto channel with no receivers")
	time.Sleep(time.Second * 2)
	slowFunc2(messages)

}

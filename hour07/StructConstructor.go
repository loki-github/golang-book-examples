package main

import "fmt"

type Alarm1 struct {
	time  string
	sound string
}

func NewAlarm1(time string) Alarm1 {
	a := Alarm1{
		time:  time,
		sound: "Kalxonn",
	}
	return a
}

func main() {
	fmt.Printf("%+v\n", NewAlarm1("07:00:01"))

}

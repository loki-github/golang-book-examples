package main

import (
	"errors"
	"fmt"
)

type MyRobot interface {
	MyPowerOn() error
}

type MyT850 struct {
	Name string
}

func (r *MyT850) MyPowerOn() error {
	return nil
}

type RD2 struct {
	Broken bool
}

func (r *RD2) MyPowerOn() error {
	if r.Broken {
		return errors.New("RD2 is broken")
	} else {
		return nil
	}
}

func MyBoot(r MyRobot) error {
	return r.MyPowerOn()
}

func main() {
	t := MyT850{
		Name: "t850robot",
	}

	r := RD2{
		Broken: false,
	}

	err1 := MyBoot(&r)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Robot is powered on")
	}

	err2 := MyBoot(&t)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Robot is powered on")
	}

}

package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() string {
	return fmt.Sprintln("Happy go lucky 123")
}

func (sa secretAgent) speak() string {
	return fmt.Sprintln("Shocking. Positively shocking!")
}

type people interface {
	speak() string
}

func info(p people) {
	fmt.Println(p.speak())
}

func main() {
	p1 := person{"Nina", "Roby"}
	sa1 := secretAgent{person{"James", "Bond"}, true}
	fmt.Println(p1)
	fmt.Println(sa1)
	p1.speak()
	sa1.speak()
}

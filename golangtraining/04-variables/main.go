package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.78
	d := true

	var f int
	var g float64
	var h string
	var i bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", h)
	fmt.Printf("%T \n", i)
}

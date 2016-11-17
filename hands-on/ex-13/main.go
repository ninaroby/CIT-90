//Adding onto this code: Can you assign “g1” to “x”? Why or why not?

package main

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 56
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = g1
	fmt.Println(x)
	// you CANNOT assign g1 to x. g1 is of type gator - they are two different types!
}

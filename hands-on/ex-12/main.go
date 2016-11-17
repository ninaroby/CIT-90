//Adding onto this code: Using var, declare an identifier “x” as type int (var x int). Print out “x”. Print the type of “x” using fmt.Printf(“%T\n”, x)

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
}

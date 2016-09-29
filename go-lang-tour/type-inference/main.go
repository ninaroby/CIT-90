package main

import "fmt"

func main() {
	v := -128 + 0.5i //Change me!
	fmt.Printf("v is of type %T\n", v)
}

/*When declaring a variable without specifying a specific type
(either by using the := syntax or var = expression syntax), the
variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

*/

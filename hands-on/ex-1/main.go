//Initialize a SLICE of int using a composite literal; print out the slice; range over the slice printing out just the index; range over the slice printing out both the index and the value

package main

import (
	"fmt"
)

func main() {
	s := []int{9, 18, 27}
	fmt.Println(s)

	for i := range s {
		fmt.Println(i)
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}

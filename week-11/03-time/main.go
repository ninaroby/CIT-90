package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	h := t.Hour()
	fmt.Println(h)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	s := greeting()
	fmt.Println(s)
}

func greeting() string {
	h := time.Now().Hour()
	fmt.Println(h)
	if h < 12 {
		return "Good morning"
	} else if h < 20 {
		return "Good evening"
	} else {
		return "Good night"
	}
}

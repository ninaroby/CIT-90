package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// h := time.Now().Hour()
	h := 25

	s, err := greeting(h)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(s)
}

func greeting(h int) (string, error) {
	if h < 12 {
		return "Good morning", nil
	} else if h < 20 {
		return "Good evening", nil
	} else if h < 24 {
		return "Good night", nil
	} else {
		return "", errors.New("24 hours in a day, bra")

	}
}

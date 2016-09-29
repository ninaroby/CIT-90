package main

import "fmt"

type animal struct {
	name    string
	color   string
	numLegs int
	numEars int
}

type jungleAnimal struct {
	animal
	earSize string
	trunk   bool
}

func main() {
	animal1 := jungleAnimal{
		animal{
			"Nelly",
			"grey",
			4,
			2,
		},
		"large",
		false,
	}
	fmt.Println(animal1)
	isEl(animal1)
}

func isEl(ja jungleAnimal) {
	if ja.trunk == true {
		fmt.Println("I am an elephant")
	}
}

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
		animal{"Nelly", "grey", 4, 2},
		"large",
		true,
	}
	fmt.Println(animal1)
}

func elephant() {
	if trunk == true {
		fmt.println("I am an elephant")
	}
}

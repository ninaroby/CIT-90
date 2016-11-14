//Give a method to both the “truck” and “sedan” types with the following signature
//transportationDevice() string
//Have each func return a string saying what they do. Create a value of type truck and populate the fields. Create a value of type sedan and populate the fields. Call the method for each value.

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (y truck) transportationDevice() string {
	return fmt.Sprintln("I am a truck with", y.doors, "doors and I am durable and cool!")
}

func (z sedan) transportationDevice() string {
	return fmt.Sprintln("I am a sedan with", z.doors, "doors and I am luxurious but a little bit boring.")
}

func main() {
	t1 := truck{
		vehicle{
			2,
			"white",
		},
		true,
	}
	fmt.Println(t1.transportationDevice())

	s1 := sedan{
		vehicle{
			4,
			"black",
		},
		false,
	}
	fmt.Println(s1.transportationDevice())
}

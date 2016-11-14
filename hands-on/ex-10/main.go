//Create a new type called “transportation”. The underlying type of this new type is interface. An interface defines functionality. Said another way, an interface defines behavior. For this interface, any other type that has a method with this signature …
//transportationDevice() string
//… will automatically, implicitly implement this interface. Create a func called “report” that takes a value of type “transportation” as an argument. The func should print the string returned by “transportationDevice()” being called for any type that implements the “transportation” interface. Call “report” passing in a value of type truck. Call “report” passing in a value of type sedan.

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
	return fmt.Sprintln("I am a sedan with", z.doors, "doors and I am luxurious, but dull.")
}

type transportation interface {
	transportationDevice() string
}

func report(g transportation) {
	fmt.Println(g.transportationDevice())
}

func main() {
	t1 := truck{
		vehicle{
			2,
			"white",
		},
		true,
	}
	report(t1)

	s1 := sedan{
		vehicle{
			4,
			"black",
		},
		false,
	}
	report(s1)
}

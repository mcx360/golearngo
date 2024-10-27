package main

import "fmt"

//1. How do you access the fourth element of an array or slice?

func main() {

	//access 4th element of an array of names
	names := [5]string{"Albert", "Jacqueline", "Jamal", "Tyrone", "Abdul"}
	fmt.Println(names[3])

	//access 4th element of a slice of ages that will resize based on the size of the array of names, so that each name index
	//has a corresponding age index
	ages := [len(names)]int{
		32,
		58,
		68,
		21,
		19,
	}
	fmt.Println(ages[3])

}

package main

import "fmt"

// N.B to run type the command go run *.go to the terminal as otherwise an error will occur,
// since just running main.go does'nt compile the other files in the directory
func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)

	//modify its fields and print
	c.x = 10
	c.y = 10
	fmt.Println(c.x, c.y, c.r)

	//pass the address of Circle struct to circleArea function and print the result
	fmt.Println(circleArea(&c))

	//show how methods work
	fmt.Println(c.area())

	//show how methods work again, but with rectangle this time
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.rectangleArea())
}

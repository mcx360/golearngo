package main

//Circle struct is defined in fields.go

import "math"

/* In between the keyword func and the name of the function, we’ve added a receiver.
   The receiver is like a parameter—it has a name and a type—but by creating the func‐
   tion in this way, it allows us to call the function using the . operator */

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

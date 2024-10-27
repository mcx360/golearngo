package main

import "fmt"

//new takes a type as an argument, allocates enough memory to fit a value of that type,
//and returns a pointer to it.

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}

package main

//How do you create a new pointer?

import "fmt"

func main() {
	xPtr := new(int)
	*xPtr = 12
	fmt.Println(*xPtr)
}

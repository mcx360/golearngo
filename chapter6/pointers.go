package main

// * = pointer
// & = address of a variable

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) //x is 0
}

package main

import "fmt"

//How do you assign a value to a pointer?

func main() {
	var x int = 10
	var ptr *int = &x
	fmt.Println(*ptr)
}

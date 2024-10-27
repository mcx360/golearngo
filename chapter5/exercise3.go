package main

import "fmt"

func main() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	y := x[2:5] //copy array x starting at index 2 up to but not including index 5
	fmt.Println(x)
	fmt.Println(y)
}

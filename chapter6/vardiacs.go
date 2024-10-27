package main

import "fmt"

//vardiac functions are functions that allow an unspecified amout of parameters to a function

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	fmt.Println(add(1, 2, 3))

	//we can pass a slice of ints to the add function using an ellipsis
	xs := []int{3, 3, 3}
	fmt.Println(add(xs...))
}

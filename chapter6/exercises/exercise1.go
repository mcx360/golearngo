package main

//sum is a function that takes a slice of numbers and adds them together. What
//would its function signature look like in Go?

import "fmt"

func sum(args ...int) int {
	total := 0

	for _, i := range args {
		total += i
		i++
	}
	return total
}

func main() {

	nums := []int{5, 5, 5, 5, 5}
	fmt.Println(sum(nums...))
}

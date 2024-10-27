package main

import "fmt"

//Write a function with one variadic parameter that finds the greatest number in a
//list of numbers.

func findGreatest(args ...int) (greatest int) {
	greatest = 0

	for _, i := range args {
		if i > greatest {
			greatest = i
		}
	}
	return greatest
}

func main() {
	nums := []int{8, 9, 10, 6, 7}
	fmt.Println(findGreatest(nums...))
}

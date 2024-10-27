package main

//Write a function that takes an integer and halves it and returns true if it was even
//or false if it was odd. For example, half(1) should return (0, false) and
//half(2) should return (1, true).

import "fmt"

func half(num int) (n int, even bool) {
	if num%2 == 0 {
		even = true
	}

	n = num / 2

	return n, even
}

func main() {
	fmt.Println(half(4))
}

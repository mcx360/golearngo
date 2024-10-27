package main

//What are defer, panic, and recover? How do you recover from a runtime panic?
//I will use dividing by zero as an example of panic

import "fmt"

func divide(a float64, b float64) float64 {
	if a == 0 || b == 0 {
		panic("Cannot divide by 0!")
	}
	return a / b
}

func main() {
	fmt.Println(divide(15, 0))
}

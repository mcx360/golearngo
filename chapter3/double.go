package main

//testing user input in golang

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

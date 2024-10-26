package main

//testing user input in golang

import "fmt"

//program that converts feet to meters

func main() {
	fmt.Println("Enter the feet value:")
	var feet float32
	fmt.Scanf("%f", &feet)
	var meters float32 = feet / 3.281
	fmt.Println(meters)
}

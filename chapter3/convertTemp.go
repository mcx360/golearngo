package main

//testing user input in golang

import "fmt"

func main() {
	fmt.Printf("Enter fahrenheit temperature:")
	var fahrenheitTemp float32
	fmt.Scanf("%f", &fahrenheitTemp)
	var celciusTemp float32 = (fahrenheitTemp - 32) * 5 / 9
	fmt.Println(celciusTemp)
}

package main

//defer functions in golang

import "fmt"

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func main() {
	defer second()
	first()
}

package main

import "fmt"

//A program that given an int between 1 and 5 will output that number english - used to test switch statements in golang

func main() {
	fmt.Printf("Enter a number between 1 and 5(inclusive)")
	var num int
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("invalid!")
	}
}

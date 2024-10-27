package main

//2. What is the length of a slice created using make([]int, 3, 9)?

import "fmt"

func main() {
	x := make([]int, 3, 9) // create a slice of length 3 with an array of length 9
	fmt.Println(len(x))    // proof of the above since it returns the length of the slice which is 3
}

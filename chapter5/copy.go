package main

//use copy function for slices in golang

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1) //copy contents of slice1 into slice2
	fmt.Println(slice1, slice2)

}

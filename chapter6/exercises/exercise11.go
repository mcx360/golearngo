package main

//Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y)
//should give you x=2 and y=1).

import "fmt"

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	x := 2
	y := 3

	fmt.Println("Before swap:")
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println("After swap:")
	fmt.Println(x, y)

}

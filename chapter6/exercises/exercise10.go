package main

import "fmt"

//What is the value of x after running this program

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 1.5
	square(&x)
	fmt.Println(x) // x = 2.25
}

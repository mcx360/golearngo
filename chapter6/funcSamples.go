package main

import "fmt"

// paramaters have to be passed implicitly to function so that it can access it
func f(x int) {
	fmt.Println(x)
}

// return types can have names
func f1() (r int) {
	r = 1
	return
}

// return multiple values from a function
func f2() (int, int) {
	return 5, 6
}

func main() {

	//test function f()
	x := 5
	f(x)

	//test function f1()
	fmt.Println(f1())

	//test function f2()
	fmt.Println(f2())

}

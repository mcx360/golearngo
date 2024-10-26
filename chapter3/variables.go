package main

import "fmt"

//testing variable in golang

func main() {
	var x string = "hello, world!!"
	fmt.Println(x)
	x = "bye, world!"
	fmt.Println(x)

	var y string = "Michal"
	var z string = "Michal"
	fmt.Println(y == z)

	p := "learn go!"
	fmt.Println(p)

	q := 666
	fmt.Println(q)

	const myName string = "Michal"
	fmt.Println(myName)

	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println(a + b + c)

}

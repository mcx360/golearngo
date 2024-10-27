package main

//local functions in golang

import "fmt"

// makeEvenGenerator returns a function that generates even numbers. Each time it’s
// called, it adds 2 to the local i variable, which—unlike normal local variables—persists
// between calls.
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	//A function like this together with the nonlocal variables it references is known as a
	//closure.
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println()

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}

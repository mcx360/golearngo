package main

import "fmt"

//testing for loop in go

func main() {
	i := 1
	for i <= 10 {
		fmt.Print(i)
		if i%2 == 0 {
			fmt.Print(" Even\n")
		} else {
			fmt.Print(" Odd\n")
		}
		i = i + 1
	}

}

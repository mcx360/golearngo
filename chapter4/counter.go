package main

import "fmt"

//testing for loop in go

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for y := 11; y <= 20; y++ {
		fmt.Println(y)
	}
}

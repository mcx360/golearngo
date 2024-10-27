package main

import "fmt"

func main() {
	mySlice := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var smallest int = mySlice[0]

	var i int = 0
	for i <= len(mySlice)-1 {
		if mySlice[i] < smallest {
			smallest = mySlice[i]
		}
		i++
	}

	fmt.Println(smallest)
}

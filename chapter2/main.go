package main

import "fmt"

func main() {
	fmt.Println(len("hello world"))
	fmt.Println("hello world"[1])
	fmt.Println("hello " + "world")

	fmt.Println("\n-------------")

	fmt.Println(" AND Truth table")
	fmt.Println("| True AND True  |", true && true)
	fmt.Println("| True AND False |", true && false)
	fmt.Println("| False AND True |", false && true)
	fmt.Println("| False AND True |", false && false)

	fmt.Println("\n-------------")

	fmt.Println(" Not truth table")
	fmt.Println("|False|", !false)
	fmt.Println("|True |", !true)

	fmt.Println(32132 * 42452)
}

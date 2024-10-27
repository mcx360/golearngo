package main

import "fmt"

// pairing panic with defer is essential as panic immediately stops the function from executing
func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}

package main

import "fmt"

//The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) =
//fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).

func fib(n int) (ret int) {
	if n == 0 {
		ret = 0
	} else if n == 1 {
		ret = 1
	} else {
		ret = fib(n-1) + fib(n-2)
	}
	return ret
}

func main() {
	fmt.Println(fib(9))
}

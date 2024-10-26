package main

//learn basics of slices in golang

import "fmt"

func main() {
	var a []float64             //slice
	b := make([]float64, 5, 10) //slice of size 5 with an array of size 10
	fmt.Println(a)
	fmt.Println(b)

	//[low : high] expresion

	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	y := arr[1:4]
	z := arr[0:]
	p := arr[:5]
	q := arr[:]

	fmt.Println(arr)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(p)
	fmt.Println(q)

}

package main

import "fmt"

//Using makeEvenGenerator as an example, write a makeOddGenerator function
//that generates odd numbers.

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
		return
	}

}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

}

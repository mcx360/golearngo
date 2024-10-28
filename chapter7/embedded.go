package main

import "fmt"

type Person struct {
	Name string
}

// Android has "is a" relationship with Person, this is known as an embedded type in go
type Android struct {
	Person
	model string
}

func (p *Person) Talk() {
	fmt.Println("Hello, my name is ", p.Name)
}

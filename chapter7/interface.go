package main

type Shape interface {
	area() float64
}

// Because both of our shapes have
// an area method, they both implement the Shape interface and we can change our
// function to this:
// All totalArea knows about each shape is that it has an area method:
// Nothing additional is required to implement an interface (there is no implements or
// extends keyword). It’s sufficient to merely have a method with the same name and
// signature.
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// Interfaces can also be used as fields. Consider a MultiShape that is made up of several
// smaller shapes:
type Multishape struct {
	shapes []Shape
}

// We can even turn MultiShape itself into a Shape by giving it an area method:
func (m *Multishape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

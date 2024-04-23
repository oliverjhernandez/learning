package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func main() {
	myTriangle := Triangle{
		height: 3.21,
		base:   5.78,
	}

	printArea(myTriangle)
}

func (t Triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	result := s.getArea()
	fmt.Println(result)
}

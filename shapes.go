package main

import "math"

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	var area float64 = math.Pi * (c.radius * c.radius)
	return area
}

func (r Rectangle) Area() float64 {
	var area float64 = r.height * r.width
	return area
}

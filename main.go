package main

import "fmt"

func main() {
	var rect = Rectangle{Width: 10, Height: 20}
	var area float64 = calculateArea(rect)
	fmt.Println("Area is", area)

}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func (r Rectangle) Area() float64 {
	var area float64 = r.Width * r.Height
	return area
}

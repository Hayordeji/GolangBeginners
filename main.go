package main

import (
	"fmt"
	"io"
)

func main() {
	var newCircle Circle = Circle{
		radius: 6,
	}
	var newRectangle Rectangle = Rectangle{
		height: 12,
		width:  15,
	}

	var CircleArea float64 = newCircle.Area()
	var RectangleArea float64 = newRectangle.Area()

	fmt.Println("Area of circle is: ", CircleArea)
	fmt.Println("Area of rectangle is: ", RectangleArea)
	var buff io.Reader
	fmt.Fscan(buff)
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

// func (r Rectangle) Area() float64 {
// 	var area float64 = r.Width * r.Height
// 	return area
// }

package main

import (
	"fmt"
	"math"
)

type Retangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Retangle) area() float64 { //定义Retangle的method
	return r.width * r.height
}

func (c Circle) area() float64 { //定义Circle的method
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Retangle{width: 8, height: 9}
	c := Circle{radius: 5}
	fmt.Println("Area of the retangle:", r.area()) //调用method
	fmt.Println("Area of the circle:", c.area())
}

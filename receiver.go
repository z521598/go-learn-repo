package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}
type Areable interface{
	area() float64
}

func (r Rectangle) area() float64{
	return r.width * r.height
}

func (c Circle) area() float64{
	return c.radius * c.radius * math.Pi
}

func main(){
	ra := Rectangle{5,5}
	ca := Circle{2.5}
	fmt.Println(ra.area())
	fmt.Println(ca.area())
}


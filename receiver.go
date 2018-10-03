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

func (r Rectangle) area() float64{
	return r.width * r.height
}

func (c Circle) area() float64{
	c.radius = 100
	return c.radius * c.radius * math.Pi
}

func main(){
	var c interface{}
	r := Rectangle{10,10}
	c1:= Circle{5}
	c = c1
	fmt.Println(r.area())
//	fmt.Println(c.area())
	fmt.Println(c)
	//value,ok = Circle.(c)
	//if ok {
	//	fmt.Println(value)
	//}
	// 只有intercace{} 才能使用comma-ok语法
	fmt.Println(c1.(Circle));
	fmt.Println(c.(Circle))
}


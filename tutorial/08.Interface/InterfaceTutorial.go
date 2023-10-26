package main

import (
	"math"
)

//interface 정의
type Shape interface{
	area() float64
	perimeter() float64
}


//structure Rectangle 정의
type Rect struct{
	width, height float64
}

//structure Circle 정의
type Circle struct{
	radius float64
}


//interface 구현
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64{
	return 2* (r.width + r.height)
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func main(){
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}


func showArea(shapes ...Shape){
	for _, s := range shapes{
		a := s.area()
		println(a)
	}
}









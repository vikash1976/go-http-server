package main

import (
	"math"
	"fmt"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (z Square) area() float64 {
	return z.side * z.side
}
func (z *Square) setValue(v float64) {
	z.side = v
}

func (z Circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

type Shape interface {
	area() float64
	//setValue(v float64)
}

func info(z Shape){
	fmt.Printf("%T\n", z)
	fmt.Println(z.area())
}
func main() {
	s := Square{side: 5}

	c := Circle{radius: 5}

	info(s)
	s.setValue(9)
	info(s)
	info(c)


}


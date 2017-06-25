package main

import (
	"fmt"
	"math"
)

type Accumulator interface {
	accumulate(i interface{}) float64
}
type DataPresenter interface {
	show(i interface{})
}
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width

}

type ShapeAreaAccumulator struct {
	areaSum float64
}

func (s ShapeAreaAccumulator) accumulate(shapes ...Shape) float64 {
	for _, shape := range shapes {
		s.areaSum += shape.area()
	}
	return s.areaSum
}

type DataValuePresenter struct{}

func (d DataValuePresenter) show(data interface{}) {
	fmt.Println(data)
}

type DataValue_DescPresenter struct{}

func (d DataValue_DescPresenter) show(data interface{}) {
	fmt.Printf("The accumulated value is: %v\n", data)
}

type Square struct {
	Rectangle
}

func main() {
	c := Circle{5}
	r := Rectangle{5, 10}
	s := Square{Rectangle{4, 4}}

	sAC := ShapeAreaAccumulator{}
	sAC.areaSum = sAC.accumulate(c, r, s)
	valuePresenter := DataValuePresenter{}
	valuePresenter.show(sAC.areaSum)

	valueDescPresenter := DataValue_DescPresenter{}
	valueDescPresenter.show(sAC.areaSum)

}

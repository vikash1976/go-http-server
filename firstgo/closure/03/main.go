package main

import "fmt"

func wrapper(i float64) func() float64 {
	x := i
	return func() float64 {
		x = x + 1.0
		return x
	}

}

func calcInterest(principal int) func(int) func(int) func() int {

	return func(time int) func(int) func() int {
		return func(rate int) func() int {
			return func() int {
				return (principal * time * rate) / 100
			}

		}
	}

}

func main() {
	increment := wrapper(10)
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Printf("%f", increment())

	for100 := calcInterest(100)
	for100n5yrs := for100(5)
	for100n5yrsAt10 := for100n5yrs(10)
	//increment12 := increment1(2)
	fmt.Println("\n", for100n5yrsAt10())

	for100n5yrsAt5 := for100n5yrs(5)
	fmt.Println("\n", for100n5yrsAt5())

}

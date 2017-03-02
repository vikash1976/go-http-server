package main

import (
	"fmt"
	"../stringutils"

)

func main() {
	fmt.Println("Hello, World!!! GOOOOOO has arrived.")
	fmt.Printf("%d\t%b\t%x\n", 42, 42,42)
	fmt.Printf("%d\t%b\t%#x\n", 42, 42,42)
	fmt.Printf("%d\t%b\t%#X", 42, 42,42)

	for i :=20; i < 200; i++ {
		fmt.Printf("%d\t%b\t%x\t%q\n", i, i, i, i)
	}
	stringutils.Reverse(stringutils.MyName)
	a := 1
	b := "abc"
	c := 4.14
	d := true

	fmt.Printf("%v\t%v\t%v\t%v\t", a,b,c,d)

	fmt.Println()
	fmt.Printf("%T\t%T\t%T\t%T\t", a,b,c,d)

	var a1 int
	var b1 string
	var c1 float64
	var d1 bool
	fmt.Println()
	fmt.Printf("%v\t%v\t%v\t%v\t", a1,b1,c1,d1)

	var a11, b11 string = "a's value", "b's value"
	fmt.Println()
	fmt.Printf("%v\t%v\t", a11,b11)

	foo()

}

func foo() {
	a := 12
	fmt.Println(a)
	/*a := 13 order matters*/
}

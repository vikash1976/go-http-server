package main

import "fmt"

func main() {

	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	//var b = &a this also works as per value of &a, it infers its type as pointer to int
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	fmt.Println(*b)

	*b = 66
	fmt.Println(a)
}

package main

import "fmt"

func main(){
	student := []string {} // creates the underlying array but no spots available, use append and not index style []
	fmt.Println(student == nil)

	var student1 [] string  // creates a slice but no underlying array, hence address header portion is nil, use append
	fmt.Println(student1 == nil)

	student2 := make([]string, 20, 40) // fully available for use, upto len can use [],
	// after that append, good to use append always
	fmt.Println(student2 == nil)
	student2[19] = "AAA"
	fmt.Println(student2[19], student2[16])

}

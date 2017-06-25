package main

import (
	"fmt"
)

func getAge() *int {
	age := 34
	return &age
}
func main() {
	fmt.Println(*getAge())
}

/*func HasIt() bool {
	return false
}

func main() {
	if HasIt() {
		fmt.Println("Have It")
	}
}*/

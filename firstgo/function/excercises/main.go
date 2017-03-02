package main

import "fmt"

/*func half1 (x int) (float64, bool) {
	return float64(x)/2, x%2 == 0
}*/

func half (x int) (int, bool) {
	return x/2, x%2 == 0
}

func main(){

	half1 := func (x int) (float64, bool) {
	return float64(x)/2, x%2 == 0
	}

	fmt.Println(half(5))
	fmt.Println(half(8))

	h, even := half(17)
	fmt.Println(h, even)

	h1, even := half1(17)
	fmt.Println(h1, even)
}

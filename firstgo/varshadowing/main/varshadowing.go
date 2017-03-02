package main

import "fmt"

func max (x int) int {
	return 42 + x
}
func main() {
	max := max(12)
	fmt.Println(max)

	//we have shadowed it, so now can't treat it as func any more
	// max(11)
}

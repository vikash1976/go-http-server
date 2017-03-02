package main

import (
	"fmt"
)
/*func () {
	fmt.Println("I am on my own!!!")
}()*/
func main() {
	func () {
		fmt.Println("I am on my own!!!")
	}()
}

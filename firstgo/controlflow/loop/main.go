package main

import (
	"fmt"
)

func main() {
	i := 0

	for {
		i++
		if(i % 2 == 0) {
			continue
		}
		fmt.Println(i)
		if i >=50 {
			break
		}
	}
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}

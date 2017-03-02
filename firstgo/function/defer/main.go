package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func emotion() {
	fmt.Println("!!!")
}

func main(){

	defer emotion() //looks like it stacks the defer listing and works as LIFO
	defer world()
	hello()
}
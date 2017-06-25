package main

import (
	"log"

	"github.com/vikash1976/go-tooling/explore-test/fibonaci"
)

func main() {

	log.Println("Fib(%d) is %d", 10, fibonaci.Fib(10))
	log.Println("Fib(%d) is %d", 10, fibonaci.Fib(20))
	log.Println("Fib(%d) is %d", 10, fibonaci.Fib(30))
	log.Println("Fib(%d) is %d", 10, fibonaci.Fib(40))
}

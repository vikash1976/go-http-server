/*
package main

import (
	"fmt"

)

func main() {
*/
/*c := incrementor()
cSum := puller(c)

for n := range cSum {
	fmt.Println(n)
}*/ /*

	c := make(chan float64)
	go func(){
		c <- 1
		close(c)
	}()
	fmt.Printf("%f\n",<-c)
	fmt.Printf("%f\n",<-c)
	fmt.Printf("%f\n",<-c)// after channel is closed,
	// any fetch will return zero value of the type channel is created for
	//if not closed any fetch will wait for a write and since no one is writing
	// it gets into deadlock

	c1 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)
	}()

	for n := range c1{
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func (){
		fmt.Println("Starting Incrementor")
		for i := 0; i <10; i++ {
			out <- i
		}
		close(out)
		fmt.Println("Incrementor done")
	}()
	return out
}

func puller(c chan int) chan int {
	sumOut := make(chan int)
	go func() {

		fmt.Println("Starting Puller")
		var sum int
		for n := range c {
			sum += n
		}
		sumOut <- sum
		close(sumOut)
		fmt.Println("Puller done")
	}()
	return sumOut
}
*/
/*package main

import (
	"fmt"
)

func main() {
	g := gen()
	c := factorial(g)
	for n := range c {
		fmt.Println(n)

	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Writing to Channel", i)
			out <- i
			
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			
			out <- fact(n)
			fmt.Println("Wrote to Channel calculated factorial")
		}

		
		close(out)
	}()
	fmt.Println("Returning")
	return out
}

func fact(n int) int {
	//fmt.Println("Factorial of: ", n)
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}*/

package main

import (
	"fmt"
)

func main() {
	g := gen()
	c := factorial(g)
	for n := range c {
		fmt.Println(n)

	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 100000; i++ {
			//fmt.Println("Writing to Channel", i)
			out <- i
			
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			
			out <- fact(n)
			//fmt.Println("Wrote to Channel calculated factorial")
		}

		
		close(out)
	}()
	fmt.Println("Returning")
	return out
}

func fact(n int) int {
	//fmt.Println("Factorial of: ", n)
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}


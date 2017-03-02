package main

import "fmt"

func visit(input []int, callback func(int)) {
	for _, v := range input {
		callback(v)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}

	}
	fmt.Printf("%T\n", callback)
	return xs
}

func factorial(x int) int {
	if(x == 0) {
		return 1;
	}
	return x * factorial(x - 1)
}


func main() {
	visit([]int{1, 2, 3}, func(n int) {
		calc := n * 2
		fmt.Println(calc)
	})
	visit([]int{10, 20, 30}, func(n int) {
		calc := n * 3
		fmt.Println(calc)
	})

	xs := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 3
	})

	fmt.Println(xs)

	fmt.Println(factorial(5))


}

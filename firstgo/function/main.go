package main

import "fmt"

func main(){
	data := []float64{23,33,44,55,66,77}
	avg := averageIt(data...) //variadic arguments
	fmt.Println(avg)

	avg = average(data)
	fmt.Println(avg)
}

func averageIt(sf ... float64) float64 { //variadic function
	var total float64

	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))

}

func average(sf []float64) float64 { //variadic function
	var total float64

	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))

}


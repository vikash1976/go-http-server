package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"

	"strconv"
)
type Price struct {
	Tick string
	Price int
}
func main() {
	res, _ := http.Get("http://localhost:1337/prices")

	//fmt.Printf("%s\n", res.Body)
	page, _ := ioutil.ReadAll(res.Body)

	var prices [] Price
	fmt.Printf("%s\n", page)
	fmt.Printf("%T\n", page)

	json.Unmarshal(page, &prices)
	fmt.Printf("%s\n", prices[0].Tick)

	var noType interface {}
	json.Unmarshal(page, &noType)
	fmt.Printf("%v\n", noType)

	m := noType.( []interface{})
	 fmt.Printf("%v\n", m)

	for i, u := range m {
		fmt.Println(i, u)
		fmt.Printf("Type of u is: %T\n", u)
		d := u.(map [string] interface {})

		fmt.Printf("Type of d is: %T\t%v\t%v\n", d, d["tick"], d["price"])
	}
	var newPrices [3] Price //:= Price{Tick: "A", Price: 300}

	for i:=0; i < 3; i++ {
		newPrices[i].Tick = "A" + strconv.Itoa(i)
		newPrices[i].Price = i + 500
	}
	newPricesAsJSON, _ := json.Marshal(newPrices)

	fmt.Println(string(newPricesAsJSON))

}

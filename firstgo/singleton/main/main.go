package main

import (
	"github.com/vikash1976/firstgo/singleton/store"
	"fmt"
)

func main(){
	var store store.Store
	store.Set("name", "vikash")
	store.Set("id", "101")

	fmt.Println(store.Get("id"))

	//var store1 store.Store
	store.Set("name", "abc")
	fmt.Println(store.Get("name"))
}

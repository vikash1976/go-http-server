package main

import "fmt"

func main(){
	var myMap = make(map [string]string)
	myMap["name"] = "AAAA"

	fmt.Println(myMap)

	var myMap1 = map [string]string {}
	myMap1["id"] = "101"
	fmt.Println(myMap1)
	
 elements := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
    "Be": map[string]string{
      "name":"Beryllium",
      "state":"solid",
    },
    "B":  map[string]string{
      "name":"Boron",
      "state":"solid",
    },
    "C":  map[string]string{
      "name":"Carbon",
      "state":"solid",
    },
    "N":  map[string]string{
      "name":"Nitrogen",
      "state":"gas",
    },
    "O":  map[string]string{
      "name":"Oxygen",
      "state":"gas",
    },
    "F":  map[string]string{
      "name":"Fluorine",
      "state":"gas",
    },
    "Ne":  map[string]string{
      "name":"Neon",
      "state":"gas",
    },
  }

  if el, ok := elements["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }

x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
smallest := x[0]
for _,n := range x {
	if n < smallest {
		smallest = n
	} 
}
fmt.Println("Smallest: ", smallest)
}

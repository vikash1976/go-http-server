package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"os"
)

type Person struct {
	First       string
	Last        string `json:"-"`        // don't include it
	Age         int    `json:"HowYoung"` //change the key to HowYoung
	notExported int
}

type Employee struct {
	First    string
	Last     string `json:"-"`        // don't include it
	Age      int    `json:"HowYoung"` //change the key to HowYoung
	Exported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 107}

	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Println(string(bs))
	var p2 Employee

	json.Unmarshal(bs, &p2)

	fmt.Println(p2)
	fmt.Println(p2.Age)

	var p3 Person
	rdr := strings.NewReader(`{"First": "Vikash", "Last": "Pan", "HowYoung": 39}`)
	json.NewDecoder(rdr).Decode(&p3)
	fmt.Println(p3)

	json.NewEncoder(os.Stdout).Encode(p1)
}

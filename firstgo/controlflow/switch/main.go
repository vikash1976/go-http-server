package main

import "fmt"

type Contact struct {
	Name string
	Zip  string
}

func switchCase(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("An Int")
	case string:
		fmt.Println("A String")
	case Contact:
		fmt.Println("A Contact")
	default:
		fmt.Println("Can't figure out")
	}
}
func main() {
	switchCase(4)
	switchCase("AAA")
	contact := Contact{Name: "Name1", Zip: "110011"}
	switchCase(contact)
	switchCase(77.89)
	age := 41
	name := "Name11"

	switch {
	case age <= 32, name == "Name <32" :
		fmt.Println("Yep")

	case age == 42, name == "Name11"  :
		fmt.Println("Yep!!!")

	default:
		fmt.Println("Nop!!!")


	}
}

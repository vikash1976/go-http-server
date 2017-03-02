package stringutils

import "fmt"

func Reverse (s string) {
	//call sort of private Func in the same package
	fmt.Printf("Exported One called... calls un-exported one: %q\n", s)
	var reversed = reverseTwo(s)
	fmt.Printf("%q\n", reversed)
}

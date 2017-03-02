package stringutils

import "fmt"

func reverseTwo (s string) string{
	fmt.Printf("Un-Exported called from exported one: %q", s)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

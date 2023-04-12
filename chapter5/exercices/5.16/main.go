// Write a variadic version of strings.Join.
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"1", "2", "3"}
	fmt.Println(strings.Join(s, "-"))
	fmt.Println(join("-", s...))

	s = []string{"I", "like", "to", "GO", "fast!"}
	fmt.Println(strings.Join(s, "   "))
	fmt.Println(join("   ", s...))
}

func join(separator string, items ...string) string {

	res := ""
	for idx, item := range items {
		res += item

		if idx != len(items)-1 {
			res += separator
		}
	}
	return res
}

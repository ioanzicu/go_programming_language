// Write a function expand(s string, f func(string) string) string that
// replaces each substring ‘‘$foo’’ within s by the text returned by f("foo").
package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\$\w+`)

func main() {
	s := "$this is very $interesting thing?"

	res := expand(s, strings.ToUpper)
	fmt.Println(res)
}

func expand(s string, f func(string) string) string {
	wrapper := func(msg string) string {
		msg = msg[1:]
		return f(msg)
	}
	return re.ReplaceAllStringFunc(s, wrapper)
}

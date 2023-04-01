// Write a non-recursive version of comma,
// using bytes.Buffer instead of string concatenation.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	// comma inserts commas in a non-negative decimal integer string
	fmt.Println(comma("1"))          // 1
	fmt.Println(comma("12"))         // 12
	fmt.Println(comma("123"))        // 123
	fmt.Println(comma("12345"))      // 12,345
	fmt.Println(comma("123456789"))  // 123,456,789
	fmt.Println(comma("1234567890")) // 1,234,567,890

}

// func comma(s string) string {
// 	n := len(s)
// 	if n <= 3 {
// 		return s
// 	}

// 	b := bytes.Buffer{}
// 	// get the first numbers up to 3 numbers form left side
// 	prefix := len(s) % 3
// 	b.WriteString(s[:prefix])

// 	for i := prefix; i < n; i += 3 {
// 		if i != 0 { // skip when prefix is 0, example 123 456
// 			b.WriteString(",")
// 		}
// 		b.WriteString(s[i : i+3])
// 	}
// 	return b.String()
// }

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	buf := bytes.Buffer{}
	// get the first numbers up to 3 numbers form left side
	prefix := len(s) % 3
	buf.WriteString(s[:prefix])

	for i := prefix; i < n; i += 3 {
		if i != 0 { // skip when prefix is 0, example 123 456
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

// Modify reverse to reverse the characters of a []byte slice that represents a
// UTF-8-encoded string, in place. Can you do it without allocating new memory?
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	t := []byte("Hello, 世界!")
	fmt.Printf("%+#v\n%s\n", t, string(t))
	reverseBytes(t)
	fmt.Printf("%+#v\n%s\n", t, string(t))

}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverse2(b []byte) {
	length := len(b)
	for i := 0; i < length/2; i++ {
		b[i], b[length-1-i] = b[length-1-i], b[i]
	}
}

func reverseBytes(b []byte) {
	fmt.Printf(" - %v\n", b)

	// reverse runes first
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}

	fmt.Printf(" + %v\n", b)

	// reverse the entire string
	reverse(b)
	fmt.Printf(" = %v\n", b)

}

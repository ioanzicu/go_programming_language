// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"
}

func basename(s string) string {
	// Discard last '/' and everything before
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]

	// Preserve everything before last '.'
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		s = s[:dot]
	}
	return s
}

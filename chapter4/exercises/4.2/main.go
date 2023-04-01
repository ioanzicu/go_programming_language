// Write a program that prints the SHA256 hash of its standard input by default but
// supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var version = flag.Uint("v", 256, "sha version, default 256")

/* first flag, then input
go run chapter4/exercises/4.2/main.go -v=512 hi
*/

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		panic("Please provide sha version")
	}
	fmt.Println(os.Args[1:], *version)
	input := os.Args[1:][1]

	var x interface{}

	switch {
	case *version == uint(512):
		x = sha512.Sum512([]byte(input))
	default:
		x = sha256.Sum256([]byte(input))

	}
	fmt.Printf("input %v\n sha%v hash is = %x\n", input, *version, x)

}

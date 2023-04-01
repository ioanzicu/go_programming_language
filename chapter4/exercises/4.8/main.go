// Modify charcount to count letters, digits, and so on in their Unicode categories,
// using functions like unicode.IsLetter.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letters := make(map[rune]int)  // counts of Unicode letter characters
	digits := make(map[rune]int)   // counts of Unicode digit characters
	graphics := make(map[rune]int) // counts of Unicode graphic character

	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)

	k := 10
	for k >= 0 {
		k--
		r, n, err := in.ReadRune() // returns rune, nbytes, error

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsDigit(r) {
			digits[r]++
		}
		if unicode.IsLetter(r) {
			letters[r]++
		}
		if unicode.IsGraphic(r) {
			graphics[r]++
		}

		utflen[n]++
	}
	fmt.Printf("rune\tdigits\n")
	for c, n := range digits {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("rune\tletters\n")
	for c, n := range letters {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("rune\tgraphics\n")
	for c, n := range graphics {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tutflen\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

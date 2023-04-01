// Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int) // count of words

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	k := 5
	for k >= 0 && input.Scan() {
		k--

		words[input.Text()]++
	}

	fmt.Printf("\n\ncount\tword\tlength\n")
	for w, i := range words {
		fmt.Printf("%d\t%q\t%d\n", i, w, len(w))
	}
}

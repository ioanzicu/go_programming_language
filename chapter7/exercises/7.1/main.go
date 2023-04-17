// Using the ideas from ByteCounter, implement counters for
// words and for lines. You will find bufio.ScanWords useful.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return int(*c), nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("hello this is Johny"))
	fmt.Println(wc) // "4" words

	wc = 0 // reset the counter
	var name = "Ioan"
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Println(wc) // "2" words

	var lc LineCounter
	lc.Write([]byte("hello\n this\n is\n Johny\n\n"))
	fmt.Println(lc) // "5"

	lc = 0 // reset the counter
	fmt.Fprintf(&lc, "hello,\n %s\n", name)
	fmt.Println(lc) // "2"
}

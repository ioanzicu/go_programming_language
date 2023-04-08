// Write a function to populate a mapping from element names—p, div, span, and
// so on—to the number of elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
go build .
cat out.txt | ./outline
*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	counter := map[string]uint16{}
	countTags(nil, doc, counter)

	fmt.Println("\nTags \t\t Count\n")
	for tag, count := range counter {
		spaces := "\t"
		if len(tag) < 4 {
			spaces = "\t\t"
		}
		fmt.Println("->", tag, spaces, count)
	}
}

func countTags(stack []string, n *html.Node, counter map[string]uint16) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		for _, tag := range stack {
			counter[tag]++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countTags(stack, c, counter)
	}
}

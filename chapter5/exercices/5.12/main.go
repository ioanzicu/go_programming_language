// The startElement and endElement functions in gopl.io/ch5/outline2
// (§5.5) share a global variable, depth. Turn them into anonymous functions that share a vari-
// able local to the outline function.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
go build .
cat out.txt | ./outline2
*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
		os.Exit(1)
	}

	outline(doc)
}

func outline(n *html.Node) {

	depth := 0

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(n, startElement, endElement)
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

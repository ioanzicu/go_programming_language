// Modify forEachNode so that the pre and post functions return a boolean result
// indicating whether to continue the traversal. Use it to write a function ElementByID with the
// following signature that finds the first HTML element with the specified id attribute. The
// function should stop the traversal as soon as a match is found.
// func ElementByID(doc *html.Node, id string) *html.Node

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

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Please enter HTML element id")
		os.Exit(1)
	}

	elemntId := os.Args[1]

	n := ElementretuByID(doc, elemntId)
	if n == nil {
		fmt.Fprintf(os.Stderr, "HTML element with id '%s' was not found\n", elemntId)
	} else {
		fmt.Fprintf(os.Stdout, "HTML element with id '%s' was FOUND!!!\n", elemntId)
		for _, attr := range n.Attr {
			fmt.Fprintf(os.Stdout, "<%s> has '%s' element with value = '%s'\n", n.Data, attr.Key, attr.Val)
		}
	}
}

func ElementretuByID(n *html.Node, id string) *html.Node {
	if n == nil {
		return nil
	}
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
		return true
	}
	return forEachElement(n, pre, nil)
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	nodes := []*html.Node{}
	nodes = append(nodes, n)
	for len(nodes) > 0 {
		n = nodes[0]
		nodes = nodes[1:]

		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nodes = append(nodes, c)
		}

		if post != nil {
			if !post(n) {
				return n
			}
		}
	}

	return nil
}

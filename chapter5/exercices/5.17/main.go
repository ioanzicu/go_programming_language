// Write a variadic function ElementsByTagName that, given an HTML node tree
// and zero or more names, returns all the elements that match one of those names.
// Here are two example calls:
//
// func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
go build .
cat out.txt | ./varadic

cat out.txt | go run main.go
*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "varadic: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Please enter HTML element tags: \"img h1 h2 h3\"\n")
		os.Exit(1)
	}

	tagNames := os.Args[1:]

	nodes := ElementsByTagName(doc, tagNames...)
	if nodes == nil {
		fmt.Fprintf(os.Stderr, "HTML element with tag names '%s' was not found\n", tagNames)
	} else {
		fmt.Fprintf(os.Stdout, "HTML element with tag names '%s' was FOUND!!!\n", tagNames)
		for _, node := range nodes {
			for _, attr := range node.Attr {
				fmt.Fprintf(os.Stdout, "<%s> has '%s' element with value = '%s'\n", node.Data, attr.Key, attr.Val)
			}
		}
	}
}

func ElementsByTagName(n *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	if len(names) == 0 {
		return nil
	}

	if n.Type == html.ElementNode {
		for _, tag := range names {
			if n.Data == tag {
				nodes = append(nodes, n)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, names...)...)
	}
	return nodes
}

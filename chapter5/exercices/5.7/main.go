// Develop startElement and endElement into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use
// short forms like <img/> instead of <img></img> when an element has no children. Write a
// test to ensure that the output can be parsed successfully. (See Chapter 11.)
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

/*
go build .
cat out.txt | ./pretty-print
*/

var depth int = 0

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pretty-print: %v\n", err)
		os.Exit(1)
	}

	prettyPrint(doc)
}

func prettyPrint(n *html.Node) {
	start(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		prettyPrint(c)
	}
	end(n)
}

func startElement(n *html.Node) {
	end := ">"

	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, "")
	}

	name := n.Data

	fmt.Printf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}

func startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	fmt.Printf("%*s%s\n", depth*2, "", n.Data)
}

func startComment(n *html.Node) {
	fmt.Printf("<!--%s-->\n", n.Data)
}

func start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		startElement(n)
	case html.TextNode:
		startText(n)
	case html.CommentNode:
		startComment(n)
	}
}

func end(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		endElement(n)
	}
}

func endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
}

// Extend the visit function so that it extracts other kinds of links
// from the document, such as images, scripts, and style sheets.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
go build .
cat out.txt | ./findlinks1
*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {

	// hyperlink
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "script" || n.Data == "link" || n.Data == "img") {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

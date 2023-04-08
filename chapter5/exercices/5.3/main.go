// Write a function to print the contents of all text nodes in an HTML document
// tree. Do not descend into <script> or <style> elements, since their contents are not visible
// in a web browser.
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

	nodeData(nil, doc)
}

func nodeData(stack []string, n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	if n.Type == html.TextNode {
		// fmt.Println(n.Data)
		fmt.Printf("%+#v\n", n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodeData(stack, c)
	}
}

package main

import (
	"fmt"
	"os"
	"testing"

	"bytes"
	"golang.org/x/net/html"
)

func TestElementretuByID(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	data, err := os.ReadFile(path + "/out.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	elementId := "menu-item-1645"
	n := ElementretuByID(doc, elementId)
	if n == nil {
		fmt.Fprintf(os.Stderr, "Expected non nil %+#v", n)
	}

	if n.Data != "li" {
		fmt.Fprintf(os.Stderr, "Expected 'li' HTML element, got %+#v\n", n.Data)
	}
	if n.Attr[0].Val != elementId {
		fmt.Fprintf(os.Stderr, "Expected HTML element to have element id %v, got %+#v\n", elementId, n.Attr[0].Val)
	}

	classVal := "menu-item menu-item-type-taxonomy menu-item-object-category menu-item-1645"
	if n.Attr[1].Val != classVal {
		fmt.Fprintf(os.Stderr, "Expected HTML element to have class value %v, got %+#v\n", classVal, n.Attr[1].Val)
	}
}

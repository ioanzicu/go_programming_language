package main

import (
	"os"
	"testing"

	"bytes"
	"golang.org/x/net/html"
)

func TestElementretuByID(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	data, err := os.ReadFile(path + "/out.txt")
	if err != nil {
		t.Error(err)
	}

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}

	nodes := ElementsByTagName(doc, "h1", "h2", "h3")
	if nodes == nil {
		t.Errorf("Expected non nil %+#v", nodes)
	}

	if nodes[0].Data != "h1" {
		t.Errorf("Expected 'h1' HTML element, got %+#v\n", nodes[0].Data)
	}

	if nodes[1].Data != "h2" {
		t.Errorf("Expected 'h2' HTML element, got %+#v\n", nodes[0].Data)
	}
	classVal := "page-title"
	if nodes[1].Attr[0].Val != classVal {
		t.Errorf("Expected HTML element to have class value %v, got %+#v\n", classVal, nodes[1].Attr[0].Val)
	}

	if nodes[2].Data != "h3" {
		t.Errorf("Expected 'h3' HTML element, got %+#v\n", nodes[0].Data)
	}
}

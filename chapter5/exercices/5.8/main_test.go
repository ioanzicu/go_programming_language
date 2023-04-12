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

	elementId := "menu-item-1645"
	n := ElementByID(doc, elementId)
	if n == nil {
		t.Errorf("Expected non nil %+#v", n)
	}

	if n.Data != "li" {
		t.Errorf("Expected 'li' HTML element, got %+#v\n", n.Data)
	}
	if n.Attr[0].Val != elementId {
		t.Errorf("Expected HTML element to have element id %v, got %+#v\n", elementId, n.Attr[0].Val)
	}

	classVal := "menu-item menu-item-type-taxonomy menu-item-object-category menu-item-1645"
	if n.Attr[1].Val != classVal {
		t.Errorf("Expected HTML element to have class value %v, got %+#v\n", classVal, n.Attr[1].Val)
	}
}

package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var wc WordCounter
	wc.Write([]byte("hello this is Johny"))

	if fmt.Sprint(wc) != "4" {
		t.Error("expected 4, got", wc)
	}

	wc = 0 // reset the counter
	var name = "Ioan"
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Println(wc) // "2" words
	if fmt.Sprint(wc) != "2" {
		t.Error("expected 2, got", wc)
	}
}

func TestLineCounter(t *testing.T) {
	var lc LineCounter
	lc.Write([]byte("hello\n this\n is\n Johny\n\n"))
	if fmt.Sprint(lc) != "5" {
		t.Error("expected 5, got", lc)
	}

	lc = 0 // reset the counter
	var name = "Ioan"
	fmt.Fprintf(&lc, "hello,\n %s\n", name)
	if fmt.Sprint(lc) != "2" {
		t.Error("expected 2, got", lc)
	}
}

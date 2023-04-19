package main

import (
	"bytes"
	"testing"
)

func TestNewReader(t *testing.T) {
	s := "Noroc Ioan"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Errorf("Expected %+v, got %+v; err %+v", len(s), n, err)
	}

	if b.String() != s {
		t.Errorf("Expected %+v, got %+v", s, b.String())
	}
}

func TestNewReaderHTML(t *testing.T) {
	s := "<html><head><title>Slipkont</title></head><body><h1>I am probably wrong, but I am better than you</h1></body></>html>"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Errorf("Expected %+v, got %+v; err %+v", len(s), n, err)
	}

	if b.String() != s {
		t.Errorf("Expected %+v, got %+v", s, b.String())
	}
}

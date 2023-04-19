package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hi Ioan"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 3)
	n, _ := b.ReadFrom(r)
	if n != 3 {
		t.Error("Expected n = 3, got", n)
	}
	if b.String() != "hi " {
		t.Errorf("Expected \"hi \", got %+#v", b.String())
	}
}

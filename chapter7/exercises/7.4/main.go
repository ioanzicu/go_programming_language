// The strings.NewReader function returns a value that satisfies
// the io.Reader interface (and others) by reading from its argument,
// a string. Implement a simple version of NewReader yourself, and
// use it to make the HTML parser (§5.2) take input from a string.
package main

import (
	"bytes"
	"fmt"
	"io"
)

type StringReader struct {
	s string
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, []byte(r.s))
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &StringReader{s}
}

func main() {
	s := "Noroc Ioan"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		fmt.Printf("Expected %+v, got %+v; err %+v", len(s), n, err)
	}

	if b.String() != s {
		fmt.Printf("Expected %+v, got %+v", s, b.String())
	}
}

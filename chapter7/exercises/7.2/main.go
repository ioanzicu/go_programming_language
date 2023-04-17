// Write a function CountingWriter with the signature below that, given an
// io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer.
// func CountingWriter(w io.Writer) (io.Writer, *int64)
package main

import (
	"bytes"
	"fmt"
	"io"
)

type CounterWrite struct {
	writer io.Writer
	count  int64
}

func (c *CounterWrite) Write(p []byte) (n int, err error) {
	n, err = c.writer.Write(p)
	c.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw = &CounterWrite{
		writer: w,
		count:  0,
	}
	return cw, &cw.count
}

func main() {
	w, c := CountingWriter(new(bytes.Buffer))
	msg1 := "hello banana!!!"
	w.Write([]byte(msg1))

	msg2 := "hello kiwi!!!"
	w.Write([]byte(msg2))

	fmt.Println("Expected", *c, "==", len(msg1)+len(msg2))
}

package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w, c := CountingWriter(new(bytes.Buffer))
	msg1 := "hello banana!!!"
	w.Write([]byte(msg1))

	msg2 := "hello kiwi!!!"
	w.Write([]byte(msg2))
	totalLen := int64(len(msg1) + len(msg2))
	if *c != totalLen {
		fmt.Println("Expected", *c, ", got", totalLen)
	}
}
